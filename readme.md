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
$ ./bin/server

# Shell 2
# client should be running after server initialized
$ ./bin/client
```

