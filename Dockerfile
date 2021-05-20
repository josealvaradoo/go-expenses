# Image to build the Go binary
FROM golang:1.16-alpine as builder

# Set Env Vars
ARG APP_ENV

# Build application
RUN mkdir -p /home/app
ADD . /home/app
WORKDIR /home/app
RUN go clean --modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Image to run our app
FROM alpine
COPY --from=builder /home/app/main /home

# Set Env Vars
ENV APP_ENV $APP_ENV

WORKDIR /home

# Expose port
EXPOSE 8000

# Run command
CMD if [ "$APP_ENV" = "production" ]; then ./main; fi