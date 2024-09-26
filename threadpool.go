package main

import (
	"fmt"
	"sync"
	"time"
)

// operations here are assumed as strings which are to be printed
type Job struct {
	operations chan string
	wg         sync.WaitGroup // waits until all operations are done by a job
}

// initializes the channel
func (job *Job) NewJob() {
	job.operations = make(chan string)
}

// increments waitgroup counter and adds operation to the channel
func (job *Job) AddOperation(
	str string,
) {
	job.wg.Add(1)
	job.operations <- str
}

// checks for all the operations in the channel, executes it
// and decrements the waitgroup counter
func (job *Job) Execute() {
	for text := range job.operations {
		fmt.Println(text) // operation execution
		job.wg.Done()
	}
}

// pool is collection of more than one channel to handle operations
type ThreadPool struct {
	Jobs []Job
}

// initializes pool with specified number of channels
// and channels are concurrently waiting for some operation
// to come to it.
func (threadPool *ThreadPool) NewThreadPool(
	numberOfJobs uint,
) {
	threadPool.Jobs = make([]Job, numberOfJobs)
	for i := 0; i < int(numberOfJobs); i += 1 {
		threadPool.Jobs[i].NewJob()
		go threadPool.Jobs[i].Execute() // concurrently wait for some operation to come to channel and execute
	}
}

// waits for all the job to execute their respective tasks
func (threadPool *ThreadPool) Execute() {
	totalJobs := len(threadPool.Jobs)
	for i := 0; i < totalJobs; i += 1 {
		threadPool.Jobs[i].wg.Wait()
	}
}

// takes up operations all together and uniformly divides operations
// to different channels inside a pool
// and each operation addition is non-blocking so order is not guaranted.
func (threadPool *ThreadPool) AddOperations(
	operations []string,
) {
	totalJobs := len(threadPool.Jobs)
	for i := range len(operations) {
		go threadPool.Jobs[i%totalJobs].AddOperation(operations[i])
	}
}

func threadpool_main() {
	// declaring threadpool
	threadPool := ThreadPool{}
	// initializing pool with 2 channels
	threadPool.NewThreadPool(2)

	// infinite loop which keeps throwing operations at interval of a second
	for i := 0; i < 4; i -= 1 {
		threadPool.AddOperations([]string{"cheers!", "hey", "dil", "bekar", "sunna", "suzuki"})
		threadPool.Execute() // waits for all job to finish above set of operations
		time.Sleep(time.Second)
	}
}
