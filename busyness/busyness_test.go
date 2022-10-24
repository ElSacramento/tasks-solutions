package busyness

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teambition/rrule-go"
)

func TestGetRowIndex(t *testing.T) {
	// one week with 30min slots
	b := NewBusyness(7, 8, time.Minute*30)
	currentTime := time.Now()

	t.Run("start day", func(t *testing.T) {
		index := b.getRowIndex(currentTime.Add(2 * time.Hour))
		require.Equal(t, 0, index)
	})
	t.Run("another day", func(t *testing.T) {
		index := b.getRowIndex(currentTime.Add(48 * time.Hour))
		require.Equal(t, 2, index)
	})
	t.Run("out of the period", func(t *testing.T) {
		index := b.getRowIndex(currentTime.Add(8 * 24 * time.Hour))
		require.Greater(t, index, len(b.value))
	})
}

func TestGetRowsInterval(t *testing.T) {
	// one week with 30min slots
	b := NewBusyness(7, 8, time.Minute*30)
	currentTime := time.Now()

	t.Run("same day", func(t *testing.T) {
		indexes, err := b.getRowsInterval(currentTime, currentTime.Add(2*time.Hour), nil)
		require.Nil(t, err)
		require.Len(t, indexes, 1)
	})
	t.Run("two days meeting", func(t *testing.T) {
		indexes, err := b.getRowsInterval(currentTime, currentTime.Add(28*time.Hour), nil)
		require.Nil(t, err)
		require.Len(t, indexes, 2)
	})
	t.Run("fromDay is out of the period", func(t *testing.T) {
		indexes, err := b.getRowsInterval(currentTime.Add(8*24*time.Hour), currentTime.Add(9*24*time.Hour), nil)
		require.Nil(t, indexes)
		require.Equal(t, "event is out of the period", err.Error())
	})
	t.Run("toDay is out of the period", func(t *testing.T) {
		indexes, err := b.getRowsInterval(currentTime, currentTime.Add(9*24*time.Hour), nil)
		require.Nil(t, err)
		require.Len(t, indexes, 7)
	})
}

func TestGetColumnIndex(t *testing.T) {
	// one week with 30min slots
	b := NewBusyness(7, 8, time.Minute*30)
	currentTime := time.Now()
	year, month, day := currentTime.Date()

	t.Run("first interval", func(t *testing.T) {
		// 00:01:00
		index := b.getColumnIndex(time.Date(year, month, day, 0, 1, 0, 0, currentTime.Location()))
		require.Equal(t, 0, index)
	})
	t.Run("last interval", func(t *testing.T) {
		// 23:55:00
		index := b.getColumnIndex(time.Date(year, month, day, 23, 55, 0, 0, currentTime.Location()))
		require.Equal(t, len(b.value[0])-1, index)
	})
	t.Run("another interval", func(t *testing.T) {
		// 02:24:00 -> 4
		index := b.getColumnIndex(time.Date(year, month, day, 02, 24, 0, 0, currentTime.Location()))
		require.Equal(t, 4, index)
	})
}

func TestGetColumnsInterval(t *testing.T) {
	// one week with 30min slots
	b := NewBusyness(7, 8, time.Minute*30)
	currentTime := time.Now()

	t.Run("min interval", func(t *testing.T) {
		indexes, err := b.getColumnsInterval(currentTime, currentTime.Add(30*time.Minute))
		require.Nil(t, err)
		require.Len(t, indexes, 1)
	})
	t.Run("one hour", func(t *testing.T) {
		indexes, err := b.getColumnsInterval(currentTime, currentTime.Add(time.Hour))
		require.Nil(t, err)
		require.Len(t, indexes, 2)
	})
	t.Run("5,5 hours", func(t *testing.T) {
		indexes, err := b.getColumnsInterval(currentTime, currentTime.Add(5*time.Hour+30*time.Minute))
		require.Nil(t, err)
		require.Len(t, indexes, 11)
	})
}

type testMeeting struct {
	from time.Time
	to   time.Time
	rule *rrule.RRule
}

func TestBusyness(t *testing.T) {
	// one week with 30min slots
	b := NewBusyness(7, 8, time.Minute*30)

	currentTime := time.Now()

	firstMeeting := testMeeting{
		from: currentTime.Add(4 * time.Hour),
		to:   currentTime.Add(5 * time.Hour),
	}
	result, err := b.BookIfPossible(firstMeeting.from, firstMeeting.to, firstMeeting.rule)
	require.Nil(t, err)
	require.True(t, result)

	secondMeeting := testMeeting{
		from: currentTime.Add(50 * time.Hour),
		to:   currentTime.Add(50*time.Hour + 30*time.Minute),
	}
	result, err = b.BookIfPossible(secondMeeting.from, secondMeeting.to, secondMeeting.rule)
	require.Nil(t, err)
	require.True(t, result)

	everydayMeeting, err := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Dtstart: currentTime.Add(2 * time.Hour),
		Count:   7,
	})
	require.Nil(t, err)
	thirdMeeting := testMeeting{
		from: currentTime.Add(6 * time.Hour),
		to:   currentTime.Add(7 * time.Hour),
		rule: everydayMeeting,
	}
	result, err = b.BookIfPossible(thirdMeeting.from, thirdMeeting.to, thirdMeeting.rule)
	require.Nil(t, err)
	require.True(t, result)
}
