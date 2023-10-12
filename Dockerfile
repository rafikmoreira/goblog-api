FROM golang:1.21-alpine

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application inside the container
RUN go get -d -v ./...
RUN go install -v ./...

# Set environment variables
ENV PORT=8080

# Expose the port the application will run on
EXPOSE $PORT

# Run the application
CMD ["app"]