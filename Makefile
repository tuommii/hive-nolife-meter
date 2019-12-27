
BIN_DIR = bin

all: crawler parser

crawler:
	go build -o $(BIN_DIR)/crawler cmd/crawler/main.go

parser:
	go build -o $(BIN_DIR)/parser cmd/parser/main.go

deploy:
	scp -P PORT run.sh USER@IP:PATH
	scp -P PORT users.json USER@IP:PATH
	scp -P PORT bin/crawler USER@IP:PATH
	scp -P PORT bin/parser USER@IP:PATH
	scp -P PORT html/template.html USER@IP:PATH

test:
	go test cmd/parser/*.go

# cross-compile:
# 	GOOS=linux GOARCH=amd64 go build
