compile:
	mkdir -p bin/
	@echo "Compiling for every OS and Platform"

	@echo "🐧 Compile for Linux"
	GOOS=linux GOARCH=amd64 go build -o ./bin/d2dc-linux-amd64 ./cmd/d2dc
	GOOS=linux GOARCH=386 go build -o ./bin/d2dc-linux-386 ./cmd/d2dc
	GOOS=linux GOARCH=arm go build -o ./bin/d2dc-linux-arm ./cmd/d2dc
	GOOS=linux GOARCH=arm64 go build -o ./bin/d2dc-linux-arm64 ./cmd/d2dc

	@echo "🍏 Compile for Apple"
	GOOS=darwin GOARCH=amd64 go build -o ./bin/d2dc-darwin-amd64 ./cmd/d2dc

	@echo "🪟 Compile for Windows"
	GOOS=windows GOARCH=amd64 go build -o ./bin/d2dc-windows-amd64 ./cmd/d2dc
	GOOS=windows GOARCH=386 go build -o ./bin/d2dc-windows-386 ./cmd/d2dc

	@echo "🐡 Compile for FreeBSD"
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/d2dc-freebsd-amd64 ./cmd/d2dc
	GOOS=freebsd GOARCH=386 go build -o ./bin/d2dc-freebsd-386 ./cmd/d2dc
	GOOS=freebsd GOARCH=arm go build -o ./bin/d2dc-freebsd-arm ./cmd/d2dc