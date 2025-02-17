package reader

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (c *connector) ListBuckets(ctx context.Context, input *s3.ListBucketsInput) ([]s3types.Bucket, error) {
	if c.svc.s3 == nil {
		c.svc.s3 = s3.NewFromConfig(c.svc.config)
	}

	opt := s3.ListBucketsOutput{}

	if input == nil {
		input = &s3.ListBucketsInput{}
	}

	if c.svc.region != "" {
		input.BucketRegion = &c.svc.region
	}

	hasNextToken := true
	for hasNextToken {
		o, err := c.svc.s3.ListBuckets(ctx, input)
		if err != nil {
			return nil, err
		}
		if o.Buckets == nil {
			hasNextToken = false
			continue
		}

		if input == nil {
			input = &s3.ListBucketsInput{}
		}
		input.ContinuationToken = o.ContinuationToken
		hasNextToken = o.ContinuationToken != nil

		opt.Buckets = append(opt.Buckets, o.Buckets...)
	}

	return opt.Buckets, nil
}
