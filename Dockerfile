FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY internal ./internal
COPY docs ./docs

RUN go mod download
RUN go build -o /event-store-service

FROM alpine
WORKDIR /app
COPY --from=build /event-store-service /

ENV APP_ENV=integration
EXPOSE 8080

ENTRYPOINT [ "/event-store-service" ]