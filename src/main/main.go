package main

import (
	"fmt"
	"container/list"
	"sync"
	//"os"
	"time"
	//"runtime"
	"runtime"
)

func init() {
	fmt.Println("init fuc exec")
}

func typeChange(input int) float32 {
	return float32(input)
}

func main() {
	//var a int
	//a = 9
	//fmt.Println(a)
	//
	//b := int(99)
	//fmt.Println(b)
	//
	//intTest := int(98)
	//strTest := string("go land")
	//fmt.Println(intTest)
	//fmt.Println(strTest)
	//
	//array1 := []string{"111", "bb", "980"}
	//strings := append(array1, "abc")
	//fmt.Println(len(array1))
	//fmt.Println(strings)
	//fmt.Println("bbbbb")
	//
	//p := new(user)
	//p.name = "qiyz"
	//p.age = 12
	//fmt.Println(p)
	//
	//m := make(map[string]string)
	//m["a"] = "i"
	//m["c"] = "9"
	//fmt.Println(m)
	//
	//for key, val := range m {
	//	//fmt.Println("key=%s,value=%s", key, val)
	//	fmt.Printf("key=%s,value=%s\n", key, val)
	//}
	////sliceTest()
	//
	//rect3 := &Rect{0, 0, 100, 200}
	//fmt.Println(*rect3)
	//
	//rect4 := &Rect{width: 100, height: 200}
	//fmt.Println(*rect4)
	//
	//user1 := new(user)
	//user1.age = 78
	//user1.name = "ddd"
	//fmt.Println(user1)
	//
	//pid := os.Getgid()
	//ppid := os.Getppid()
	//fmt.Println(pid)
	//fmt.Println(ppid)
	//
	//fmt.Println(typeChange(99))

	exampleGo()

	//runtime.GOMAXPROCS(90)
}

func Contains(l *list.List, value string) (bool, *list.Element) {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return true, e
		}
	}
	return false, nil
}

func consolePrint(context string) {
	fmt.Println(context)
}

func sliceTest() {
	var ss []string;
	fmt.Printf("[ local print ]\t:\t length:%v\taddr:%p\tisnil:%v\n", len(ss), ss, ss == nil)
	print("func print", ss)
	//切片尾部追加元素append elemnt
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}
	fmt.Printf("[ local print ]\t:\tlength:%v\taddr:%p\tisnil:%v\n", len(ss), ss, ss == nil)
	print("after append", ss)
	//删除切片元素remove element at index
	index := 5;
	ss = append(ss[:index], ss[index+1:]...)
	print("after delete", ss)
	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], "inserted")
	ss = append(ss, rear...)
	print("after insert", ss)
}

func print(msg string, ss []string) {
	fmt.Printf("[ %20s ]\t:\tlength:%v\taddr:%p\tisnil:%v\tcontent:%v", msg, len(ss), ss, ss == nil, ss)
	fmt.Println()
}

type Rect struct {
	x, y          float64
	width, height float64
}

func NewRect(x, y, width, height float64) *Rect {

	return &Rect{x, y, width, height}

}

type user struct {
	lock sync.Mutex
	name string
	age  int
}

func testTimeerChannel() {
	//创建了一个定时器，2秒后会发送事件到timer1.C channel
	timer1 := time.NewTimer(time.Second * 2)

	//等待定时器到期
	data := <-timer1.C //接收到的数据 2016-07-16 15:24:19.337701998 +0800 CST
	fmt.Println("Timer 1 expired")
	fmt.Println("Timer 1 expired", data)
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	//关闭定时器
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func testChannel2() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//启动三个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	//循环9次，发送任务
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	//循环得到结果
	for a := 1; a <= 9; a++ {
		<-results
	}
}

func exampleGo() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello,  %s.\n", who)
		}(name)
	}
	runtime.Gosched()
}

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	// 将计算结果发送到channel中
	resultChan <- sum
}

func testChannel8() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 3)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	go sum(values[len(values)/3:], resultChan)
	sum1, sum2, sum3 := <-resultChan, <-resultChan, <-resultChan
	fmt.Println("Result:", sum1, sum2, sum3)
}

func testChannel9() {
	arrayChan := make(chan int, 20)
	arrayInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for t := 0; t < 10; t++ {
		length := len(arrayInt)
		go sum(arrayInt[length-t:], arrayChan)
	}

	arrayResult := [10]int{0}
	for i := 0; i < 10; i++ {
		arrayResult[i] = <-arrayChan
	}
	fmt.Println(arrayResult)
}
