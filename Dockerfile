FROM golang:1.17.8

WORKDIR /go/MyMiniProject
COPY go.mod go.sum ./
RUN go mod download
ADD main.go /go/MyMiniProject/
COPY . .
EXPOSE 1234
CMD ["go", "run", "main.go"]