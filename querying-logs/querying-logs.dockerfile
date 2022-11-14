FROM alpine:latest

RUN mkdir /app

COPY queryinglogs /app

CMD [ "/app/queryinglogs"]