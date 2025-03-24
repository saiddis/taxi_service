#syntax=docker/dockerfile:1

FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server_session

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /server_session /server_session

EXPOSE 8080

USER nonroot:nonroot

CMD ["/server_session"]
