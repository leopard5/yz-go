package main

import "fmt"

func init()  {
	fmt.Println("init fuc exec")
}

func main() {
	go consolePrint("Hello Go Lang")
	// fmt.Println("bbbbb")
}

func consolePrint(context string) {
	fmt.Println(context)
}
