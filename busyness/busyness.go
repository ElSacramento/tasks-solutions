package busyness

import (
	"errors"
	"sync"
	"time"

	"github.com/teambition/rrule-go"
)

/*
	10:00 - 11:00 	11:00 - 12:00	12:00 - 13:00

day1	0		1		0
day2	1		1		1
day3	1		0		1
*/
type Busyness struct {
	interval    time.Duration
	startDate   time.Time
	dayBytesLen int

	rwLock *sync.RWMutex
	value  []byte
}

const (
	minPossibleInterval = time.Minute * 15
	maxPossibleInterval = time.Hour * 3
	bitsLen             = 8
)

func NewBusyness(storedDays int, interval time.Duration) (*Busyness, error) {
	if interval < minPossibleInterval {
		return nil, errors.New("too small interval")
	}
	if interval > maxPossibleInterval {
		return nil, errors.New("too big interval")
	}

	// for interval=30min we need 48bits = 6bytes
	// for 7 days we need 6bytes*7 = 42bytes
	bytesNumberForDay := int((24 * time.Hour / interval) / bitsLen)
	value := make([]byte, bytesNumberForDay*storedDays, bytesNumberForDay*(storedDays+1))

	year, month, day := time.Now().Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())

	return &Busyness{
		interval:    interval,
		value:       value,
		startDate:   startDate,
		dayBytesLen: bytesNumberForDay,
		rwLock:      &sync.RWMutex{},
	}, nil
}

func (b *Busyness) getDaysLen() int {
	return len(b.value) / b.dayBytesLen
}

func (b *Busyness) getDayBitsLen() int {
	return b.dayBytesLen * bitsLen
}

func (b *Busyness) geStoredBitsLen() int {
	return len(b.value) * bitsLen
}

func (b *Busyness) getBitIndex(t time.Time) int {
	index := int(t.Sub(b.startDate).Minutes() / b.interval.Minutes())
	if index < 0 {
		return -1
	}
	return index
}

func (b *Busyness) getTimeFromIndex(index int) time.Time {
	return b.startDate.Add(time.Duration(index * int(b.interval)))
}

func (b *Busyness) getMeetingLength(m *Meeting) (int, error) {
	fromIndex := b.getBitIndex(m.from)
	toIndex := b.getBitIndex(m.to)

	if fromIndex == -1 || toIndex == -1 {
		return 0, errors.New("wrong fromDay or toDay time")
	}

	return toIndex - fromIndex, nil
}

func (b *Busyness) getStartIndexes(m *Meeting) ([]int, error) {
	fromIndex := b.getBitIndex(m.from)
	if fromIndex >= b.geStoredBitsLen() {
		return nil, errors.New("event is out of the period")
	}

	if m.rule == nil {
		return []int{fromIndex}, nil
	}

	indexes := make([]int, 0)
	iterValue := m.rule.Iterator() // complexity?
	event, hasNext := iterValue()
	for hasNext {
		index := b.getBitIndex(event)
		if index == -1 {
			return nil, errors.New("wrong day")
		}

		if index >= b.geStoredBitsLen() {
			return indexes, nil
		}

		indexes = append(indexes, index)
		event, hasNext = iterValue()
	}

	return indexes, nil
}

func getByteIndexes(index int) (int, int) {
	div := index / bitsLen
	mod := index % bitsLen
	return div, mod
}

type ByteInfo struct {
	index int
	value byte
}

func (b *Busyness) convertToBytes(startIndexes []int, meetingLength int) []ByteInfo {
	expectedBytes := make([]ByteInfo, 0)
	for _, i := range startIndexes {
		byteFromIndex, bitFromIndex := getByteIndexes(i)
		bitToIndex := bitFromIndex + meetingLength - 1
		byteToIndex := byteFromIndex
		if bitToIndex >= bitsLen {
			byteToIndex += 1
			bitToIndex -= bitsLen
		}

		if byteFromIndex == byteToIndex {
			// create 00001110
			// meetingLength = 3
			// bitFromIndex = 4, bitToIndex = 6
			res := byte(1)
			for i := 0; i < bitToIndex-bitFromIndex; i++ {
				res = res << 1
				res = res + byte(1)
			}
			// now we have 00000111
			res = res << byte(bitsLen-bitToIndex-1)
			// now we have 00001110
			expectedBytes = append(expectedBytes, ByteInfo{
				index: byteFromIndex,
				value: res,
			})
			continue
		}

		// create 00000011 and 11000000
		// meetingLength = 4
		// bitFromIndex = 6, bitToIndex = 1
		res := byte(1)
		for i := 0; i <= bitsLen-bitFromIndex-1; i++ {
			res = res << 1
			res = res + byte(1)
		}
		// now we have 00000011
		expectedBytes = append(expectedBytes, ByteInfo{
			index: byteFromIndex,
			value: res,
		})

		// if period longer that 2 days
		for i := 1; i < byteToIndex-byteFromIndex; i++ {
			expectedBytes = append(expectedBytes, ByteInfo{
				index: byteFromIndex + i,
				value: byte(255),
			})
		}

		res = byte(1)
		for i := 0; i < bitToIndex; i++ {
			res = res << 1
			res = res + byte(1)
		}
		// now we have 00000011
		res = res << byte(bitsLen-bitToIndex-1)
		// now we have 11000000
		expectedBytes = append(expectedBytes, ByteInfo{
			index: bitToIndex,
			value: res,
		})
	}
	return expectedBytes
}

