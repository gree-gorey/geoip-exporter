FROM golang:1.10.1

RUN go get -d -v github.com/gree-gorey/geoip-exporter/cmd/geoip-exporter
WORKDIR /go/src/github.com/gree-gorey/geoip-exporter/cmd/geoip-exporter
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o geoip-exporter .

FROM alpine:3.7
WORKDIR /root/
RUN apk install -y ca-certificates
COPY --from=0 /go/src/github.com/gree-gorey/geoip-exporter/cmd/geoip-exporter/geoip-exporter .
CMD ["./geoip-exporter"]
