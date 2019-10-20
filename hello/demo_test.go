package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/garyburd/redigo/redis"
	"github.com/robfig/cron/v3"
	"sync"
	"testing"
	"time"
)

func TestTest(t *testing.T) {
	c := cron.New()
	c.AddFunc("@every 1h1m", func() { fmt.Println("111") })
	c.AddFunc("@hourly 1h1m", func() { fmt.Println("111") })
	//c.AddFunc("0/1 * * * *", func() { fmt.Println("222") })
	//id, _ := c.AddFunc("0/1 * * * *", func() { fmt.Println("444") })
	//c.AddFunc("0/1 * * * *", func() { fmt.Println("333") })
	//c.AddFunc("0/1 * * * *", func() { fmt.Println("555") })
	//c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour1") })
	//c.AddJob("0/1 * * * *", TestJob{})
	//c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	//c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	//c.AddFunc("@hourly", func() { fmt.Println("Every hour, starting an hour from now") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()
	//println(id)
	for _, v := range c.Entries() {
		fmt.Println("v ", v.ID)
	}
	c.Remove(4)
	//for _, v := range c.Entries() {
	//	fmt.Println("v ", v.ID)
	//}

	println(len(c.Entries()))
	//c.AddFunc("0/1 * * * *", func() { fmt.Println("666") })
	for _, v := range c.Entries() {
		fmt.Println("v ", v.ID)
	}

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	uchan := make(chan int)
	defer close(uchan)
	<-uchan
}

func TestCronExp(t *testing.T) {
	var millis int64
	millis = 1570874110340
	millis = ToTimestamp("2019-10-31 01:01:01")
	fmt.Println(millis)
	//fmt.Println("no		", CreateCronSpec(millis, UNREPEATED))
	//fmt.Println("day		", CreateCronSpec(millis, DAY))
	//fmt.Println("week	", CreateCronSpec(millis, WEEK))
	//fmt.Println("month	", CreateCronSpec(millis, MONTH))
	//fmt.Println("year	", CreateCronSpec(millis, YEAR))
}

func TestTimeFormat(t *testing.T) {
	format2 := time.Unix(GetMillis()/1000, 0).Format(timeTemplateStandard)
	fmt.Println(format2)
	parse, _ := time.Parse("2006-01-02 15-04-05", "2019-10-12 12-24-51")

	fmt.Println(parse.Weekday())
}

type TestJob struct {
}

func (this TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("testJob2...")
}

func TestProject2(t *testing.T) {
	// 1. 从  reminds remindUsers  中查出 从当前时间开始 5-10分钟将要执行的提醒
	// 2. 将计划都增加到同一cron中，定时执行
	// 2.1. 对目标用户发送提醒 (remindUsers)
	// 2.2. 更新 (reminds, remindLog) 中数据
}

func TestTemp(t *testing.T) {
	c := make(chan int)
	num := 0
	//创建一个启动goroutine的匿名函数
	go func() {
		for v := range c {
			num++
			if num&15 == 0 {
				fmt.Println()
			}
			fmt.Print(v, " ")
		}
	}()
	go func() {
		for v := range c {
			num++
			if num&15 == 0 {
				fmt.Println()
			}
			fmt.Print(v, "@")
		}
	}()

	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}

func TestTimer1(t *testing.T) {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stop")
	}
}
func TestTimer2(t *testing.T) {
	input := make(chan interface{})

	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	i := 0
	for {
		i++
		select {
		case <-input:
			println("input")
		case <-t1.C:
			println("5s ", i)
			t1.Reset(time.Second * 5)
		case <-t2.C:
			println("10s ", i)
			t2.Reset(time.Second * 5)
		}
	}
}

func TestTimer3(t *testing.T) {
	input := make(chan interface{})

	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	i := 0
	for {
		i++
		select {
		case <-input:
			println("input")
		case <-t1.C:
			println("5s ", i)
			t1.Reset(time.Second * 5)
		case <-t2.C:
			println("10s ", i)
			t2.Reset(time.Second * 5)
		}
	}
}

