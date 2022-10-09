FROM golang:1.19 AS builder
ENV CGO_ENABLED 0
ARG VERSION
WORKDIR /go/src/app
ADD . .
RUN go build -o /trunclog

FROM busybox
COPY --from=builder /trunclog /trunclog
ENTRYPOINT ["/trunclog"]
