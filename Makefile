Build:
	GO111MODULE=on go mod vendor

Build-10:
	go get -u golang.org/x/vgo
	${GOPATH}/bin/vgo mod vendor
