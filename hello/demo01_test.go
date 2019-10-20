package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

// write data to channel
func writer(max int, bufChan chan int) {
	for {
		for i := 0; i < max; i++ {
			bufChan <- i
			fmt.Fprintf(os.Stderr, "%v write: %d\n", os.Getpid(), i)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

// read data fro m channel
func reader(name string, bufChan chan int) {
	for {
		r := <-bufChan
		fmt.Printf("%s read value: %d\n", name, r)
	}
}

func testWriterAndReader(max int) {
	var bufChan chan int = make(chan int, 1000)
	//var msgChan chan string = make(chan string)

	// 开启多个writer的goroutine，不断地向channel中写入数据
	for i := 0; i < 3; i++ {
		go writer(max, bufChan)
	}

	// 开启多个reader的goroutine，不断的从channel中读取数据，并处理数据
	for i := 0; i < 10; i++ {
		go reader("read "+strconv.Itoa(i), bufChan)
	}

	//// 获取三个reader的任务完成状态
	//name1 := <-msgChan
	//name2 := <-msgChan
	//name3 := <-msgChan
	//
	//fmt.Println("------", name1, name2, name3)
}

func TestDd(t *testing.T) {
	go testWriterAndReader(10)
	select {}
}
