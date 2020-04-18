#build stage
FROM golang:1.13 AS builder
WORKDIR $GOPATH/src/counter-queue
COPY . .
RUN go get -d -v .
RUN go install -v .

# Build the Go app
RUN CGO_ENABLED=0 go build -o /counter-queue

# final stage
FROM golang:1.13

### set root as working directory
WORKDIR /root/

####### COPYING to DOCKER CONTAINER #######
## COPY binary file
COPY --from=builder /counter-queue /root/

CMD ["/root/counter-queue"]