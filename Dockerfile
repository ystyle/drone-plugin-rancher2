FROM golang:alpine AS build-env
ADD . /go/src/app
WORKDIR /go/src/app
RUN go build -v -o /go/src/app/drone-plugin main.go

FROM alpine
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=build-env /go/src/app/drone-plugin /bin/drone-plugin
ENTRYPOINT /bin/drone-plugin
