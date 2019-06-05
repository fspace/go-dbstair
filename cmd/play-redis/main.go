package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-config"
	"github.com/go-redis/redis"
)

func main() {
	// create a Config object
	c := config.New()

	// load configuration from a JSON string
	c.Load("app.json")

	// get the "Version" value, return "1.0" if it doesn't exist in the config
	version := c.GetString("Version", "1.0")

	var author struct {
		Name, Email string
	}
	// populate the author object from the "Author" configuration
	c.Configure(&author, "Author")

	fmt.Println(version)
	fmt.Println(author.Name)
	fmt.Println(author.Email)
	// Output:
	// 2.0
	// Foo
	// bar@example.com

	ExampleNewClient()
	ExampleClient()
}

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
func ExampleClient() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// copy from above method !

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
