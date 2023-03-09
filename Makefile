mini: mini.go
ifdef VERSION
	go build -ldflags="-X 'main.version=$(VERSION)'" -o mini ./cmd/mini
else
	go build -o mini ./cmd/mini
endif

clean:
	rm mini
