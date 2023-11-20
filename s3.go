package stat

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Stat struct {
	httpClient *http.Client
}

func headObject(client *s3.Client, ctx context.Context, bucket, key string) (*s3.HeadObjectOutput, error) {
	input := &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}
	return client.HeadObject(ctx, input)
}

func NewS3Stat(client *http.Client) *S3Stat {
	return &S3Stat{httpClient: client}
}

func (ss *S3Stat) Stat(name string) (*StatInfo, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	uri, err := url.Parse(name)
	if err != nil {
		return nil, err
	}

	var bucket, key string
	if subs := strings.Split(uri.Path, "/"); len(subs) > 1 {
		bucket = subs[1]
		key = strings.TrimPrefix(strings.TrimPrefix(uri.Path, "/"+bucket), "/")
	}

	if ss.httpClient != nil {
		cfg.HTTPClient = ss.httpClient
	}

	var (
		headOutput *s3.HeadObjectOutput
		statInfo   *StatInfo
	)
	if headOutput, err = headObject(s3.NewFromConfig(cfg), context.TODO(), bucket, key); err == nil {
		var partsCount int64
		if headOutput.PartsCount != nil {
			partsCount = int64(*headOutput.PartsCount)
		}

		statInfo = &StatInfo{
			Name:     name,
			Size:     *headOutput.ContentLength,
			Mtime:    *headOutput.LastModified,
			Blocks:   partsCount,
			Metadata: headOutput.Metadata,
		}
	}
	return statInfo, err
}
