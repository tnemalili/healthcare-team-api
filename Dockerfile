FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build

ENV CGO_ENABLED=0

RUN go build -o auth.api .

FROM alpine:latest
RUN addgroup -g 1000 noroot
RUN adduser -u 1000 -G noroot -h /home/noroot -D noroot
RUN mkdir /home/noroot/app
WORKDIR /home/noroot/app
COPY --from=builder /build/auth.api /home/noroot/app/
CMD ["./auth.api"]