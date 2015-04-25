export GOPATH=$(pwd)

go install driver
go install network
go install def
go install eventDetection
go install stateMachine
go install timer
go install encdec


go run main.go
