FROM golang:1.23.0-alpine3.20 as builder

COPY go.mod go.sum /go/src/app/
WORKDIR /go/src/app/
RUN go mod download
COPY . /go/src/app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-extldflags -static" -tags musl -o build/app github.com/tumivn/goblog/cmd/web


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /go/src/app/internal/server/static/ /usr/bin/build/internal/server/static/
COPY --from=builder /go/src/app/internal/server/views/ /usr/bin/build/internal/server/views/
COPY --from=builder /go/src/app/.env /usr/bin/build/.env
COPY --from=builder /go/src/app/build/app /usr/bin/build/app

WORKDIR /usr/bin/build/

EXPOSE 8080

ENTRYPOINT ["/usr/bin/build/app"]
