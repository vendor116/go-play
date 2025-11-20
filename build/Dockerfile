ARG ALPINE_VERSION=3.22

FROM alpine:${ALPINE_VERSION}

RUN apk add --no-cache ca-certificates

COPY ./bin/playgo /playgo

CMD ["/playgo"]