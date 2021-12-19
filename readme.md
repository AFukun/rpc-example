# Client-Server Example Using gRPC

## Build Excutable

```shell
$ mkdir -p $GOPATH/src/github.com/AFukun
$ cd $GOPATH/src/github.com/AFukun
$ git clone https://github.com/AFukun/rpc-example.git
$ cd rpc-example
$ go mod tidy
$ make build
```

## Run Example

```shell
# Shell 1 
$ ./bin/server 9000
# Shell 2
$ ./bin/server 9001
# Shell 3
$ ./bin/server 9002
# Shell 4
$ ./bin/server 9003

# Shell 0 (client should be running after server initialized)
# 只要有一个server正确运行则可以返回结果
$ ./bin/client
```

