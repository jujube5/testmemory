# Comparison In-Memory Caching

## Requirements

- Redis
- Redis client for Go

## Installation

```shell
go install github.com/jujube5/testmemorymain
```

```shell
go get github.com/jujube5/testmemory
```

## Quickstart

```shell
cd testmemorymain
```
```shell
go run main.go
```
### Output

```shell
Init rdb
 Redis Set 5 000 000 records 
Memory Available before test --  4433 MB
17:34:54
Memory Available after test --  4047 MB
17:36:23
 Redis Get 5 000 000 records 
Memory Available before test --  4047 MB
17:36:23
Memory Available after test --  3845 MB
17:37:56
 Redis Delete 5 000 000 records 
Memory Available before test --  3845 MB
17:37:56
Memory Available after test --  4287 MB
17:39:21
 Native Set 5 000 000 records 
Memory Available before test --  4287 MB
17:39:21
Memory Available after test --  3661 MB
17:39:24
 Native Get 5 000 000 records 
Memory Available before test --  3661 MB
17:39:24
Memory Available after test --  3653 MB
17:39:25
 Native Delete 5 000 000 records 
Memory Available before test --  3653 MB
17:39:25
Memory Available after test --  3623 MB
17:39:26
```
