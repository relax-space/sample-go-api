FROM pangpanglabs/golang:builder AS builder

RUN go get github.com/fatih/structs \
    && go get github.com/relax-space/go-kit/...

ADD . /go/src/go-api
WORKDIR /go/src/go-api
ENV CGO_ENABLED=0
RUN go build -o go-api

FROM alpine
RUN apk --no-cache add ca-certificates
# FROM scratch
<<<<<<< HEAD
WORKDIR /go/src/sample-go-api
COPY --from=builder /go/src/sample-go-api/*.yml /go/src/sample-go-api/
COPY --from=builder /go/src/sample-go-api/sample-go-api /go/src/sample-go-api/
COPY --from=builder /go/src/sample-go-api/sample-go-api/sample_view.sql /go/src/sample-go-api/
COPY --from=builder /swagger-ui/ /go/src/sample-go-api/swagger-ui/
COPY --from=builder /go/src/sample-go-api/index.html /go/src/sample-go-api/swagger-ui/
=======
WORKDIR /go/src/go-api
COPY --from=builder /go/src/go-api/*.yml /go/src/go-api/
COPY --from=builder /go/src/go-api/go-api /go/src/go-api/
COPY --from=builder /swagger-ui/ /go/src/go-api/swagger-ui/
COPY --from=builder /go/src/go-api/index.html /go/src/go-api/swagger-ui/
>>>>>>> #2 add fruit sample


EXPOSE 8080

CMD ["./go-api"]
