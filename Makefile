build:
	go build -o ml ./mylinter/cmd/mylinter

test:
	go vet -vettool ml ./...

clean:
	rm -f ml