FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /app

# Install git to clone the repository
RUN apt-get update && apt-get install -y --no-install-recommends git

# Clone the repository
RUN git clone https://github.com/M1chlCZ/go-webapi-skeleton.git .

# Install the necessary packages
RUN go get -d -v ./...
RUN go install -v ./...

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Final stage
FROM alpine:3.14

# Copy the binary from the build stage to the final stage
COPY --from=build /app/app .

# Copy the .env file to the final stage
COPY .env .

# Expose the API port
EXPOSE 8080

# Run the API
CMD ["./app"]