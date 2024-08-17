FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build
COPY data.json ./
CMD ["./tt-ranking-calculator"]
EXPOSE 8080