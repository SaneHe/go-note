FROM golang:1.15-alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /work-wechat
COPY . .
RUN go env && go build -o server .

FROM alpine:latest

WORKDIR /work-wechat
COPY --from=builder /work-wechat/server ./
COPY --from=builder /work-wechat/conf ./conf
COPY --from=builder /work-wechat/static ./static

EXPOSE 8080

ENTRYPOINT ./server
