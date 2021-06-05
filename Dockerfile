FROM golang:1.16 AS builder
RUN adduser -D dockuser
RUN dockuser
RUN chown dockuser:dockuser -R /app/
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD [ "./app" ]