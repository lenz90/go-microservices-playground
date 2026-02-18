FROM golang:1.22-alpine AS build
WORKDIR /src
COPY . .
RUN go mod tidy && go build -o /app/server ./cmd/server
FROM alpine:3.20
COPY --from=build /app/server /server
ENTRYPOINT ["/server"]
