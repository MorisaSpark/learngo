package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	go testTimeTaskA()

	flg := make(chan int)
	defer close(flg)
	<-flg
	for{
		fmt.Println("xxx")
	}
}

func testTimeTaskWait() {
	c := make(chan int)
	select {
	case <-c:
		fmt.Println("没有数据")
	case <-time.After(5 * time.Second):
		fmt.Println("超时退出")
	}
}

func testTimeTaskA() {
	uch := make(chan int)
	defer close(uch)

	i := 0
	c := cron.New()
	spec := "*/1 * * * * ?"
	err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
		if i == 4 {
			uch <- i
		}
	})
	c.Start()
	defer c.Stop()
	if err != nil {

	}

	<-uch
}

func testTimeTask() {
	i := 0
	c := cron.New()
	spec := "*/1 * * * * ?"
	err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
	if err != nil {

	}
	var str string
	fmt.Scan(&str)
}

type myInt int

func testFunctionAndMethod() {
	a := 1
	b := 2
	addFunction(a, b)
	var c myInt = 3
	var d myInt = 4
	c.addMethod(d)
}

func addFunction(a int, b int) {
	fmt.Println(a + b)
}
func (c myInt) addMethod(d myInt) {
	fmt.Println(c + d)
}

func testBasicType() {

}

func test1p1() {
	var b int
	var c int
	_, b, c = testNoName()
	fmt.Println(b, c)
}

func testNoName() (a, b, c int) {
	return 1, 2, 3
}

type BaseResult struct {
	Flag       string //是否成功
	StatusCode string // 状态码
	Message    string // 携带信息
	Data       interface{}
}

type PageResult struct {
	Total int64         //总数量
	Rows  []interface{} //数组
}

type ErrorResult struct {
	Detail  string
	ErrorId string
}

func NewBaseResult(flag, statusCode, message string, data interface{}) BaseResult {
	result := BaseResult{}
	result.Flag = flag
	result.StatusCode = statusCode
	result.Message = message
	result.Data = data
	return result
}
func NewPageResult(total int64, rows []interface{}) PageResult {
	result := PageResult{}
	result.Total = total
	result.Rows = rows
	return result
}
func NewErrorResult(detail, errorId string) ErrorResult {
	result := ErrorResult{}
	result.Detail = detail
	result.ErrorId = errorId
	return result
}
