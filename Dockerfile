FROM alpine:latest
RUN apk add --update ca-certificates
COPY bin/proxy /
RUN chmod +x proxy

ENTRYPOINT ["/proxy"]
EXPOSE 8080