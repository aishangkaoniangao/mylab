package main

import "fmt"
import "strconv"

var ch = make(chan string, 10)

func main() {
	fmt.Println("==start==")

	var produce = func(i int) {
		ch <- "test"
		fmt.Println("channel " + strconv.Itoa(i) + " is producing...")
	}
	//输入100次,因为chan长度为10，所以每次只能送进去10个元素，等消费进程处理后，后再送进去10个元素，依次往复
	//这种操作类似于写日志，写日志用一个进程，产生日志会有多个进程，如果队列太短，则只有等待写日志进程把缓冲区的内容刷到磁盘，产生日志进程才能继续执行，这样会把正常请求hang住在写日志的点，所以日志缓冲队列设置一定得长一些才能避免这个问题
	for i := 0; i < 100; i++ {
		go produce(i)
	}

	var input string
	var num int

	for {
		if num >= 20 {
			break
		}
		input = <-ch
		fmt.Println("out:", input)
		num++
	}

	fmt.Println("==over==")
}
