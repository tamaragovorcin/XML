FROM golang:1.16-alpine3.13
WORKDIR app
RUN apk update \
    && apk add sqlite \
    && apk add socat \
    && apk add build-base
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o campaign-service
EXPOSE 8080
CMD ["./campaign-service"]