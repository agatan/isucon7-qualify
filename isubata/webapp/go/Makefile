GOPATH := ${PWD}
export GOPATH

build:
	cd src/isubata && dep ensure -v -vendor-only
	go build -v isubata

vet:
	go vet ./src/isubata/...
