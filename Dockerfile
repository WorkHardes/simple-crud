FROM golang:1.19-alpine

COPY . /simple-crud

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

WORKDIR /simple-crud

CMD [ "./app" ]
