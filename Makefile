build:
	@echo "building app"
	go build -o app cmd/server/main.go

run:
	@echo "Running app"
	./app