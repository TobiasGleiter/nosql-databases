package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9" // 2. Import the Redis package
)

func main() {
    // 3. Connect to the Redis client
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // Create a context
    ctx := context.Background()

    // Set a key-value pair in Redis
    err := client.Set(ctx, "foo", "bar", 0).Err()
    if err != nil {
        panic(err) // If there's an error, panic
    }

    // Get the value of the "foo" key from Redis
    val, err := client.Get(ctx, "foo").Result()
    if err != nil {
        panic(err) // If there's an error, panic
    }
    
    // Print the value retrieved from Redis
    fmt.Println("foo", val)
}
