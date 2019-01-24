package slice

import (
	"strconv"
	"unsafe"
)

func SliceInt2String1(s []int) string {
	if len(s) < 1 {
		return ""
	}

	ss := strconv.Itoa(s[0])
	for i := 1; i < len(s); i++ {
		ss += "," + strconv.Itoa(s[i])
	}

	return ss
}

func SliceInt2String2(s []int) string {
	if len(s) < 1 {
		return ""
	}

	b := make([]byte, 0, 256)
	b = append(b, strconv.Itoa(s[0])...)
	for i := 1; i < len(s); i++ {
		b = append(b, ',')
		b = append(b, strconv.Itoa(s[i])...)
	}

	//return *(*string)(unsafe.Pointer(&b))
	return string(b)
}

func SliceInt2String3(s []int) string {
	if len(s) < 1 {
		return ""
	}

	b := make([]byte, 0, 256)
	b = append(b, strconv.Itoa(s[0])...)
	for i := 1; i < len(s); i++ {
		b = append(b, ',')
		b = append(b, strconv.Itoa(s[i])...)
	}

	return *(*string)(unsafe.Pointer(&b))
}

func SliceInt2String4(s [][]int) []string {
	res := make([]string, len(s))
	for i, v := range s {
		if len(v) < 1 {
			res[i] = ""
			continue
		}

		res[i] += strconv.Itoa(v[0])
		for j := 1; j < len(v); j++ {
			res[i] += "," + strconv.Itoa(v[j])
		}
	}

	return res
}

func SliceInt2String5(s [][]int) []string {
	res := make([]string, len(s))
	b := make([]byte, 0, 256)
	for i, v := range s {
		if len(v) < 1 {
			res[i] = ""
			continue
		}

		b := b[:0]
		b = append(b, strconv.Itoa(v[0])...)
		for j := 1; j < len(v); j++ {
			b = append(b, ',')
			b = append(b, strconv.Itoa(v[j])...)
		}
		res[i] = string(b)
	}

	return res
}
