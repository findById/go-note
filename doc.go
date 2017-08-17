package main

import (
	"time"
)

type Doc struct {
	Permalink string

	Title  string
	Desc   string
	Date   string
	Tag    string
	Author string
}

type DocSlice []Doc

func (a DocSlice) Len() int {
	return len(a)
}
func (a DocSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a DocSlice) Less(i, j int) bool {
	time1 := a[i].Date
	time2 := a[j].Date
	t1, err := time.Parse("2006-01-02 15:04:05 -0700", time1)
	if err != nil {
		return true
	}
	t2, err := time.Parse("2006-01-02 15:04:05 -0700", time2)
	if err != nil {
		return false
	}
	return t1.Before(t2)
}
