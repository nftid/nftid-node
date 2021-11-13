build:
	go build -o nftid-node main.go

run: build
	./nftid-node

watch:
	reflex -s -r '\.go$$' make run