# Redis Lists

## Introduction to Redis Lists

Redis lists are linked lists of string values, commonly used to:

- Implement stacks and queues.
- Build queue management for background worker systems.

## Basic Commands

- `LPUSH element1 element2 ...`: Adds elements to the head of a list.
- `RPUSH element1 element2 ...`: Adds elements to the tail of a list.
- `LPOP`: Removes and returns an element from the head of a list.
- `RPOP`: Removes and returns an element from the tail of a list.
- `LLEN`: Returns the length of a list.
- `LMOVE`: Atomically moves elements from one list to another.
- `LTRIM`: Reduces a list to the specified range of elements.

## Blocking Commands

Redis lists support several blocking commands, such as:

- `BLPOP timeout key1 key2 ...`: Removes and returns an element from the head of a list, blocking until an element becomes available or the timeout is reached.
- `BLMOVE source destination LEFT RIGHT timeout`: Atomically moves elements from a source list to a target list, blocking until a new element becomes available in the source list or the timeout is reached.

## Examples

### Treating a List like a Queue (First In, First Out)

```
LPUSH bikes:repairs bike:1
LPUSH bikes:repairs bike:2
RPOP bikes:repairs
RPOP bikes:repairs
```

### Treating a List like a Stack (First In, Last Out)

```
LPUSH bikes:repairs bike:1
LPUSH bikes:repairs bike:2
LPOP bikes:repairs
LPOP bikes:repairs
```

### Checking the Length of a List

```
LLEN bikes:repairs
```

### Atomically Popping an Element from One List and Pushing to Another

```
LPUSH bikes:repairs bike:1
LPUSH bikes:repairs bike:2
LMOVE bikes:repairs bikes:finished LEFT LEFT
LRANGE bikes:repairs 0 -1
LRANGE bikes:finished 0 -1
```

### Limiting the Length of a List

```
RPUSH bikes:repairs bike:1 bike:2 bike:3 bike:4 bike:5
LTRIM bikes:repairs 0 2
LRANGE bikes:repairs 0 -1
```

## What are Lists?

Lists are implemented using linked lists in Redis, allowing for fast addition of elements to very long lists. However, accessing elements by index is not as fast as with arrays.

## First Steps with Redis Lists

```
RPUSH bikes:repairs bike:1 bike:2
LPUSH bikes:repairs bike:important_bike
LRANGE bikes:repairs 0 -1
```

## Common Use Cases for Lists

- Remembering the latest updates in social networks.
- Communication between processes using a producer-consumer pattern.

## Capped Lists

Lists can be capped to a certain length using LTRIM, allowing for efficient management of memory.

## Blocking Operations on Lists

Blocking commands like BRPOP and BLPOP make Redis lists suitable for implementing queues.

## Automatic Creation and Removal of Keys

Redis automatically creates or removes keys associated with lists as needed.

## Limits

The max length of a Redis list is 2^32 - 1 (4,294,967,295) elements.

## Performance

List operations accessing head or tail are O(1), but manipulating elements within a list can be O(n).

## Alternatives

Consider Redis streams as an alternative to lists when dealing with an indeterminate series of events.
