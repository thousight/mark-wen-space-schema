clean:
	rm -rf markwen.space build node_modules

go: clean
	protoc -I. --go_out=. --go-grpc_out=. ./proto/*/*.proto ./proto/*/*/*.proto

ts: clean 
	yarn
	mkdir build
	protoc --plugin='protoc-gen-ts=./node_modules/.bin/protoc-gen-ts' --ts_out=./build proto/*/*.proto proto/*/*/*.proto