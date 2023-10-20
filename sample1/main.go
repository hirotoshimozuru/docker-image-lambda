package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hirotoshimozuru/docker-image-lambda/internal/ldate"
	"golang.org/x/exp/slog"
)

func main() {
	lambda.Start(Sample1)
}

func Sample1(ctx context.Context) (string, error) {
	now := time.Now().In(ldate.JST).Format("2006-01-02 15:04:05")
	slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return now, nil
}
