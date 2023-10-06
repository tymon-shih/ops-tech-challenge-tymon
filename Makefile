build:
    go build -o ops-tech-challenge

test:
    go test -v ./...

docker:
    docker build -t ops-tech-challenge .