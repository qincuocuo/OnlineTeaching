FROM golang:1.17.2-alpine3.13 as builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache \
    make \
    git \
    gcc \
    libc-dev

WORKDIR /workspace

COPY . .

RUN cd build && make build

FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && mkdir -p /workspace/bin

WORKDIR /workspace/bin

COPY --from=builder /workspace/bin/learning /workspace/bin

RUN chmod +x /workspace/bin/learning

CMD ["/workspace/bin/learning"]
