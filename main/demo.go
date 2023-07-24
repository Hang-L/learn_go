package main

import (
	"fmt"
	"time"
)
// func main() {
// 	// var l list.List
// 	// l.PushBack(1)
// 	// l.PushBack(2)
// 	// l.PushBack(3)
// 	// l.PushBack(4)
// 	// l.PushBack(5)
// 	// element := l.Front()
// 	// for i := 0; i <= l.Len(); i++ {
// 	// 	fmt.Println(i, l)
// 	// }
// 	// lenth := l.Len()
// 	// fmt.Println(element.Value)
// 	// fmt.Println(lenth)
// 	// fmt.Printf("hello word")
// 	//chan_test1()
// 	// cat := Cat{
// 	// 	name: "aaa",
// 	// 	age:  10,
// 	// 	sex:  "M",
// 	// }
// 	// cat.SetName("bbbb")
// 	// fmt.Println(cat)
// 	// cat.SetName1("eee")
// 	// fmt.Println(cat)
// 	///添加注释测试
// 	for {

// 		resp, err := http.Get("https://gw-boki.starpay.jp/wxpos/version")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(resp)

// 	}

// }

func chan_test() {
	chan1 := make(chan int, 3)
	chan1 <- 1
	chan1 <- 3
	chan1 <- 2
	chan1 <- 4
	elem := <-chan1
	fmt.Println("The first element received from channel ch1:%v\n", elem)
}

func chan_test1() {
	intChan := make(chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("the canditate case is closed")
			break
		}
		fmt.Println("The canditate case is selected")
	}

}

type Cat struct {
	name string
	age  int
	sex  string
}

func (cat Cat) SetName(name string) {
	cat.name = name
}

func (cat *Cat) SetName1(name string) {
	cat.name = name
}
