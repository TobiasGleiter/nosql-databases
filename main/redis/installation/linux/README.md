# Installing Redis on Linux

Redis can be installed on Linux using different methods depending on your distribution or preferences.

## Installing from Official Repository (Ubuntu/Debian)

For Ubuntu/Debian systems, you can install Redis from the official APT repository.

### Prerequisites:

If you're using a minimal distribution, ensure you have the necessary tools installed:

```bash
sudo apt install lsb-release curl gpg
```

### Installation Steps:

1. Add the Redis repository to APT:

```bash
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list
```

2. Update APT:

```bash
sudo apt-get update
```

3. Install Redis:

```bash
sudo apt-get install redis
```

## Installing via Snapcraft

Alternatively, Redis can be installed using Snapcraft, a package management system.

1. Install Redis via Snap:

```bash
sudo snap install redis
```

If Snap isn't installed on your system, you can follow the instructions to install it separately.

Once installed, Redis is ready to use.
