package main

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

type NoticeON struct {
	Old     map[string][]int
	OldKeys []string
	New     map[string][]int
	NewKeys []string
}

func Selector(myTicker *time.Ticker, noticeON *NoticeON) {
	tempMap := make(map[string][]int)
	millis := GetMillis()
	rand.Seed(millis)
	for i := 0; i < 5; i++ {
		n := rand.Int63n(60000)
		str := strconv.Itoa(int(millis + n))
		list := []int{i, i + 1, i + 2, i + 3}
		tempMap[str] = list
	}
	noticeON.OldKeys = noticeON.OldKeys[0:0]
	for key := range tempMap {
		noticeON.OldKeys = append(noticeON.OldKeys, key)
	}
	sort.Strings(noticeON.OldKeys)
	noticeON.Old = tempMap
	for {
		tempMap = make(map[string][]int)
		millis = GetMillis()
		rand.Seed(millis)
		for i := 0; i < 5; i++ {
			n := rand.Int63n(60000)
			str := strconv.Itoa(int(millis + n + 60000))
			list := []int{i, i + 1, i + 2, i + 3}
			tempMap[str] = list
			//println("运行在这个时候：", TimestampToString((millis+n+60000)/1000))
		}
		noticeON.NewKeys = noticeON.NewKeys[0:0]
		for key := range tempMap {
			noticeON.NewKeys = append(noticeON.NewKeys, key)
		}
		sort.Strings(noticeON.NewKeys)
		noticeON.New = tempMap

		<-myTicker.C

		noticeON.Old = noticeON.New
		noticeON.OldKeys = noticeON.NewKeys
	}
}

// write data to channel
func Writer(noticeON *NoticeON, bufChan chan int) {
	for {
		if noticeON.Old != nil && len(noticeON.OldKeys) > 0 {
			key := noticeON.OldKeys[0]
			r, err := strconv.ParseInt(key, 10, 64)
			if err != nil {
				println(err)
			}
			millis := GetMillis()

			internal := r - millis
			//println("millis：", TimestampToString(GetMillis()/1000))
			println("r     ：", TimestampToString(r/1000))
			//println("Sleep ：", internal)

			time.Sleep(time.Duration(internal) * time.Millisecond)
			//println("时间到：", TimestampToString(GetMillis()/1000))
			for _, v2 := range noticeON.Old[key] {
				bufChan <- v2
				//fmt.Fprintf(os.Stderr, "%v write: %d", os.Getpid(), v2)
			}
			delete(noticeON.Old, key)
			i := 0
			noticeON.OldKeys = append(noticeON.OldKeys[0:i], noticeON.OldKeys[i+1:]...)
		}
	}
}

func Reader(name string, bufChan chan int) {
	for {
		r := <-bufChan
		//fmt.Printf("%s read: %d", name, r)
		r = r
	}
}

func NewBee() {
	println("定时器启动：" + TimestampToString(GetMillis()/1000))
	bufChan := make(chan int, 1000)

	noticeON := &NoticeON{Old: nil, New: nil}
	for i := 0; i < 1; i++ {
		myTicker := time.NewTicker(time.Minute * 1)
		go Selector(myTicker, noticeON)
	}

	// 开启多个writer的goroutine，不断地向channel中写入数据
	for i := 0; i < 1; i++ {
		go Writer(noticeON, bufChan)
	}
	// 开启多个reader的goroutine，不断的从channel中读取数据，并处理数据
	for i := 0; i < 2; i++ {
		go Reader("read "+strconv.Itoa(i), bufChan)
	}
}

func TestDdd(t *testing.T) {
	//go NewBee()
	go Tesgsegs()
	select {}
}

func Tesgsegs() {
	for {
		b := 5 + rand.Intn(5)
		println("b:", b)
		time.Sleep(1 * time.Second)
		println("---------")
		a := b
		println("a:", a)
	}
}
