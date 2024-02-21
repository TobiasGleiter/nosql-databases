# Redis Strings

## Basics

Strings can be set different ways. The following is a list of the commands:

- `SET bike:1 Deimos` (set but also replaces existing data)
- `GET bike:1`(gets the value with the key bike:1)
- `set bike:1 bike nx` (is the same as SET but returns nil if the key exists)
- `set bike:1 bike xx`(the set only succeeds if the key exists)
- `mset bike:1 "Deimos" bike:2 "Ares" bike:3 "Vanth"` (set multiple key-value pairs which reduces latency)
- `mget bike:1 bike:2 bike:3` (returns an array of value)

## Limits

By default, a single Redis sing can be maximum of 512 MB

## Uses cases

Caching html fragments or pages.

## Performance

Most string operations are O(1).
Be careful when using the SUBSTR, GETRANGE, and SETRANGE commands, which can be O(n).
