ARG BUILD_IMAGE=grafana/loki-build-image:0.2.1
# Directories in this file are referenced from the root of the project not this folder
# This file is intented to be called from the root like so:
# docker build -t grafana/promtail -f cmd/promtail/Dockerfile .
FROM golang:1.11.4-alpine as goenv
RUN go env GOARCH > /goarch && \
    go env GOARM > /goarm

FROM --platform=linux/amd64 $BUILD_IMAGE as build
COPY --from=goenv /goarch /goarm /
COPY . /go/src/github.com/grafana/loki
WORKDIR /go/src/github.com/grafana/loki
RUN GOARCH=$(cat /goarch) GOARM=$(cat /goarm) make clean && make promtail

FROM alpine:3.9
# tzdata required for the timestamp stage to work
RUN apk add --update --no-cache ca-certificates tzdata
COPY --from=build /go/src/github.com/grafana/loki/cmd/promtail/promtail /usr/bin/promtail
COPY cmd/promtail/promtail-local-config.yaml /etc/promtail/local-config.yaml
COPY cmd/promtail/promtail-docker-config.yaml /etc/promtail/docker-config.yaml
ENTRYPOINT ["/usr/bin/promtail"]
