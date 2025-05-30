module github.com/cycloidio/terracognita

go 1.24.0

require (
	github.com/Azure/azure-sdk-for-go v65.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.27
	github.com/adrg/xdg v0.2.3
	github.com/aws/aws-sdk-go-v2 v1.36.2
	github.com/aws/aws-sdk-go-v2/config v1.29.7
	github.com/aws/aws-sdk-go-v2/credentials v1.17.60
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.28.12
	github.com/aws/aws-sdk-go-v2/service/athena v1.49.11
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.51.13
	github.com/aws/aws-sdk-go-v2/service/batch v1.49.13
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.44.11
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.43.15
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.45.14
	github.com/aws/aws-sdk-go-v2/service/configservice v1.51.14
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.49.1
	github.com/aws/aws-sdk-go-v2/service/dax v1.23.16
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.30.13
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.30.18
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.40.2
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.203.1
	github.com/aws/aws-sdk-go-v2/service/ecs v1.53.16
	github.com/aws/aws-sdk-go-v2/service/efs v1.34.12
	github.com/aws/aws-sdk-go-v2/service/eks v1.58.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.44.14
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.28.17
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.28.18
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.43.13
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.32.19
	github.com/aws/aws-sdk-go-v2/service/emr v1.47.13
	github.com/aws/aws-sdk-go-v2/service/fsx v1.52.1
	github.com/aws/aws-sdk-go-v2/service/glue v1.105.10
	github.com/aws/aws-sdk-go-v2/service/iam v1.39.2
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.32.20
	github.com/aws/aws-sdk-go-v2/service/lambda v1.69.14
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.42.17
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.24.16
	github.com/aws/aws-sdk-go-v2/service/mq v1.27.17
	github.com/aws/aws-sdk-go-v2/service/neptune v1.35.18
	github.com/aws/aws-sdk-go-v2/service/rds v1.93.14
	github.com/aws/aws-sdk-go-v2/service/redshift v1.53.13
	github.com/aws/aws-sdk-go-v2/service/route53 v1.48.8
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.34.14
	github.com/aws/aws-sdk-go-v2/service/s3 v1.77.1
	github.com/aws/aws-sdk-go-v2/service/ses v1.29.11
	github.com/aws/aws-sdk-go-v2/service/sqs v1.37.15
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.35.2
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.15
	github.com/aws/smithy-go v1.22.3
	github.com/chr4/pwgen v1.1.0
	github.com/cycloidio/mxwriter v1.0.4
	github.com/cycloidio/tfdocs v0.0.0-20230516095646-1dc8f8412d50
	github.com/gertd/go-pluralize v0.2.1
	github.com/golang/mock v1.6.0
	github.com/hashicorp/go-azure-helpers v0.40.0
	github.com/hashicorp/go-cty v1.5.0
	github.com/hashicorp/hcl/v2 v2.23.0
	github.com/hashicorp/terraform v0.13.0
	github.com/hashicorp/terraform-plugin-go v0.27.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.37.0
	github.com/hashicorp/terraform-provider-aws v1.60.1-0.20210513231836-489654890359
	github.com/hashicorp/terraform-provider-azurerm v1.44.1-0.20201029183808-d721bcc1bb55
	github.com/hashicorp/terraform-provider-google v1.20.1-0.20210510171431-a764cf3da527
	github.com/hashicorp/terraform-provider-vsphere v1.26.1-0.20220510172607-30f37d268d79
	github.com/pascaldekloe/name v1.0.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.10.0
	github.com/vmware/govmomi v0.28.0
	github.com/zclconf/go-cty v1.16.2
	golang.org/x/sync v0.14.0
	golang.org/x/text v0.25.0
	google.golang.org/api v0.235.0
	google.golang.org/grpc v1.72.2
	gopkg.in/yaml.v2 v2.4.0
)

