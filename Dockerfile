#Stage 1: build
FROM golang:1.22.5-alpine AS build

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod download && go mod tidy

RUN apk add --no-cache make

RUN make build

RUN cat .env

RUN chmod +x /app/main


#Stage 2 : Run

FROM alpine:latest

RUN apk add --no-cache curl bash

COPY --from=build /app/main /main

COPY --from=build /app/.env /.env

EXPOSE 3000

CMD ["./main"]