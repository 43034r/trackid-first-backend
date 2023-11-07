FROM golang:1.21.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY *.go ./
ADD controllers /app/controllers
ADD database /app/database
RUN ls -l
RUN ls -l /app/controllers
RUN ls -l /app/database
RUN go build -o /monitoriong.wiki/trackid-first-backend
EXPOSE 5000
CMD [ "monitoriong.wiki/trackid-first-backend" ]