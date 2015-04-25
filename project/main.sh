export GOPATH=$(pwd)

go install driver
go install network
go install def
go install eventDetection
go install stateMachine
go install timer


go run main.go
