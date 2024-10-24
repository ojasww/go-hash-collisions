## Hash Collisions

This repository contains **Go** code to handle collisions in hash maps through [Separate Chaining](https://en.wikipedia.org/wiki/Hash_collision#Separate_chaining).

The data structure used for hash maps is an array (**buckets**) of Linked Lists. Caller function can **Insert**, **Get**, **Delete** the entries in the hash map. Caller can also **Rehash** the hash map which doubles it's size keeping all the key - value pairs intact in the map.

## Testing

`git clone` this repository and run the following command to test the individual functions mentioned above: 

```go
    go test -v -count=1 ./...
```