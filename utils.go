package main

import (
	"strconv"
	"time"
)

func GenerateTimestampBasedID() int64 {
	return time.Now().UnixNano()
}

func generateBookID(number int64) string {
	return "book" + strconv.FormatInt(number, 10)
}

func generateMemberID(number int64) string {
	return "member" + strconv.FormatInt(number, 10)
}
