FROM golang:1.19-alpine
RUN mkdir post_service
COPY . /post_service
WORKDIR /post_service
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9093