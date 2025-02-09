FROM golang

WORKDIR /app

COPY ./app/go.mod .
COPY ./app/go.sum .

RUN go mod download
RUN go mod tidy

EXPOSE 8080