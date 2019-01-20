FROM alpine:latest
RUN apk add --update ca-certificates
COPY bin/proxy /
RUN chmod +x proxy

ENTRYPOINT ["/arbitrage"]
EXPOSE 8080