    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=/usr/local/go
    export GOPATH=$GOPATH:$PWD
    export PATH=$PATH:$PWD
    GO111MODULE=on  go run src/main/main.go