FROM golang:alpine

WORKDIR /usr/app

COPY . .

RUN GOARCH=amd64 go build -o main .

EXPOSE 3000

CMD [ "./main" ]