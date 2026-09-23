FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /routing ./cmd/api

FROM alpine:3.21
RUN addgroup -S app && adduser -S app -G app
COPY --from=build /routing /routing
USER app
EXPOSE 8081
ENTRYPOINT ["/routing"]
