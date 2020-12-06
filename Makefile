all:
	CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o pac-server.exe .