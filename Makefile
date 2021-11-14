build:
	go build -o build/nftid-node cmd/nftid-node/main.go

run: build
	build/nftid-node

watch:
	reflex -s -r '\.go$$' make run