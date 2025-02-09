FROM golang

WORKDIR /app

COPY ./app/go.mod .
COPY ./app/go.sum .

RUN go mod tidy

COPY ./app .

EXPOSE 8080