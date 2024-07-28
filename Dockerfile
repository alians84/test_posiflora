FROM golang:1.20 as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN apt install -y tzdata

RUN CGO_ENABLED=0 go build -o main .

FROM alpine:3.9 as app

ARG APP_PORT=:80

COPY ./docs /app/docs
COPY --from=builder /app/main /


EXPOSE 80

CMD ["/main"]