FROM golang:1.18.3-alpine3.16 as build

COPY . /app
WORKDIR /app
RUN go build  main.go

FROM alpine 
WORKDIR /app
RUN apk update && apk add wget

COPY --from=build /app/main ./
COPY --from=build /app/.example.env ./.env

CMD ./main serve

