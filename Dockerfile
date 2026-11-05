FROM alpine:3.22 AS certs
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY domain_exporter /domain_exporter
ENTRYPOINT ["/domain_exporter"]
