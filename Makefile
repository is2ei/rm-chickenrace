NAME := rm

bin/${NAME}:
	GOOS=linux GOARCH=amd64 go build -o bin/$(NAME)
	cp -a bin vagrant/

.PHONY: clean deps build run test

clean:
	rm -rf bin/*
	rm -rf tmp

test:
	mkdir -p tmp/a/a1/a2/a3 tmp/b tmp/c/d/e
	touch tmp/aa tmp/a/aa tmp/a/a1/a2/aaaaaa tmp/b/bb tmp/c/d/e/ffff
	dd if=/dev/zero of=tmp/a/a1/data bs=1029 count=1
	@make run ARG="tmp/a*"
	rm -rf tmp

test2:
	mkdir -p tmp/DD
	touch tmp/DD/aaa
	sudo chown root:wheel tmp/DD
	@make run ARG="tmp/DD"
	sudo rm -rf tmp

run-jq:
	go run *.go ${ARG} | jq .

run:
	go run *.go ${ARG}