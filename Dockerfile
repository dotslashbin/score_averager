FROM golang

WORKDIR /go/src/

RUN go get github.com/githubnemo/CompileDaemon

COPY . .

ENTRYPOINT go get -d github.com/gorilla/mux && CompileDaemon -build="go build -o api" -command="./api"
