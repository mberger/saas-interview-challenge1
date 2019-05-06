package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	//set
	fmt.Println("Publishing to Redis")

	c.Do("SET", "message1", "Hello ForgeRock")

	//sleep for a 100 ms
	fmt.Println("Sleeping for 5 seconds")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Here we go...")
	fmt.Println(" ")
	//get
	world, err := redis.String(c.Do("GET", "message1"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
	//ENDINIT OMIT
}
