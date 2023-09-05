FROM golang:1.21 as build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go
FROM debian:12.1
RUN apt-get update
RUN apt-get -y install postgresql
COPY --from=build /app/main /app/main
COPY --from=build /app/config /app/config
COPY --from=build /app/run.sh /app/run.sh
RUN chmod +x /app/run.sh
EXPOSE 8080
USER postgres
CMD [ "/app/run.sh" ]