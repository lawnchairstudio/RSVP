# RSVP 
A gRPC service to help manage RSVPs to an event. This is a little side project to help further our knowledge of Go and get exposure to gRPC.

## Setup

1. Make sure that you have postgres installed
2. Run `createdb rsvp-dev` to get the db up.
3. Setup the needed go packages by running `go get google.golang.org/grpc` and `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`. These packages will be needed anytime you modify the /proto/RSVP.proto file. 
4. Also, run `go get` to get all the other packages needed.
5. Change directories to /proto using `cd proto` and then run `protoc --go_out=plugins=grpc:. RSVP.proto` to generate RSVP.pb.go (already included in the repo).
6. Change to the root directory of the repo `cd ..` and run `make db` to get the db all setup.
7. Run `go run main.go` to start the service. The service is now running on port 8000!

## Testing
Right now the app only has testing on the service layer and not around the actual gRPC part. Testing around the gRPC part at some point would be nice.

1. Change into the services directory `cd services`.
2. Run the tests using `go test`. 