FROM golang:1.15-alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go-note
COPY . .
RUN go env && go mod tidy && go build -o server .

FROM alpine:latest

WORKDIR /go-note
COPY --from=builder /go-note/server ./
COPY --from=builder /go-note/conf ./conf
COPY --from=builder /go-note/static ./static

EXPOSE 8080

ENTRYPOINT ./server
