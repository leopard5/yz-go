package main

import (
	"fmt"
	"container/list"
)

type user struct {
	name string // name
	age  int    // age
}

func init() {
	fmt.Println("init fuc exec")
}

func main() {

	var a int
	a = 9
	fmt.Println(a)

	b := int(99)
	fmt.Println(b)

	intTest := int(98)
	strTest := string("go land")
	fmt.Println(intTest)
	fmt.Println(strTest)

	array1 := []string{"111", "bb", "980"}
	strings := append(array1, "abc")
	fmt.Println(len(array1))
	fmt.Println(strings)
	fmt.Println("bbbbb")

	p := new(user)
	p.name = "qiyz"
	p.age = 12
	fmt.Println(p)

	m := make(map[string]string)
	m["a"] = "i"
	m["c"] = "9"
	fmt.Println(m)

	for key, val := range m {
		//fmt.Println("key=%s,value=%s", key, val)
		fmt.Printf("key=%s,value=%s\n", key, val)
	}
	//sliceTest()

	rect3 := &Rect{0, 0, 100, 200}
	fmt.Println(*rect3)

	rect4 := &Rect{width: 100, height: 200}
	fmt.Println(*rect4)

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
