FROM golang:1.21.1-alpine
WORKDIR /app
COPY * ./
COPY go.mod ./
RUN go mod download
RUN go mod tidy
RUN go build -o /monitoriong.wiki/trackid-first-backend
EXPOSE 5000
CMD [ "monitoriong.wiki/trackid-first-backend" ]