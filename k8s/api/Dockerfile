FROM golang:1.21-alpine3.18 as build
RUN apk update && apk add git gcc g++ && apk add --no-cache make
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . /app

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest
RUN apk update && apk add wget
#RUN apk --no-cache add ca-certificates

WORKDIR /app
# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main ./
COPY --from=build /app/.example.env ./.env

# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ./main serve

