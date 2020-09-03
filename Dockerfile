FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD app -port ${PORT:-:443} -public_cert ${PUBLIC_CERT} -private_cert ${PRIVATE_CERT} 