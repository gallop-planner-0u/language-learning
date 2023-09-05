FROM golang:1.21 as build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go
RUN chmod +x main
EXPOSE 8080
CMD [ "./main" ]