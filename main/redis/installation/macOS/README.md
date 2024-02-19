To install and manage Redis on macOS using Homebrew, follow these steps:

**Prerequisites:**

Ensure Homebrew is installed. Check by running:

```bash
brew --version
```

If not installed, follow Homebrew installation instructions.

**Installation:**

Run in Terminal:

```bash
brew install redis
```

This installs Redis on your system.

**Starting and Stopping Redis:**

To test your Redis installation, run:

```bash
redis-server
```

If successful, you'll see startup logs.

To stop Redis, press `Ctrl-C`.

**Using launchd to Manage Redis:**

Alternatively, start Redis in the background using launchd:

```bash
brew services start redis
```

This also sets Redis to restart at login. Check status with:

```bash
brew services info redis
```

To stop the service:

```bash
brew services stop redis
```

**Connect to Redis:**

After installation, test by running:

```bash
redis-cli
```

This opens the Redis REPL. Try commands like:

```bash
127.0.0.1:6379> lpush demos redis-macOS-demo
OK
127.0.0.1:6379> rpop demos
"redis-macOS-demo"
```

Redis is now ready for use on your macOS system.