func TestTicker01(t *testing.T) {
	ticker := time.NewTimer(time.Second * 2)
	quit := make(chan interface{})

	go func() {
		fmt.Println("start")
		for {
			select {
			case <-ticker.C:
				println("ticker")
			case <-quit:
				println("quit")
				ticker.Stop()
				return
			}
		}
		fmt.Println("stop")
	}()
	time.Sleep(10 * time.Second)

	quit <- 1
}

func TestTicker02(t *testing.T) {
	// 阻塞当前线程，等待所有任务完成
	var wg sync.WaitGroup
	wg.Add(2)
	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	//以释放相关资源。
	ticker1 := time.NewTicker(2 * time.Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1	", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(2 * time.Second)

	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer	", time.Now().Format("2006-01-02 15:04:05"))
			//Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时t
			//还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(2 * time.Second)
		}
	}(timer1)

	wg.Wait()
}

func TestTicker03(t *testing.T) {
	a1 := make([]map[string]string, 5)
	a2 := make([]map[string]string, 5)

	ticker1 := time.NewTicker(time.Minute * 1)
	go funcName(ticker1, a1, a2)
	<-ticker1.C
}

func funcName(ticker1 *time.Ticker, a1, a2 []map[string]string) {
	func(ticker1 *time.Ticker) {
		for {
			<-ticker1.C
		}
	}(ticker1)
}

func TestFuncName2(t *testing.T) {
	uc := make(chan int)
	go bbbb()
	go aaaa()
	<-uc
}

func aaaa() {
	c := cron.New()

	c.AddFunc("0/1 * * * *", func() { fmt.Println("-------------------------------") })
	c.Start()
	fmt.Println(len(c.Entries()))
	defer c.Stop()
	select {}
}
func bbbb() {
	c := cron.New()

	rst := &NoticeTwist{new: make([]*Notice, 0), old: make([]*Notice, 0)}
	_, _ = c.AddFunc("0/1 * * * *", RunNoticeTimers(rst))
	rst.old = SelectFromDBF()
	rst.old = SelectFromDBF()
	c.Start()
	fmt.Println(len(c.Entries()))
	defer c.Stop()
	select {}
}
func RunNoticeTimers(rst *NoticeTwist) func() {
	return func() {
		fmt.Println("oldlen ", len(rst.old))
		fmt.Println("newlen ", len(rst.new))

		fmt.Println("ticker start")
		for k, _ := range rst.old {
			fmt.Print(k)
		}
		fmt.Println()
		rst.new = SelectFromDBF()
		for k, v := range rst.old {
			ak := k
			av := v
			pre := v.NextAt.Unix()
			nowMillis := GetMillis()
			if pre-nowMillis < 0 {
				continue
			}
			println(v.Name, " add ", v.NextAt.Unix())

			time.AfterFunc(time.Millisecond*time.Duration(pre-nowMillis), func() {
				fmt.Println("已执行", rst.old[ak].Name, ak, av)
			})
		}
		rst.old = rst.new
	}
}

type NoticeTwist struct {
	old []*Notice
	new []*Notice
}

func SelectFromDBF() []*Notice {
	//db, err := sql.Open("mysql", "root:1234@/mattermost?charset=utf8")
	//if err != nil {
	//	println(err)
	//}
	//now := GetMillis()
	//s1 := TimestampToString(now + 3000000)
	//s2 := TimestampToString(now - 3000000)
	//stmt, err := db.Prepare("SELECT * FROM Notices where NextAt>" + s1 + " and NextAt<" + s2 + " and Status='running'")
	//if err != nil {
	//	println(err)
	//}
	//rs := make([]*Notice, 0)
	//for rows.Next() {
	//	var r *Notice
	//	err := rows.Scan(r)
	//	if err != nil {
	//		println(err)
	//	}
	//	rs = append(rs, r)
	//}
	//return rs
	return nil
}

