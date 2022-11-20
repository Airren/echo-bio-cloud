FROM golang:alpine3.16 as builder
COPY ./ /echo-bio-cloud
WORKDIR /echo-bio-cloud
RUN go build -o ./bin/echo-bio-cloud main.go



FROM alpine
COPY --from=builder /echo-bio-cloud/bin/echo-bio-cloud /echo-bio-cloud
ENTRYPOINT "/metric-exporter"
