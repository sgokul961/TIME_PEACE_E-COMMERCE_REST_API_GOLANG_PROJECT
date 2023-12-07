FROM golang:1.21.0 as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /api -v ./cmd/api



EXPOSE 3000


CMD ["/api"]