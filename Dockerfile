FROM golang:1.21.1
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go mod tidy
COPY * ./
ADD controllers /app/controllers
ADD database /app/database
RUN ls
RUN go build -o /monitoriong.wiki/trackid-first-backend
EXPOSE 5000
ENTRYPOINT [ "monitoriong.wiki/trackid-first-backend" ]