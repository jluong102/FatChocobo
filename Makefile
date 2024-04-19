VERSION="v0.0.0"
BUILD_DATE=$(shell date +%F)
BIN="fatchocobo"
LD_FLAGS="-s"

build:
	go build \
		-ldflags="${LD_FLAGS} -X 'main.VERSION=${VERSION}' -X 'main.BUILD_DATE=${BUILD_DATE}'" \
		-o ${BIN} ./cmd/${BIN}/*.go
debug:
	go build \
		-ldflags="-X 'main.VERSION=${VERSION}-debug' -X 'main.BUILD_DATE=${BUILD_DATE}'" \
		-o ${BIN}-debug ./cmd/${BIN}/*.go
docker:
	sudo docker build -t localhost/fatchocobo .
clean:
	sudo docker image rm localhost/fatchocobo