func (b *Busyness) bookTimeSlot(bytesInfo []ByteInfo) bool {
	defer b.rwLock.Unlock()
	b.rwLock.Lock()

	for _, info := range bytesInfo {
		if info.index >= len(b.value) {
			break
		}
		if b.value[info.index]&info.value != 0 {
			return false
		}
	}
	for _, info := range bytesInfo {
		if info.index >= len(b.value) {
			break
		}
		b.value[info.index] = b.value[info.index] | info.value
	}
	return true
}

func (b *Busyness) checkTimeSlot(bytesInfo []ByteInfo) bool {
	defer b.rwLock.RUnlock()
	b.rwLock.RLock()

	for _, info := range bytesInfo {
		if info.index >= len(b.value) {
			break
		}
		if b.value[info.index]&info.value != 0 {
			return false
		}
	}
	return true
}

type Meeting struct {
	from time.Time
	to   time.Time
	rule *rrule.RRule
}

func (b *Busyness) AppendDay(meetings []Meeting) error {
	// todo: check that day to append doesn't exist in b.value

	defer b.rwLock.Unlock()
	b.rwLock.Lock()

	bytesInfo := make([]ByteInfo, 0)
	for _, m := range meetings {
		meetingLength, getMLengthErr := b.getMeetingLength(&m)
		if getMLengthErr != nil {
			return getMLengthErr
		}

		if meetingLength == 0 {
			continue
		}

		fromIndex := b.getBitIndex(m.from)
		bytesInfo = append(bytesInfo, b.convertToBytes([]int{fromIndex}, meetingLength)...)
	}

	newDay := make([]byte, b.dayBytesLen)
	b.value = append(b.value, newDay...)

	for _, info := range bytesInfo {
		if info.index >= len(b.value) {
			continue
		}
		if b.value[info.index]&info.value != 0 {
			return errors.New("crossing events")
		}
		b.value[info.index] = b.value[info.index] | info.value
	}

	b.value = b.value[b.dayBytesLen:]
	b.startDate = b.startDate.Add(24 * time.Hour)
	return nil
}

// func (b *Busyness) getPossiblePairs(currentIndex int, startPair pairInt) []pairInt {
// 	result := make([]pairInt, 0)
// 	i := 1
// 	withLeft := startPair[0]-i > currentIndex
// 	// O(n)
// 	for (startPair[1]+i < b.getColumnsLen()) || (startPair[0]-i >= 0) {
// 		if withLeft && (startPair[0]-i >= 0) {
// 			leftPair := pairInt{startPair[0] - i, startPair[1] - i}
// 			result = append(result, leftPair)

// 			withLeft = startPair[0]-(i+1) > currentIndex
// 		}

// 		if startPair[1]+i < b.getColumnsLen() {
// 			rightPair := pairInt{startPair[0] + i, startPair[1] + i}
// 			result = append(result, rightPair)
// 		}

// 		i += 1
// 	}
// 	return result
// }

func (b *Busyness) BookOrGetAvailableSlots(m *Meeting, maxNumber int) (bool, []Meeting, error) {
	// todo: validate that from, to and rrule in our stored period of time
	// todo: validate that reccurent event from and to happen in one day
	// todo: validate that from < to, from > time.Now

	meetingLength, getMLengthErr := b.getMeetingLength(m)
	if getMLengthErr != nil {
		return false, nil, getMLengthErr
	}

	startIndexes, getIndexes := b.getStartIndexes(m)
	if getIndexes != nil {
		return false, nil, getIndexes
	}

	if meetingLength == 0 {
		return false, nil, errors.New("empty period")
	}

	bytesInfo := b.convertToBytes(startIndexes, meetingLength)
	bookStatus := b.bookTimeSlot(bytesInfo)
	if bookStatus {
		return true, nil, nil
	}

	// dont need available slots
	if maxNumber <= 0 {
		return false, nil, nil
	}
	return false, nil, nil

	// currentIndex := b.getBitIndex(time.Now())

	// possibleIntervals := b.getPossiblePairs(currentIndex, columnsInterval)

	// possibleMeetings := make([]Meeting, 0)
	// for _, interval := range possibleIntervals {
	// 	chekStatus := b.checkTimeSlot(rows, interval)
	// 	if chekStatus {
	// 		newMeeting := Meeting{
	// 			from: b.getTimeFromRowColumn(rows[0], interval[0]),
	// 			to:   b.getTimeFromRowColumn(rows[0], interval[1]+1),
	// 		}
	// 		if m.rule != nil {
	// 			newRule := *m.rule
	// 			newRule.DTStart(newMeeting.from)
	// 			newMeeting.rule = &newRule
	// 		}

	// 		possibleMeetings = append(possibleMeetings, newMeeting)
	// 	}

	// 	if len(possibleMeetings) == maxNumber {
	// 		return false, possibleMeetings, nil
	// 	}
	// }
	// return false, possibleMeetings, nil
}
