### Build Stage
FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go build cmd/search/main.go

### Prod Stage
FROM golang:alpine AS prod

WORKDIR /opt/app

COPY --from=build /app/main ./
COPY .env ./

CMD ["./main"]
