FROM golang:1.15.10-alpine3.13 as builder

ENV GOPROXY https://goproxy.io,direct
ENV GOPRIVATE *.git.moresec.cn
ENV GO111MODULE on


RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache \
    make

WORKDIR /workspace

COPY . .

RUN cd build && make webapi-srv

FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache bash \
    mkdir -p /workspace/data

WORKDIR /workspace/bin

COPY --from=builder /workspace/bin/webapi_srv /workspace/bin
COPY config/config.yaml /workspace/config/config.yaml
# COPY --from=builder /home/moresec/ds_server/deployments/health.sh /run

RUN chmod +x /workspace/bin/webapi_srv

CMD ["/workspace/bin/webapi_srv", "--conf_path", "/workspace/config/config.yaml"]
