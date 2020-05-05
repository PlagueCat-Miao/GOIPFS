all :
	make cloud
	make edge
	make gateway

cloud :

edge :

gateway :

MySql :
	cd $GOPATH/src/github.com/\!plague\!cat-\!miao/GOIPFS/Db/MySql
	go build -o ../../output/MySql.exe  MySqlUp.go
	cd ../..
clean :
	rm -rf ./output