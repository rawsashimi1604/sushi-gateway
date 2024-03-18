FROM golang:1.22 as build-stage

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /proxy /app/cmd

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:3.15 AS build-release-stage
WORKDIR /

COPY --from=build-stage /proxy /app/proxy

EXPOSE 8008 8443
CMD ["/app/proxy"]