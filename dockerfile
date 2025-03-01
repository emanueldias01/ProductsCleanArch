FROM golang:1.23.2-alpine AS build

WORKDIR /app

COPY ./cmd /app/cmd
COPY ./controller /app/controller
COPY ./db /app/db
COPY ./model /app/model
COPY ./repository /app/repository
COPY ./router /app/router
COPY ./usecase /app/usecase
COPY ./migrations /app/migrations
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum


RUN go build ./cmd/main.go

FROM alpine:latest AS prod

EXPOSE 8080
WORKDIR /app

COPY --from=build /app/main /app/

CMD [ "./main" ]