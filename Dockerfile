FROM golang:1.23 AS builder

RUN useradd -m -s /bin/bash htmx-user

USER htmx-user

RUN go install github.com/air-verse/air@latest

RUN export PATH=$PATH:~/go/bin/

WORKDIR /htmx

COPY . .

EXPOSE 8000/tcp

ENTRYPOINT [ "air" ]