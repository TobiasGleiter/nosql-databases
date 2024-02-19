# 🚀 Use Redis with Go

## 📋 Checklist

1. Setup your go project `go mod init github.com/my/repo`
2. Run `go get github.com/redis/go-redis/v9` to install go for redis
3. Connect to Redis using this code:

```Go
import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
}
```

4. Store and retrieve simple string

```Go
ctx := context.Background()

err := client.Set(ctx, "foo", "bar", 0).Err()
if err != nil {
    panic(err)
}

val, err := client.Get(ctx, "foo").Result()
if err != nil {
    panic(err)
}
fmt.Println("foo", val)
```
