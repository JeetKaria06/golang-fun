package main

import (
	"fmt"
	"time"
)

const iters = 5000
const ms1 = time.Millisecond

func main() {

	c := Config{Readers: 5, Writers: 1, ReadPause: ms1, WritePause: ms1}
	mutexesDuration := c.Run(iters)
	fmt.Println("configs: ", c, "mutexesDuration: ", mutexesDuration)

	c.Readers = 4
	c.Writers = 2
	mutexesDuration = c.Run(iters)
	fmt.Println("configs: ", c, "mutexesDuration: ", mutexesDuration)

	c.Readers = 3
	c.Writers = 3
	mutexesDuration = c.Run(iters)
	fmt.Println("configs: ", c, "mutexesDuration: ", mutexesDuration)

	c.Readers = 2
	c.Writers = 4
	mutexesDuration = c.Run(iters)
	fmt.Println("configs: ", c, "mutexesDuration: ", mutexesDuration)

	c.Readers = 1
	c.Writers = 5
	mutexesDuration = c.Run(iters)
	fmt.Println("configs: ", c, "mutexesDuration: ", mutexesDuration)
}
