package utils

import (
	"strconv"
	"fmt"
)

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

func PageUtil(count64 int64, pageNo int, pageSize int) *Page {
	string := strconv.FormatInt(count64, 10)
	count, _ := strconv.Atoi(string)
	tp := count / pageSize
	if count%pageSize > 0 {
		tp += 1
	}
	fmt.Println("tp:",tp,"num",pageNo)
	if tp < pageNo {
		pageNo = tp
	}
	return &Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp}
}