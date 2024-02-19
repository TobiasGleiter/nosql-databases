# Installing Redis from Source

Redis can be installed from source on various platforms like Linux and macOS.

## Installation Steps:

1. **Download the Source Files:**

   - Retrieve the source files from Redis:
     ```
     wget https://download.redis.io/redis-stable.tar.gz
     ```

2. **Compile Redis:**

   - Extract the files and compile Redis:

     ```
     tar -xzvf redis-stable.tar.gz
     cd redis-stable
     make
     ```

   - For TLS support, install OpenSSL development libraries and compile with:

     ```
     make BUILD_TLS=yes
     ```

   - After successful compilation, Redis binaries will be available in the `src` directory, including:

     - `redis-server`: Redis Server
     - `redis-cli`: Command-line interface for Redis

   - To install these binaries in `/usr/local/bin`, use:
     ```
     sudo make install
     ```

3. **Starting and Stopping Redis:**

   - Start Redis in the foreground:

     ```
     redis-server
     ```

   - To stop Redis, press `Ctrl-C`.
