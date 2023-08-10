gen:
# 	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
# gen:
# 	protoc proto/*.proto \
# 			--go_out=pb \
# 			--go_opt=module=github.com/rob-bender/go-test \
# 			proto/*.proto
# gen:
# 	protoc proto/*.proto \
# 			--go_out=pb \
# 			--go_opt=paths=source_relative \
# 			--proto_path=.
	protoc --proto_path=proto proto/*.proto --go_out=.

clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...