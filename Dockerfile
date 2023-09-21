# Use the official Go image to create a build environment
FROM golang:1.17 AS build

# Set the current working directory
WORKDIR /src

# Copy the local package files to the container's workspace
COPY . .

# Build the binary
RUN go build -o /out/ops-tech-challenge

# Use a small base image to create a final image
FROM alpine:3.14 AS runtime

# Copy the binary from the build stage
COPY --from=build /out/ops-tech-challenge /app/ops-tech-challenge

# Set the environment variable for the shared secret
ENV SECRET=yoursharedsecret

# Set the working directory
WORKDIR /app

# Run the service on container startup
CMD ["./ops-tech-challenge"]
