FILENAME="awskeys"

.PHONY: clean
clean:
	rm -f $(FILENAME)

.PHONY: build
build: clean
	go build -o $(FILENAME) main.go

.PHONY: install
install:
	go build -o $(GOPATH)/bin/$(FILENAME) main.go

.PHONY: winrelease
winrelease:
	GOOS=windows GOARCH=amd64 go build -o $(FILENAME).exe main.go

.PHONY: macrelease
macrelease:
	GOOS=darwin GOARCH=amd64 go build -o $(FILENAME)-mac main.go

.PHONY: linuxrelease
linuxrelease:
	GOOS=linux GOARCH=amd64 go build -o $(FILENAME)-linux main.go

.PHONY: buildall
buildall: winrelease macrelease linuxrelease