package reader

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

//go:generate go run ../cmd/ -output reader.go

// New returns an object which also contains the accountID and the region to use.
//
// The accountID is helpful to return only the AMI or snapshots that belong to the account.
//
// # While the region has to be a valid AWS region
//
// An error is returned if any of the needed AWS request for creating the reader returns an AWS error, in such case it
// will have any of the common error codes (see below) or EmptyStaticCreds code or a go standard error in case that no
// regions are matched with the ones available, at the time, in AWS.
// See:
//   - https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html#CommonErrors
//   - https://docs.aws.amazon.com/STS/latest/APIReference/CommonErrors.html
func New(ctx context.Context, accessKey, secretKey, region, sessionToken string, config *aws.Config) (Reader, error) {
	var c = connector{}

	creds, ec2s, sts, err := configureAWS(accessKey, secretKey, region, sessionToken)
	if err != nil {
		return nil, err
	}
	c.creds = creds
	if err := c.setAccountID(ctx, sts); err != nil {
		return nil, err
	}

	if err = c.setRegion(ctx, ec2s, region); err != nil {
		return nil, err
	}

	c.setService(config)

	return &c, nil
}

// The connector provides easy access to AWS SDK calls.
//
// By using it, calls can be made directly through multiple regions, and will filter only data that belongs to you.
// For example, when fetching the list of AMI, or snapshots.
//
// In order to start making calls, only calling New is required.
type connector struct {
	region    string
	svc       *serviceConnector
	creds     aws.CredentialsProvider
	accountID string
}

func (c *connector) GetAccountID() string {
	return c.accountID
}

func (c *connector) GetRegion() string {
	return c.region
}

type serviceConnector struct {
	config                   aws.Config
	apigateway               *apigateway.Client
	athena                   *athena.Client
	autoscaling              *autoscaling.Client
	batch                    *batch.Client
	cloudfront               *cloudfront.Client
	cloudwatch               *cloudwatch.Client
	configservice            *configservice.Client
	databasemigrationservice *databasemigrationservice.Client
	dax                      *dax.Client
	directconnect            *directconnect.Client
	directoryservice         *directoryservice.Client
	dynamodb                 *dynamodb.Client
	ec2                      *ec2.Client
	ecs                      *ecs.Client
	efs                      *efs.Client
	eks                      *eks.Client
	elasticache              *elasticache.Client
	elasticbeanstalk         *elasticbeanstalk.Client
	elasticsearchservice     *elasticsearchservice.Client
	elb                      *elasticloadbalancing.Client
	elbv2                    *elasticloadbalancingv2.Client
	emr                      *emr.Client
	fsx                      *fsx.Client
	glue                     *glue.Client
	iam                      *iam.Client
	kinesis                  *kinesis.Client
	lambda                   *lambda.Client
	lightsail                *lightsail.Client
	mediastore               *mediastore.Client
	mq                       *mq.Client
	neptune                  *neptune.Client
	rds                      *rds.Client
	redshift                 *redshift.Client
	region                   string
	route53resolver          *route53resolver.Client
	route53                  *route53.Client
	s3                       *s3.Client
	ses                      *ses.Client
	sqs                      *sqs.Client
	storagegateway           *storagegateway.Client
}

/* The default region is only used to (1) get the list of region and
 * (2) get the account ID associated with the credentials.
 *
 * It is not used as a default region for services, therefore if no
 * region is specified when instantiating the connector, then it will
 * not try to establish any connections with AWS services.
 */
const defaultRegion string = "eu-west-1"

// configureAWS creates a new static credential with the passed accessKey and
// secretKey and with it, a sessions which is used to create a EC2 client and
// a Security Token Service client.
// The only AWS error code that this function return is
// * EmptyStaticCreds
func configureAWS(accessKey, secretKey, region, token string) (aws.CredentialsProvider, *ec2.Client, *sts.Client, error) {
	if region == "" {
		region = defaultRegion
	}

	creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, token))
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithCredentialsProvider(creds))
	if err != nil {
		return nil, nil, nil, err
	}

	return creds, ec2.NewFromConfig(cfg), sts.NewFromConfig(cfg), nil
}

