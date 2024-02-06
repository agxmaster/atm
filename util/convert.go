package util

import (
	"bytes"
	"strconv"
)

func ArrayToString[T int | int8 | int16 | int32 | int64](Arr []T, delim string) string {

	var buffer bytes.Buffer
	for i := 0; i < len(Arr); i++ {
		buffer.WriteString(strconv.Itoa(int(Arr[i])))
		if i != len(Arr)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}
