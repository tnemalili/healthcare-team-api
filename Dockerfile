FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build

ENV CGO_ENABLED=0

RUN go build -o suppliers.api .

FROM alpine:latest
RUN addgroup -g 1000 noroot
RUN adduser -u 1000 -G noroot -h /home/noroot -D noroot
RUN mkdir /home/noroot/app
WORKDIR /home/noroot/app
EXPOSE API_PORT=3540
EXPOSE API_PORT=80
COPY --from=builder /build/suppliers.api /home/noroot/app/
CMD ["./suppliers.api"]