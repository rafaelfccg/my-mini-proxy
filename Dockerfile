FROM alpine:latest
RUN apk add --update ca-certificates
COPY bin/arbitrage /
RUN chmod +x arbitrage

ENTRYPOINT ["/arbitrage"]
EXPOSE 8080