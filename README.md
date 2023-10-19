# docker-image-lambda

## local environment setup
### build
```bash
# local環境用のイメージを作成する
$ docker build -t lambda-sample1-local -f './sample1/Dockerfile' . --build-arg GO_VERSION=$(GO_VERSION) --target local
```

### 起動
```bash
# lambdaのlocal環境用イメージをENTRYPOINTを指定して起動するとhttpリクエストを受けつける状態になる(8080)
$ docker run --name lambda-sample1-local -p 9000:8080 --rm lambda-sample1-local /functions/sample1
```

### 動作確認
```bash
# httpリクエストを受け付けているのでcurlでPOSTし起動する
$ curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations"
```

