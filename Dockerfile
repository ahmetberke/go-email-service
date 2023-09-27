FROM golang:1.19

ENV PROFILE=docker

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /go-email-service

EXPOSE 8080

CMD [ "/go-email-service" ]