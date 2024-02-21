# Redis JSON

## Basics

Redis JSON allows for storing, updating, and retrieving JSON values within Redis. Key features include:

- Full support for the JSON standard.
- JSONPath syntax for selecting/updating elements inside documents.
- Documents stored as binary data in a tree structure for fast access to sub-elements.
- Typed atomic operations for all JSON value types.

Here are some basic commands to get started:

- `JSON.SET key path value`: Sets a Redis key with a JSON value.
- `JSON.GET key path`: Retrieves the JSON value associated with a key.
- `JSON.TYPE key path`: Returns the type of the value at the specified path within the JSON document.
- `JSON.STRLEN key path`: Returns the length of a string value within the JSON document.
- `JSON.NUMINCRBY key path value`: Increments a numeric value by the specified amount.
- `JSON.NUMMULTBY key path value`: Multiplies a numeric value by the specified amount.
- `JSON.SET obj $ '{"name":"John","age":30}'`: Example of setting a JSON object.
- `JSON.ARRAPPEND arr $ "new_element"`: Example of appending to a JSON array.

## Limits

- JSON values passed to commands are limited to a depth of up to 128 levels. Beyond this, commands return an error.

## Use Cases

- Storing and manipulating structured data within Redis.
- Integrating JSON data with Redis' search and query capabilities.

## Performance

- Most JSON operations are optimized for speed, with typical operations being O(1).
- Operations like JSONPath traversal or manipulation of deeply nested structures may have higher complexity (e.g., O(n)).

## Indexing and Searching JSON Documents

Redis JSON seamlessly integrates with Redis' Search and Query capabilities, allowing indexing and querying of JSON documents.

## Developer Notes

- Debugging, testing, and documentation notes for Redis JSON.

This breakdown should help you understand the basic operations, limitations, use cases, and performance considerations when working with JSON data in Redis.
