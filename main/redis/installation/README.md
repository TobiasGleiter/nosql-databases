# 👑 Install Redis

This is an installation Guide and shows how to install Redis on Linux, macOS, and Windows.

## Install Redis

This installation guide is similar to the Guide from [Redis](https://redis.io/docs/install/install-redis/).

The Redis installation depends on your operating system, pick your system:

- [Install Redis from Source](./source/README.md)
- [Install Redis on Linux](./linux/README.md)
- [Install Redis on macOS](./macOS/README.md)
- [Install Redis on Windows](./windows/README.md)

## Test if you can connect using the CLI

After you have Redis up and running (`redis-server`) you can connect using the **redis-cli**.

The **redis-cli** makes it easy to interact with the database.

The first thing after the installation is to check if Redis is working proberly by running:

```
redis-cli ping
```

The response should be `PONG`. Use `redis-cli --help` for more information.

A different approach to use redis is to type `redis-cli` without any arguments. The terminal will show the ip adress and port of the Redis-Server.
