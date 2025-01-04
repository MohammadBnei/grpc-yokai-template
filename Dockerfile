## Multistage build
FROM golang:1.23 AS build
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /src

COPY ["go.mod", "go.sum", "./"]
RUN  go mod download

COPY . .
RUN go build -o /app

## Multistage deploy
FROM scratch

WORKDIR /
COPY --from=build /src/configs /configs
COPY --from=build /app /app

ENTRYPOINT ["/app"]
