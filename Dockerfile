# Stage 1 - Base APP
FROM golang:1.19.3 AS builder
LABEL maintainer="Hafidz98 <github.com/hafidz98>"
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a installsuffix cgo -o main .

# Stage 2 - Deploy
FROM gcr.io/distroless/static-debian11
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
# COPY --from=builder /app/.env .
CMD [ "." ]
EXPOSE 8991