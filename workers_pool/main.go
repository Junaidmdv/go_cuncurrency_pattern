package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, jobchan chan int, reschan chan int) {
	defer wg.Done()

	for job := range jobchan {
		reschan <- job * job
	}
}

func FinalResult(wg *sync.WaitGroup, reschan chan int) {

	defer wg.Done()

	for results := range reschan {
		fmt.Println(results)
	}

}

func processJob(workers, jobs int) {

	jobChan := make(chan int, jobs)
	resChan := make(chan int, jobs)

	var wg sync.WaitGroup

	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go worker(&wg, jobChan, resChan)
	}

	for i := 1; i <= jobs; i++ {
		jobChan <- i
	}
	close(jobChan)

	var resWg sync.WaitGroup
	resWg.Add(1)
	go FinalResult(&resWg, resChan)

	wg.Wait()
	close(resChan)
	resWg.Wait()
}

func main() {

	workers := 5
	jobs := 100

	processJob(workers, jobs)

}
