# Install Redis from Source

Redis can be installed from source on variety of plattforms and operation systems inluding Linux and macOS.

## Installation process

1. Downloading the source files
2. Compiling Redis
3. Starting and stopptin Redis n the foreground

## 1. Downloading the source files

Download the source files from Redis:

`wget https://download.redis.io/redis-stable.tar.gz`

## 2. Compiling Redis

After run the tarbal, change the root directory, and then run **make**:

```
tar -xzvf redis-stable.tar.gz
cd redis-stable
make
```

If TLS necessary you need to installe the OpenSSL development libaries and then run:

`make BUILD_TLS=yes`

If the compile succeeds ✅ you will find several Redis bianries in the src directory, including:

- **redis-server**: the Redis Server itself
- **redis-cli**: is the command line interface utility to talk with Redis

To instal these binaries in `/usr/local/bin`, run:

`sudo make install`

## 3. Starting and stopping Redis in the foreground

Start Redis by running:

`redis-server`

If succesful, you will see the startup logs for Redis.
Stop Redis with `Ctrl-C`
