build:
	protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=.  proto/consignment/consignment.proto
	sudo docker build -t shippy-service-consignment .

run: 
	sudo docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment