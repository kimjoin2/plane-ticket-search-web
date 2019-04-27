FROM golang:1.11-alpine AS build
COPY ./ /go/src/plane-ticket-search-web/
WORKDIR /go/src/plane-ticket-search-web/
RUN go build -o plane-ticket-search-web

FROM alpine:3.7
RUN apk --no-cache --update upgrade \
  && apk add --update --no-cache tzdata ca-certificates \
  && update-ca-certificates --fresh
RUN mkdir -p /app
COPY --from=build /go/src/plane-ticket-search-web/plane-ticket-search-web /app/plane-ticket-search-web

EXPOSE 80
EXPOSE 443

ENTRYPOINT ["/app/plane-ticket-search-web"]