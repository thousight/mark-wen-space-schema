clean:
	rm -rf pb build node_modules

go: clean
	mkdir pb
	protoc -I. --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ./proto/*/*.proto ./proto/*/*/*.proto

ts: clean 
	yarn
	mkdir build
	protoc --plugin='protoc-gen-ts=./node_modules/.bin/protoc-gen-ts' --ts_out=./build proto/*/*.proto proto/*/*/*.proto