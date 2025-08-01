FROM golang:alpine as build
ARG VERSION=1.1.4

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=1 go build -mod=readonly -a -tags 'netgo osusergo static_build' -ldflags="-X 'main.version=v$VERSION' -X 'purestorage/fb-openmetrics-exporter/internal/rest-client.UserAgentVersion=$VERSION'" -v -o /usr/local/bin/pure-fb-om-exporter cmd/fb-om-exporter/main.go


# alpine is used here as it seems to be the minimal image that passes quay.io vulnerability scan
FROM alpine
# update ssl packages for CVEs
RUN apk update && apk add --upgrade libcrypto3 libssl3 && rm -rf /var/cache/apk/*
COPY --from=build  /usr/local/bin/pure-fb-om-exporter /pure-fb-om-exporter

# create an empty tokens file for use with volumes if required. You can use a mounted volume to /etc/pure-fb-om-exporter/ to pass the `tokens.yaml` file. File must be named `tokens.yaml`.
RUN mkdir /etc/pure-fb-om-exporter && touch /etc/pure-fb-om-exporter/tokens.yaml

EXPOSE 9491
ENTRYPOINT ["/pure-fb-om-exporter"]
CMD ["--address", "0.0.0.0", "--port", "9491", "--tokens", "/etc/pure-fb-om-exporter/tokens.yaml"]
