FROM ubuntu:latest

WORKDIR /app

ADD . /app

EXPOSE 8080

CMD ["./lookup","-s"]
