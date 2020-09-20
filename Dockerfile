FROM golang:1.15.1-buster AS builder

COPY ./cache /cache

WORKDIR /build
ENV GOPROXY https://goproxy.cn
ENV GOMODCACHE=/cache
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o sendmail_app .

FROM alpine:3.10 AS final
ENV TZ=Asia/Shanghai
EXPOSE 8080

WORKDIR /app
COPY --from=builder /build/sendmail_app /app/
#COPY --from=builder /build/config/app.toml /app/config/app.toml
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app/sendmail_app"]
#
##docker build -t echinacities/sendmail:v1 .
##docker run -ti -v /Users/Weile/go/src/echina/sendmail/config:/app/config echinajobs/sendmail:v1