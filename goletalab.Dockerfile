FROM golang:1.26
WORKDIR /workspace
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go build ./...
CMD ["go", "test", "./..."]
