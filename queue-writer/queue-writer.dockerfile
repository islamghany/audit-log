FROM alpine:latest

RUN mkdir /app

COPY queueWriter /app

CMD [ "/app/queueWriter"]