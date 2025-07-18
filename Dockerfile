FROM debian:latest

WORKDIR /app

RUN apt update && apt-get install -y wget ca-certificates

RUN wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz && rm ./go1.24.0.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

COPY router /app/router
COPY public /app/public

RUN cd './router/' && go mod download && go build -o server .

EXPOSE 8080

CMD ["./router/server"]

