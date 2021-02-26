GOCMD=go
GOBUILD=$(GOCMD) build

# Построить образ прежде чем собрать под windows
docker_prebuild_image:
	docker build -t exitstop/golang_bakend_msys2 -f docker/cross/Dockerfile .

windows:
	#GOOS=windows GOARCH=amd64 CGO_ENABLED=1 $(GOBUILD) -v -o build/speaker.exe cmd/voice/main.go
	./scripts/cross.sh
