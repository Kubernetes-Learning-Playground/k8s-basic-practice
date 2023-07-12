package main

import (
	"math/rand"
	"time"
)

/*
	mkdir myapp

	cd myapp
	echo 1g  > memory.limit_in_bytes
	echo 0 > memory.swappiness


	然后进入myapp目录执行
	cgexec -g memory:myapp  ./myapp

*/

func mem() {
	var list []int
	for i := 0; i < 10000; i++ {
		list = append(list, rand.Int())
	}
	select {}
}
func main() {
	for{
		go mem()
		time.Sleep(time.Millisecond*200)
	}

}

