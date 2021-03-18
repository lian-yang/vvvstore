FROM golang:1.15-alpine as builder

WORKDIR /go/src/github.com/lian-yang/vvvstore

ENV GOPROXY=https://goproxy.cn

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache upx ca-certificates tzdata

COPY . .

RUN go mod download

RUN make && \
    upx --best vvvstore -o upx_vvvstore && \
    mv -f upx_vvvstore vvvstore

FROM scratch as runner

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/lian-yang/vvvstore/vvvstore /app/

WORKDIR /app

RUN addgroup -S app && \
    adduser -S -g app app && \
    chown -R app:app ./

USER app

EXPOSE 9595

CMD ["/app/vvvstore"]

