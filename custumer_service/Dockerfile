FROM golang:1.19-alpine
RUN mkdir custumer_service
COPY . /custumer_service
WORKDIR /custumer_service
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9088