require (
	bitbucket.org/creachadair/stringset v0.0.8 // indirect
	cel.dev/expr v0.24.0 // indirect
	cloud.google.com/go v0.121.2 // indirect
	cloud.google.com/go/auth v0.16.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.8 // indirect
	cloud.google.com/go/bigtable v1.37.0 // indirect
	cloud.google.com/go/compute/metadata v0.7.0 // indirect
	cloud.google.com/go/iam v1.5.2 // indirect
	cloud.google.com/go/longrunning v0.6.7 // indirect
	cloud.google.com/go/monitoring v1.24.2 // indirect
	cloud.google.com/go/storage v1.54.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.5 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/GoogleCloudPlatform/declarative-resource-client-library v1.78.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.27.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.51.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.51.0 // indirect
	github.com/ProtonMail/go-crypto v1.1.6 // indirect
	github.com/YakDriver/go-version v0.1.0 // indirect
	github.com/YakDriver/regexache v0.24.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/apparentlymart/go-versions v1.0.1 // indirect
	github.com/aws/aws-sdk-go v1.55.7 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.29 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.63 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.33 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.37.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/account v1.22.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v1.30.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.38.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/amp v1.31.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/amplify v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.36.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/appfabric v1.11.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/appflow v1.45.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/appintegrations v1.30.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.34.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.29.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationsignals v1.7.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.29.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.32.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/appstream v1.44.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/appsync v1.43.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.37.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.24.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/backup v1.40.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/bcmdataexports v1.7.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.26.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrockagent v1.36.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/billing v1.1.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/budgets v1.29.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/chatbot v1.9.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/chime v1.34.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines v1.21.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkvoice v1.20.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.21.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.28.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.23.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.57.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore v1.8.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.29.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.26.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.47.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.33.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.53.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecatalyst v1.17.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.27.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeconnections v1.5.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.29.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.24.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/codegurureviewer v1.29.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.39.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.29.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.26.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.28.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.49.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.35.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.41.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/connect v1.125.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/connectcases v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/controltower v1.20.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.28.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.46.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/costoptimizationhub v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/customerprofiles v1.44.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/databrew v1.33.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/dataexchange v1.33.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.25.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/datasync v1.45.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/datazone v1.25.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/detective v1.31.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.28.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/devopsguru v1.34.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/dlm v1.29.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdb v1.40.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.14.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/drs v1.30.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.41.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.31.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.27.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrcontainers v1.34.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.27.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.36.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/evidently v1.23.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/finspace v1.28.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/firehose v1.36.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/fis v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/fms v1.39.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.39.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/glacier v1.26.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.29.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/grafana v1.26.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/greengrass v1.27.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/groundstation v1.31.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.53.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.29.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.27.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.40.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector v1.25.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.34.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internetmonitor v1.20.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/invoicing v1.0.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/iot v1.62.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/iotanalytics v1.26.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/iotevents v1.27.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivs v1.42.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.16.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafka v1.38.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.22.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/kendra v1.55.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.16.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.25.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.31.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.27.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/kms v1.37.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.39.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/launchwizard v1.8.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice v1.28.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.49.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/licensemanager v1.29.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/location v1.43.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/lookoutmetrics v1.31.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/m2 v1.19.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.44.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.37.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.67.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/medialive v1.68.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.34.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackagev2 v1.20.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackagevod v1.34.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.25.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/mgn v1.32.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.33.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/neptunegraph v1.16.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.45.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.32.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkmonitor v1.7.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/oam v1.15.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.45.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.26.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/organizations v1.37.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/osis v1.14.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/outposts v1.48.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/paymentcryptography v1.16.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/pcaconnectorad v1.9.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/pcs v1.2.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.34.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2 v1.18.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/pipes v1.18.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/polly v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/pricing v1.32.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/qbusiness v1.21.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/qldb v1.25.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.83.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ram v1.29.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/rbin v1.21.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.31.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.25.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.45.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.29.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.16.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.27.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.25.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.16.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.28.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53profiles v1.4.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.25.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.21.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/rum v1.21.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3control v1.53.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3outposts v1.28.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3tables v1.1.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.177.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.12.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/schemas v1.28.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.34.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.55.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.19.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.24.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.32.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.30.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.34.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.25.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.42.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sfn v1.34.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.29.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/signer v1.26.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.33.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.56.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.26.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.34.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmquicksetup v1.3.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmsap v1.19.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.29.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/swf v1.27.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.31.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/taxsettings v1.7.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamquery v1.29.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.29.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.43.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.56.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.21.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.13.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/waf v1.25.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.25.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.56.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.34.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/worklink v1.23.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.52.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspacesweb v1.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/xray v1.30.13 // indirect
	github.com/beevik/etree v1.5.0 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/btubbs/datetime v0.1.0 // indirect
	github.com/cedar-policy/cedar-go v0.1.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.6.0 // indirect
	github.com/cncf/xds/go v0.0.0-20250501225837-2ac532fd4443 // indirect
	github.com/coreos/go-systemd v0.0.0-20190620071333-e64a0ec8b42a // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/dmarkham/enumer v1.5.10 // indirect
	github.com/envoyproxy/go-control-plane/envoy v1.32.4 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.2.1 // indirect
	github.com/evanphx/json-patch v4.12.0+incompatible // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gammazero/deque v0.0.0-20180920172122-f6adf94963e4 // indirect
	github.com/gammazero/workerpool v0.0.0-20181230203049-86a96b5d5d92 // indirect
	github.com/go-jose/go-jose/v4 v4.1.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/golang/glog v1.2.4 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-cpy v0.0.0-20211218193943-a9c933c06932 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.6 // indirect
	github.com/googleapis/gax-go/v2 v2.14.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.23.0 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.62 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2 v2.0.0-beta.63 // indirect
	github.com/hashicorp/awspolicyequivalence v1.7.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-azure-sdk v0.20220824.1090858 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.7.8 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.6.3 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/hc-install v0.9.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.23.0 // indirect
	github.com/hashicorp/terraform-json v0.25.0 // indirect
	github.com/hashicorp/terraform-plugin-framework v1.14.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-jsontypes v0.2.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timeouts v0.5.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timetypes v0.5.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.17.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/hashicorp/terraform-plugin-mux v0.18.0 // indirect
	github.com/hashicorp/terraform-plugin-testing v1.13.1 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.5 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/manicminer/hamilton v0.44.0 // indirect
	github.com/manicminer/hamilton-autorest v0.2.0 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20240118010651-0ba75a80ca38 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rickb777/date v1.12.5-0.20200422084442-6300e543c4d9 // indirect
	github.com/rickb777/plural v1.2.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spiffe/go-spiffe/v2 v2.5.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tombuildsstuff/giovanni v0.20.0 // indirect
	github.com/ulikunitz/xz v0.5.12 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/zeebo/errs v1.4.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/detectors/gcp v1.36.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.59.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.61.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.61.0 // indirect
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	go4.org/netipx v0.0.0-20231129151722-fdeea329fbba // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/exp v0.0.0-20240409090435-93d18d7e34b8 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/oauth2 v0.30.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Force an specific version if not the AWS provider does not compile
replace (
	github.com/hashicorp/aws-sdk-go-base v0.6.0 => github.com/hashicorp/aws-sdk-go-base v0.5.0
	github.com/hashicorp/terraform => github.com/cycloidio/terraform v1.1.9-cy
	github.com/hashicorp/terraform-plugin-log => github.com/gdavison/terraform-plugin-log v0.0.0-20230928191232-6c653d8ef8fb
	github.com/hashicorp/terraform-provider-aws => github.com/appcd-dev/terraform-provider-aws v0.0.0-20250224044323-5bfa3cc416f1
	// Fork of Azurerm that has the V2 of the SDK
	github.com/hashicorp/terraform-provider-azurerm => github.com/cycloidio/terraform-provider-azurerm v1.44.1-0.20230517144901-90a36c6b8ed4
	github.com/hashicorp/terraform-provider-google => github.com/hashicorp/terraform-provider-google v1.20.1-0.20250318170011-31ee137da177
)

tool (
	github.com/dmarkham/enumer
	github.com/golang/mock/mockgen
	golang.org/x/lint/golint
	golang.org/x/tools/cmd/goimports
)
