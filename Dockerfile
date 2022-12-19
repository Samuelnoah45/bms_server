FROM golang:alpine

COPY go.mod ./
COPY go.sum ./

# COPY ./firebase-auth.json /server/firebase-auth.json
COPY *.go ./
WORKDIR /bms_server

RUN go mod download
RUN go build -o /bms_server
EXPOSE 8181
CMD ["/bms_server"]