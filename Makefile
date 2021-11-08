build:
	go build -o ./bin/server github.com/AFukun/rpc-example/server
	go build -o ./bin/client github.com/AFukun/rpc-example/client

clean:
	rm -rf ./bin/*
