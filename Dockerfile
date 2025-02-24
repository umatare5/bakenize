FROM golang:1.24 as builder

WORKDIR /app
COPY . .
RUN make build

FROM busybox:glibc AS app

WORKDIR /app
COPY --from=builder /app/dist/bakenize /bin/

ENTRYPOINT ["/bin/bakenize"]
