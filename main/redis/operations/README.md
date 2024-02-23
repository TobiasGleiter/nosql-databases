# Operations

# Working with the Redis cli

### Command Line Usage

It is easy to use Redis with the command line for example testing purposes or while developing.
First make sure redis server is running (rin in terminal `redis-server`).
If it is running connect to the redis-cli via `redis-cli` in your terminal.

The command `INCR mycounter` does increase the value of the key (which is mycounter). The output after the first execution of the command is `1`. Human readable output is only enabled in the tty or terminal. For all other outputs it will enable the _raw output mode_.

### String quoting and escaping

In the redis-cli whitespace characters automatically delimit the arguments.
If there is the need for whietsprache between words use double (") or single (') quotation marks.
An escape sequence contains backslash (\) symbol followed by one of the escape sequence characters.

Doubly-quoted strings support the following escape sequences:

- \" - double-quote
- \n - newline
- \r - carriage return
- \t - horizontal tab
- \b - backspace
- \a - alert
- \\ - backslash
- \xhh - any ASCII character represented by a hexadecimal number (hh)

Single quotes assume the string is a letral, and allow only the following esacpe sequences

- \' - single quote
- \\ - backslash

### Host, port, password and database

- Use the `-h` flag to change host name or IP adress
- Use the `-p` flag to change the port.

Example: `redis-cli -h redis15.localnet.org -p 6390`

It is also possible to connect using one uri:

`redis-cli -u redis://LJenkins:p%40ssw0rd@redis-16379.hosted.com:16379/0`

### Mass insertion of data using redis-cli

#### CSV output

Export data from into a CSV file using the following commands:

```
redis-cli LPUSH mylist a b c d
```

The output is: (integer) 4

```
redis-cli --csv LRANGE mylist 0 -1
```

The output is: "d","c","b","a"

## Redis data types

Redis supports different types of data e.g. json, lists, string. IN the folloing we explain three used data types:

- [strings](./strings/README.md)
- [json](./json/README.md)
- [lists](./lists/README.md)

Other data types are sets, hashes, sorted sets, streams, geospatial, bitmaps, bitfields, probabilistic, and time series.
Find more about them on the [official documentation](https://redis.io/docs/data-types/)

## Redis commands

Find all commands on the [official website](https://redis.io/commands/)
