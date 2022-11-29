FROM golang:1.19-alpine
RUN mkdir reating_service
COPY . /reating_service
WORKDIR /reating_service
RUN go mod tidy
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9084