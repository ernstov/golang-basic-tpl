FROM golang:1.20-alpine

# Create new directory /app and make it our current directory
WORKDIR /app

# Copy go.mod files from source to /app. and download dependencies
COPY go.mod ./
RUN go mod download

# Copy source files to /app folder
COPY *.go ./

# Build our Go program to go-example binary
RUN go build -o /go-example

# Expose port 80
EXPOSE 80

# Start our program
CMD [ "/go-example" ]
