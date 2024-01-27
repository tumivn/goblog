FROM golang:1.21.4-alpine3.18 as builder

COPY go.mod go.sum /go/src/app/
WORKDIR /go/src/app/
RUN go mod download
COPY . /go/src/app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-extldflags -static" -tags musl -o build/api github.com/legangs/cms/cmd/api


FROM alpine

WORKDIR /usr/bin/

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /go/src/app/.env /usr/bin/.env
COPY --from=builder /go/src/app/.env.docker /usr/bin/.env.docker
COPY --from=builder /go/src/app/build/app /usr/bin/app

EXPOSE 8080

ENTRYPOINT ["/usr/bin/app"]
