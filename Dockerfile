FROM golang:1-alpine as builder
RUN apk update && apk add make
WORKDIR /build
ADD . .
RUN make build

FROM alpine
COPY --from=builder /build/wechat-channel /bin/wechat-channel
RUN chmod +x /bin/wechat-channel

ENTRYPOINT ["/bin/wechat-channel"]
