FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
# use the correct file name for building
RUN go build -o main /app/cmd/main.go
# make the main file executable
RUN chmod +x main
EXPOSE 8080
CMD [ "./main" ]