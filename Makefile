test:
	go build -o ml ./mylinter/cmd/mylinter
	go vet -vettool ml ./...