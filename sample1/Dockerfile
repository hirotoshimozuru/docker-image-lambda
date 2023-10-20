FROM golang:1.21.0 AS builder

ENV ARCH="arm64"
ENV APPDIR $GOPATH/src/docker-image-lambda

COPY go.mod go.sum $APPDIR/
COPY vendor/ $APPDIR/vendor/
COPY sample1/main.go $APPDIR/
COPY internal/ $APPDIR/internal/
WORKDIR $APPDIR

RUN --mount=type=cache,target=/root/.cache/go-build GOOS=linux GOARCH=${ARCH} CGO_ENABLED=0 go build -mod=vendor -ldflags "-s -w" -o /functions/main main.go

# aws環境実行用
FROM gcr.io/distroless/static as aws
COPY --from=builder /functions /functions
ENTRYPOINT ["/functions/sample1"]

# local環境実行用
# lambdaのlocal環境用イメージをENTRYPOINTを指定して起動すると、8080でhttpリクエストを受けつける
FROM public.ecr.aws/lambda/provided:al2 as local
COPY --from=builder /functions /functions
ENTRYPOINT ["/usr/local/bin/aws-lambda-rie"]
