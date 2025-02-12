FROM golang

WORKDIR /app

COPY ./app/go.mod .
COPY ./app/go.sum .

RUN go install github.com/air-verse/air@latest

RUN go mod download
RUN go mod tidy

#CMD ["air", "-c", ".air.toml"]