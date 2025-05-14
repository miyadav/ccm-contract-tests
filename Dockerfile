FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest

ENTRYPOINT ["ginkgo", "-v", "./..."]

