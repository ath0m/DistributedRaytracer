# Use the official Go image as the base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /agent

# Copy the Go module files
COPY go.mod ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go program
RUN go build -o agent

# Expose the port that the application runs on
EXPOSE 8090

# Set the entry point for the container
ENTRYPOINT ["./agent"]