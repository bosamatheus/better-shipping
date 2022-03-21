FROM golang:1.17-alpine

ARG CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /shipping-recommendations

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} \
    go build -o bin/shipping_recommendations cmd/shipping_recommendations/main.go

CMD ["./bin/shipping_recommendations"]
