FROM golang:alpine as build
ARG VERSION=development

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=1 go build -a -tags 'netgo osusergo static_build' -ldflags="-X main.version=$VERSION" -v -o /usr/local/bin/pure-fb-om-exporter cmd/fb-om-exporter/main.go


# alpine is used here as it seems to be the minimal image that passes quay.io vulnerability scan
FROM alpine
COPY --from=build  /usr/local/bin/pure-fb-om-exporter /pure-fb-om-exporter
EXPOSE 9491
ENTRYPOINT ["/pure-fb-om-exporter"]
CMD ["--host", "0.0.0.0", "--port", "9491"]
