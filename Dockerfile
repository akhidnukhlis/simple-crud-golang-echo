# Load golang image to build
FROM golang:1.20 as builder

# make directory and copy go project
RUN mkdir -p /app
WORKDIR /app
COPY . .

# application builder step
RUN go clean && go mod tidy && go mod download && go mod vendor
RUN go build -o main

# Resevre port
EXPOSE 8081

# run project if container start
CMD ["./main"]