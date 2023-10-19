package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/exp/slog"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

func main() {
	lambda.Start(Sample1)
}

func Sample1(ctx context.Context) (string, error) {
	now := time.Now().In(jst).Format("2006-01-02 15:04:05")
	slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return now, nil
}
