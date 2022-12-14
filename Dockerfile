FROM golang:1.18

ARG VERSION_HASH
ENV VERSION_HASH=$VERSION_HASH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /portfolio .

EXPOSE 3001

CMD ["/portfolio"]