func CreateDBFForNotices() {
	db, err := sql.Open("mysql", "root:1234@/mattermost?charset=utf8")
	if err != nil {
		println("err")
	}
	stmt, err := db.Prepare("INSERT INTO " +
		"`Notices` ( " +
		"`CreateAt`, `UpdateAt`, `ChannelId`, `Name`,`NextAt` ,`Frequency`, `Status`, `Remark`, `Mode`, `UserIds` " +
		")VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		println("err")
	}
	millis := GetMillis()
	int64s := []int64{millis - 3000000,
		millis - 2000000,
		millis - 1000000,
		millis - 100000,
		millis - 1000,
		millis - 100,
		millis - 10,
		millis - 1,
		millis + 1,
		millis + 10,
		millis + 100,
		millis + 1000,
		millis + 5000,
		millis + 10000,
		millis + 15000,
		millis + 20000,
		millis + 250000,
		millis + 250000,
		millis + 250000,
		millis + 300000,
		millis + 320000,
		millis + 320300,
		millis + 322300,
		millis + 322340,
		millis + 352340,
		millis + 372340,
		millis + 572340,
		millis + 672340,
		millis + 721340,
		millis + 972340,
		millis + 1000000,
		millis + 1000000,
		millis + 2000000,
		millis + 3000000}
	for _, v := range int64s {
		s := TimestampToString(v / 1000)
		println(s)
		_, err = stmt.Exec(s, s, "1", "sss", s, "weekly", "finish", "wwwaw", "inner", "[2,3]")
		_, err = stmt.Exec(s, s, "1", "bbbb", s, "dayly", "cancel", "qwwefqw", "inner", "[1,2]")
		_, err = stmt.Exec(s, s, "1", "dddd", s, "unrepeated", "running", "322br", "inner", "[1,3]")
		if err != nil {
			println(err.Error())
		}
	}
}

func TestDBTimer(t *testing.T) {
	CreateDBFForNotices()

	//go bbbb()
	//<-make(chan int)
}

func TestRedisGo(t *testing.T) {
	c := Pool.Get()
	defer c.Close()

	// 存入数据
	_, err := c.Do("SET", "kkk", "vvv")
	if err != nil {
		fmt.Println("err while setting:", err)
	}

	// 判断是否存在
	isExit, err := redis.Bool(c.Do("EXISTS", "kkk"))
	if err != nil {
		fmt.Println("err while checking keys:", err)
	} else {
		fmt.Println(isExit)
	}

	// 获取value并转成字符串
	value, err := redis.String(c.Do("GET", "kkk"))
	if err != nil {
		fmt.Println("err while getting:", err)
	} else {
		fmt.Println(value)
	}

	_, err = c.Do("DEL", "kkk")
	if err != nil {
		fmt.Println("err while deleting:", err)
	}

	_, err = c.Do("SET", "ke", "va", "EX", "5")
	if err != nil {
		fmt.Println("err while setting:", err)
	}

	_, err = c.Do("SET", "notices_old", SelectFromDBF())
	if err != nil {
		fmt.Println("err")
	}

	redis.MultiBulk(c.Do("GET", "notices_old"))

	_, err = c.Do("DEL", "notices_old")
	if err != nil {
		fmt.Println("err")
	}

}


type DataContainer struct {
	Queue chan interface{}
}

func NewDataContainer(max_queue_len int) (dc *DataContainer){
	dc = &DataContainer{}
	dc.Queue = make(chan interface{}, max_queue_len)
	return dc
}

//非阻塞push
func (dc *DataContainer) Push(data interface{}, waittime time.Duration) bool{
	click := time.After(waittime)
	select {
	case dc.Queue <- data:
		return true
	case <- click:
		return false
	}
}

//非阻塞pop
func (dc *DataContainer) Pop(waittime time.Duration) (data interface{}){
	click := time.After(waittime)
	select {
	case data =<-dc.Queue:
		return data
	case <- click:
		return nil
	}
}

//test
var MAX_WAIT_TIME = 10 *time.Millisecond
func TestNoBlock(t *testing.T){

	datacotainer := NewDataContainer(2)
	//add
	go funcName002(datacotainer)

	go funcName001(datacotainer)
	select{}
}
type dataItem struct {
	name string
	age int
}

func funcName002(datacotainer *DataContainer) {
	for {
		datacotainer.Push(&dataItem{"zhangsan", 25}, MAX_WAIT_TIME)
	}
}

func funcName001(datacotainer *DataContainer) {
	//get
	var item interface{}
	item = datacotainer.Pop(MAX_WAIT_TIME)
	if item != nil {
		if tmp, ok := item.(*dataItem); ok { //interface转为具体类型
			fmt.Printf("item name:%v, age:%v\n", tmp.name, tmp.age)
		}
	}
}