// setAccountID retrieves the caller ID from the Security Token Service and set
// it in the connector.
// An AWS error can be returned with one of the common error codes.
// See https://docs.aws.amazon.com/STS/latest/APIReference/CommonErrors.html
func (c *connector) setAccountID(ctx context.Context, sts *sts.Client) error {
	resp, err := sts.GetCallerIdentity(ctx, nil)
	if err != nil {
		return err
	}
	c.accountID = *resp.Account
	return nil
}

// setRegion retrieves the AWS available regions and matches with the passed
// region.
// A AWS error can be returned with one of the common error codes or a standard
// go error if enabledRegions is empty or if 0 AWS regions has been matched.
// See https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html#CommonErrors
func (c *connector) setRegion(ctx context.Context, ec2 *ec2.Client, region string) error {
	if region == "" {
		return errors.New("at least one region name is required")
	}

	regions, err := ec2.DescribeRegions(ctx, nil)
	if err != nil {
		return err
	}

	for _, r := range regions.Regions {
		if region == *r.RegionName {
			c.region = region
			return nil
		}
	}

	if c.region == "" {
		return fmt.Errorf("found 0 regions matching: %v", region)
	}

	return nil
}

func (c *connector) setService(config *aws.Config) {
	if config != nil {
		config.Credentials = c.creds
	} else {
		config = &aws.Config{
			Region:      c.region,
			Credentials: c.creds,
		}
	}

	cfg := *config

	svc := &serviceConnector{
		config:                   cfg,
		region:                   c.region,
		apigateway:               apigateway.NewFromConfig(cfg),
		athena:                   athena.NewFromConfig(cfg),
		autoscaling:              autoscaling.NewFromConfig(cfg),
		batch:                    batch.NewFromConfig(cfg),
		cloudfront:               cloudfront.NewFromConfig(cfg),
		cloudwatch:               cloudwatch.NewFromConfig(cfg),
		configservice:            configservice.NewFromConfig(cfg),
		databasemigrationservice: databasemigrationservice.NewFromConfig(cfg),
		dax:                      dax.NewFromConfig(cfg),
		directconnect:            directconnect.NewFromConfig(cfg),
		directoryservice:         directoryservice.NewFromConfig(cfg),
		dynamodb:                 dynamodb.NewFromConfig(cfg),
		ec2:                      ec2.NewFromConfig(cfg),
		ecs:                      ecs.NewFromConfig(cfg),
		efs:                      efs.NewFromConfig(cfg),
		eks:                      eks.NewFromConfig(cfg),
		elasticache:              elasticache.NewFromConfig(cfg),
		elasticbeanstalk:         elasticbeanstalk.NewFromConfig(cfg),
		elasticsearchservice:     elasticsearchservice.NewFromConfig(cfg),
		elb:                      elasticloadbalancing.NewFromConfig(cfg),
		elbv2:                    elasticloadbalancingv2.NewFromConfig(cfg),
		emr:                      emr.NewFromConfig(cfg),
		fsx:                      fsx.NewFromConfig(cfg),
		glue:                     glue.NewFromConfig(cfg),
		iam:                      iam.NewFromConfig(cfg),
		kinesis:                  kinesis.NewFromConfig(cfg),
		lambda:                   lambda.NewFromConfig(cfg),
		lightsail:                lightsail.NewFromConfig(cfg),
		mediastore:               mediastore.NewFromConfig(cfg),
		mq:                       mq.NewFromConfig(cfg),
		neptune:                  neptune.NewFromConfig(cfg),
		rds:                      rds.NewFromConfig(cfg),
		redshift:                 redshift.NewFromConfig(cfg),
		route53resolver:          route53resolver.NewFromConfig(cfg),
		route53:                  route53.NewFromConfig(cfg),
		s3:                       s3.NewFromConfig(cfg),
		ses:                      ses.NewFromConfig(cfg),
		sqs:                      sqs.NewFromConfig(cfg),
		storagegateway:           storagegateway.NewFromConfig(cfg),
	}
	c.svc = svc
}
