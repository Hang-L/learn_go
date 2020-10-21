package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer1 bytes.Buffer
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Waiting contents %q ....\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer :%d\n", buffer1.Len())
	fmt.Printf("The capcity of buffer :%d\n", buffer1.Cap())
	p1 := make([]byte, 7)
	n, _ := buffer1.Read(p1)
	fmt.Printf("%d bytes were read.(call Read)", n)
	fmt.Printf("The length of buffer :%d\n", buffer1.Len())
	fmt.Printf("The capcity of buffer :%d\n", buffer1.Cap())

	//内容泄露
	contents = "ab"
	buffer2 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q:%d\n", contents, buffer2.Cap()) //内容容器的容量为：8
	unreadBytes := buffer2.Bytes()
	fmt.Printf("The unread bytes of the buffer:%v\n", unreadBytes) //读取结果为[97 98]
	buffer2.WriteString("cdefg")
	fmt.Printf("The capacity of  buffer: %d\n", buffer2.Cap()) //底层数组不变，内容容器的容量为：8
	//只需扩充一下unreadbytes就可以读取甚至修改buffer中的后续内容
	// unreadBytes = unreadBytes[:cap(unreadBytes)]
	// fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes) //读取结果为[97 98 99 100 101 102 103 0]，泄露内容

	// ///更改buffer2的值
	// unreadBytes[len(unreadBytes)-2] = byte('X') //'X'的ASCII码为88
	// fmt.Printf("The unread bytes of the buffer: %v\n", buffer2.Bytes())

	//buffer内容容器真正扩充之后就不会存在泄露问题了
	contents = "hijklmn"
	buffer2.WriteString(contents)
	fmt.Printf("The capacity of  buffer: %d\n", buffer2.Cap()) //扩容了内容容器
	///更改buffer2的值
	unreadBytes = unreadBytes[:cap(unreadBytes)] //'X'的ASCII码为8
	unreadBytes[len(unreadBytes)-2] = byte('X')  //'X'的ASCII码为8
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer2.Bytes())
}
