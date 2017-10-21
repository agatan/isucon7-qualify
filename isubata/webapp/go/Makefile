GOPATH := ${PWD}
export GOPATH

build:
	go build -v isubata

dep:
	cd src/isubata && dep ensure -v -vendor-only

vet:
	go vet ./src/isubata/...
