# Operations

# Working with the Redis cli

It is easy to use Redis with the command line for example testing purposes or while developing.
First make sure redis server is running (rin in terminal `redis-server`).
If it is running connect to the redis-cli via `redis-cli` in your terminal.

The command `INCR mycounter` does increase the value of the key (which is mycounter). The output after the first execution of the command is `1`. Human readable output is only enabled in the tty or terminal. For all other outputs it will enable the _raw output mode_.

# Redis Commands

There are some commands that are useful for the database interaction:

| Command | Description                    | Function                                                         |
| ------- | ------------------------------ | ---------------------------------------------------------------- |
| SET     | Set key                        | Set values such as strings and integers by linking them to keys. |
| GET     | Get key value                  | Get the value associated with the key                            |
| DEL     | delete key                     | Delete one or more keys                                          |
| EXPIRE  | Set expiration date            | Set auto-deletion expiration time (TTL) for keys                 |
| INCR    | increment value                | Increase integer value by 1                                      |
| DECR    | Decrement value                | Decrease integer value by 1                                      |
| LPUSH   | Add to top of list             | Add one or more values to the beginning of the list              |
| RPOP    | Get from end of list           | Get the last value of the list and remove it from the list       |
| LLEN    | Get list length                | Get the number of values in a list                               |
| LINDEX  | Get an element at any position | Get the element at any index of the list                         |
| LRANGE  | Get range of list              | Get elements in the specified range of the list                  |
| LSET    | Update elements in list        | Update the element at the specified index in the list            |
| LREM    | Delete elements in list        | Delete element with specified value in list                      |
| SADD    | Add to set                     | Add unique values to the set                                     |
| ZADD    | Add to sorted set              | Associate score and value and add to the sorted set              |
| ZRANGE  | Get range of sorted set        | Get elements in a specified range in sorted order                |
| HSET    | set field to hash              | Set value in the specified field of hash key                     |
| HGET    | Get field value of hash        | Get the value of the specified field of the hash key             |
| HGETALL | Get all fields of hash         | Get all fields and values included in a hash key                 |
| HDEL    | Remove field from hash         | Delete hash key specified field                                  |

Find all commands on the [official website](https://redis.io/commands/)
