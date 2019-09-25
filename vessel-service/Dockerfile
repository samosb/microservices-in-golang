# vessel-service/Dockerfile
FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/micro-service-vessel

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o micro-service-vessel


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/micro-service-vessel .

CMD "./micro-service-vessel"
