build: 
	protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=.   proto/vessel/vessel.proto
	sudo docker build -t shippy-service-vessel .

run:
	sudo docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-vessel

