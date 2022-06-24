FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# install git 
RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /app

# Note here: To avoid downloading dependencies every time we
# build image. Here, we are caching all the dependencies by
# first copying go.mod and go.sum files and downloading them,
# to be used every time we build the image if the dependencies
# are not changed.

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies
RUN go mod download

# Now, copy the source code
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application
RUN CGO_enabled=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

# Finally ou multi-stage to build a small image
# Start a new stage from scratch
FROM scratch

# Copy pre-built binary file
COPY --from=builder /app/bin/main .

# Run executable
CMD ["./main"]