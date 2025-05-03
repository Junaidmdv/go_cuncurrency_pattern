package main

import (
	"fmt"
	"sync"
)


type job struct{
	Id  int
	Value int
}

type Result struct{
	Id int 
	Square int
}

func worker(w int,jobs <-chan job,result chan<- Result,wg *sync.WaitGroup){
    defer wg.Done()
    for job:=range jobs{
       result <-Result{Id: job.Id,Square: job.Value*job.Value}
	}

}
func FinalResults(resWg *sync.WaitGroup, results <-chan Result){
    defer resWg.Done()

	for result:=range results{
		 fmt.Println(result)
	}


}




func dispacture(jobcount int ,workers_count int){
	jobs:=make(chan job,jobcount)
	results:=make(chan Result,jobcount)
	 
	var wg sync.WaitGroup

	wg.Add(workers_count)

	for w:=1;w<=workers_count;w++{
		 go worker(w,jobs,results,&wg)
	}
    
	var resWg sync.WaitGroup
	resWg.Add(1)

	go FinalResults(&resWg,results)

	for i:=1;i<=jobcount;i++{
		jobs <- job{Id: i,Value: i}
	}
	close(jobs)
	wg.Wait()
	close(results)
	resWg.Wait()
}

func main(){
	const (
		workers_count=10
		jobs_count=100

	)

	fmt.Println("starting batch collection processing along with syncronisation")

	dispacture(jobs_count,workers_count)
}