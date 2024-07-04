package utils

import (
	"time"

	"github.com/sqids/sqids-go"
)

type IDMaker struct {
}

func NewIDMaker() *IDMaker {
	return &IDMaker{}
}

func (idm *IDMaker) MakeID() string {
	now := time.Now().Unix()
	s, _ := sqids.New()
	id, _ := s.Encode(Int64ToUint64Array(now))
	return id
}

// get time from id
func (idm *IDMaker) GetTimeFromID(id string) time.Time {
    s, _ := sqids.New()
	t := s.Decode(id)
	return time.Unix(Uint64ArrayToInt64(t), 0)
}

// get date from id
func (idm *IDMaker) GetDateFromID(id string) string {
    tm := idm.GetTimeFromID(id)
	return tm.Format("2006-01-02")
}

// transform int64 to uint64 array
func Int64ToUint64Array(num int64) []uint64 {
	var result []uint64
	for num > 0 {
		result = append(result, uint64(num%10))
		num /= 10
	}
	return result
}

// transform uint64 array to int64
func Uint64ArrayToInt64(arr []uint64) int64 {
    var result int64
    for i := len(arr) - 1; i >= 0; i-- {
        v := arr[i]
        result = result*10 + int64(v)
    }
    return result
}