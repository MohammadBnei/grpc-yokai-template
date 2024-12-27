## Multistage build
FROM golang:1.23 as build
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /src
COPY . .
RUN  go mod download
RUN go build -o /app

## Multistage deploy
FROM scratch

WORKDIR /
COPY --from=build /src/configs /configs
COPY --from=build /app /app

ENTRYPOINT ["/app"]
