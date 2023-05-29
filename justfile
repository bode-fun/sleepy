build:
    mkdir -p ./dist
    go build -ldflags="-s -w" -o dist/ .

install: build
    mv ./dist/sleepy ~/.local/bin/