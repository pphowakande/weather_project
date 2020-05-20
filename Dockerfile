# We specify the base image we need for our
# go application
FROM golang:1.12.0-alpine3.9
# We copy everything in the root directory
# into our /weather directory
ADD . /go/src/weather
# We specify that we now wish to execute 
# any further commands inside our /weather
# directory
WORKDIR /go/src/weather
# we run go build to compile the binary
# executable of our Go program
RUN go get weather
# install required packages
RUN go install
# Our start command which kicks off
# our newly created binary executable
ENTRYPOINT ["/go/bin/weather"]