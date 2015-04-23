export GOPATH=$(pwd)

go install driver
go install network

go run main.go
