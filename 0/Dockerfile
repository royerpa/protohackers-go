FROM golang:latest as build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
# COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -o /main .

FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /main /main

WORKDIR /app
EXPOSE 7
CMD ["/main"]