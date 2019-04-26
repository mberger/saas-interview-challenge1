package main

import (
	"fmt"
	"math/rand"
	"time"
        "github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
	panic(err)
	}
	defer c.Close()


	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// create 3 Workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Give them jobs j
	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	// make sure to close jobs
	close(jobs)

	// Wait for the results r
	for r := 1; r <= 6; r++ {
		fmt.Println("The result received from worker is: ", <-results)
	}
}

func worker(ID int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("The worker with ID ", ID, " is working on the job with the number = ", job)
		duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
		time.Sleep(duration)
		fmt.Println("The worker with ID ", ID, " has completed work on job with the number ", job, " within ", duration)
		results <- ID
	}

}

func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second) // <------------ here
}

func delayMinute(n time.Duration) {
	time.Sleep(n * time.Minute)
}
