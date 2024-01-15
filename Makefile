build:
	@go build -o build/main main.go

build-release:
	@go build -ldflags "-s -w" -o build/main main.go

run: clean build
	./build/main

clean:
	@rm -rf build || true

watch:
	@CompileDaemon -polling -color="true" -directory="." -graceful-kill="true" -build="make build" -command="make run"|| echo "CompileDaemon not found"
