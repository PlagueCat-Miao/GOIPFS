OS_NAME = $(shell go env GOOS)
all :
	make cloud
	make edge
	make gateway

cloud :

edge :

gateway :

MySql :
	cd $GOPATH/src/github.com/\!plague\!cat-\!miao/GOIPFS/Db/MySql; \
	go build -o ../../output/MySql_$(OS_NAME)  MySqlUp.go

clean :
	rm -rf ./output