FROM golang:1.19-alpine
RUN mkdir exam_api
COPY . /exam_api
WORKDIR /exam_api
RUN go mod tidy
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9079