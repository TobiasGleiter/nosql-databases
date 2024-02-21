# Redis Strings

## Basics

Redis strings store sequences of bytes, including text, serialized objects, and binary arrays. They're the simplest type of value associated with a Redis key, often used for caching, implementing counters, or performing bitwise operations. Here are some key commands:

- `SET key value`: Sets a string value, replacing existing data if the key already exists.
- `GET key`: Retrieves the value associated with a key.
- `SETNX key value`: Sets a string value only if the key doesn't already exist.
- `GETSET key value`: Sets a key to a new value, returning the old value.
- `MSET key value [key value ...]`: Sets multiple key-value pairs, reducing latency.
- `MGET key [key ...]`: Retrieves the values associated with multiple keys.

## Limits

A single Redis string can be a maximum of 512 MB by default.

## Uses Cases

- Caching HTML fragments or pages.
- Implementing counters for tracking events or statistics.
- Performing bitwise operations on binary data.

## Performance

Most string operations are O(1), ensuring high efficiency. However, operations like SUBSTR, GETRANGE, and SETRANGE can be O(n) and may impact performance with large strings.

For alternatives, consider Redis hashes or JSON for storing structured data as serialized strings.
