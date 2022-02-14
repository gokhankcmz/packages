FROM golang:1.17-alpine

WORKDIR /Packages
COPY . .
RUN go mod download

WORKDIR /Packages/src/api
RUN go build -o /api
EXPOSE 80
CMD ["/api"]
