FROM haproxytech/haproxy-alpine:2.9.3

COPY haproxy/haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg
COPY haproxy/dataplaneapi.yaml /usr/local/etc/haproxy/dataplaneapi.yaml