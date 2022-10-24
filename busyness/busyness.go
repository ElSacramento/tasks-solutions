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
	minInterval time.Duration
	startDate   time.Time

	rwLock *sync.RWMutex
	value  [][]byte
}

func NewBusyness(storedLen, storedCap int, minInterval time.Duration) *Busyness {
	value := make([][]byte, storedLen, storedCap)

	var dayIntervalsNumber = 24 * time.Hour / minInterval
	for i := 0; i < len(value); i++ {
		value[i] = make([]byte, dayIntervalsNumber)
	}

	startDate := time.Now()

	return &Busyness{
		minInterval: minInterval,
		value:       value,
		startDate:   startDate,
		rwLock:      &sync.RWMutex{},
	}
}

func isEqualDate(time1, time2 time.Time) bool {
	year1, month1, day1 := time1.Date()
	year2, month2, day2 := time2.Date()
	return year1 == year2 && month1 == month2 && day1 == day2
}

func (b *Busyness) getRowIndex(t time.Time) int {
	index := int(t.Sub(b.startDate).Hours() / 24)
	if index < 0 {
		return -1
	}
	return index
}

func (b *Busyness) getColumnIndex(t time.Time) int {
	index := (t.Hour()*60 + t.Minute()) / int(b.minInterval.Minutes())
	if index < 0 || index >= len(b.value[0]) {
		return -1
	}
	return index
}

func (b *Busyness) getReccurentRowsIDs(rule *rrule.RRule) ([]int, error) {
	indexes := make([]int, 0)

	iterValue := rule.Iterator() // complexity?
	event, hasNext := iterValue()
	for hasNext {
		index := b.getRowIndex(event)
		if index == -1 {
			return nil, errors.New("wrong day")
		}

		if index >= len(b.value) {
			return indexes, nil
		}

		indexes = append(indexes, index)
		event, hasNext = iterValue()
	}

	return indexes, nil
}

func (b *Busyness) getColumnsInterval(from, to time.Time) ([]int, error) {
	fromColumn := b.getColumnIndex(from)
	toColumn := b.getColumnIndex(to)

	if fromColumn == -1 || toColumn == -1 {
		return nil, errors.New("impossible time of the day")
	}

	indexes := make([]int, 0)
	for i := fromColumn; i < toColumn; i++ {
		indexes = append(indexes, i)
	}
	return indexes, nil
}

func (b *Busyness) getRowsInterval(from, to time.Time, rule *rrule.RRule) ([]int, error) {
	indexes := make([]int, 0)
	if rule == nil {
		fromRow := b.getRowIndex(from)

		if fromRow == -1 {
			return nil, errors.New("wrong fromDay")
		}

		if fromRow >= len(b.value) {
			return nil, errors.New("event is out of the period")
		}

		toRow := fromRow

		if !isEqualDate(from, to) {
			// event that starts today and will be finished tomorrow
			toRow = b.getRowIndex(to)

			if toRow == -1 {
				return nil, errors.New("wrong toDay")
			}
		}

		for i := fromRow; i <= toRow && i < len(b.value); i++ {
			indexes = append(indexes, i)
		}
		return indexes, nil
	}

	indexes, err := b.getReccurentRowsIDs(rule)
	if err != nil {
		return nil, err
	}
	return indexes, nil
}

func (b *Busyness) BookIfPossible(from, to time.Time, rule *rrule.RRule) (bool, error) {
	// todo: validate that from, to and rrule in our stored period of time
	// todo: validate that reccurent event from and to happen in one day

	columns, getColumnsErr := b.getColumnsInterval(from, to)
	if getColumnsErr != nil {
		return false, getColumnsErr
	}

	rows, getRowsErr := b.getRowsInterval(from, to, rule)
	if getRowsErr != nil {
		return false, getRowsErr
	}

	return b.bookSlot(rows, columns)
}

func (b *Busyness) bookSlot(rows, columns []int) (bool, error) {
	defer b.rwLock.Unlock()
	b.rwLock.Lock()

	// O(days*intervals)*2
	// intervals max value is 48 for 30min minInterval
	// days max value is 90 for 3 months stored period
	for _, i := range rows {
		for _, j := range columns {
			if b.value[i][j] == 1 {
				return false, nil
			}
		}
	}
	for _, i := range rows {
		for _, j := range columns {
			b.value[i][j] = 1
		}
	}
	return true, nil
}

type Meeting struct {
	from time.Time
	to   time.Time
	rule *rrule.RRule
}

type AppendMeeting struct {
	from time.Time
	to   time.Time
}

func (b *Busyness) AppendDay(meetings []AppendMeeting) error {
	// todo: check that day to append doesn't exist in b.value

	defer b.rwLock.Unlock()
	b.rwLock.Lock()

	columnsIndexesToChange := make([]int, 0)
	for _, m := range meetings {
		columns, getColumnsErr := b.getColumnsInterval(m.from, m.to)
		if getColumnsErr != nil {
			return getColumnsErr
		}
		columnsIndexesToChange = append(columnsIndexesToChange, columns...)
	}

	if len(columnsIndexesToChange) > len(b.value[0]) {
		return errors.New("impossible situation")
	}

	newRow := make([]byte, len(b.value[0]))
	for _, ind := range columnsIndexesToChange {
		if ind >= len(newRow) || ind < 0 {
			return errors.New("wrong index to change")
		}
		if newRow[ind] == 1 {
			return errors.New("crossing events")
		}
		newRow[ind] = 1
	}

	b.value = append(b.value, newRow)
	b.value = b.value[1:]
	return nil
}
