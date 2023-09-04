clean:
	rm -rf pb build node_modules

go: clean
	mkdir pb
	protoc --go_out=./pb --go-grpc_out=./pb ./proto/*/*.proto ./proto/*/*/*.proto

ts: clean 
	yarn
	mkdir build
	protoc --plugin='protoc-gen-ts=./node_modules/.bin/protoc-gen-ts' --ts_out=./build proto/*/*.proto proto/*/*/*.proto