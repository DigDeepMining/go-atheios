FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.5 https://github.com/atheioschain/go-atheios && \
  (cd go-atheios && make gath) && \
  cp go-atheios/build/bin/gath /gath && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-atheios && rm -rf /var/cache/apk/*

EXPOSE 8696
EXPOSE 30696

ENTRYPOINT ["/gath"]
