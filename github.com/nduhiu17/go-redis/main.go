package main

import (
	"fmt"
	"encoding/json"
	"github.com/go-redis/redis"

)



type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main()  {
	fmt.Println("Hello redis in go")
	// setting client
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()

	fmt.Println(pong,err)

	//adding values to redis

	err = client.Set("name","Nduhiu",0).Err()

	if err != nil {
		fmt.Println(err)
	}

	// retrieving values from redis
	val,err := client.Get("name").Result()
	fmt.Println("Name retrieved is: ",val)

	//adding composite values
	json, err := json.Marshal(Author{Name: "Nduhiu",Age: 31})

	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234",json,0).Err()

	if err != nil {
		fmt.Println(err)
	}

	val, err = client.Get("id1234").Result()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)


}