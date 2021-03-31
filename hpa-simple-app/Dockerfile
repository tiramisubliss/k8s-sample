FROM golang:1.8-alpine
ADD . /go/src/hpa-simple-app
RUN go install hpa-simple-app

FROM alpine:latest
COPY --from=0 /go/bin/hpa-simple-app .
ENV PORT 80
CMD ["./hpa-simple-app"]
