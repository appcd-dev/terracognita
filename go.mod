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
	github.com/aws/aws-sdk-go-v2/service/configservice v1.51.13
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.49.1
	github.com/aws/aws-sdk-go-v2/service/dax v1.23.16
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.30.13
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.30.18
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.40.2
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.203.1
	github.com/aws/aws-sdk-go-v2/service/ecs v1.53.16
	github.com/aws/aws-sdk-go-v2/service/efs v1.34.12
	github.com/aws/aws-sdk-go-v2/service/eks v1.58.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.44.13
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
	github.com/aws/aws-sdk-go-v2/service/rds v1.93.13
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
	github.com/gertd/go-pluralize v0.1.7
	github.com/go-kit/kit v0.9.0
	github.com/golang/mock v1.6.0
	github.com/hashicorp/go-azure-helpers v0.40.0
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/hcl/v2 v2.13.0
	github.com/hashicorp/terraform v0.13.0
	github.com/hashicorp/terraform-plugin-go v0.10.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.18.0
	github.com/hashicorp/terraform-provider-aws v1.60.1-0.20210513231836-489654890359
	github.com/hashicorp/terraform-provider-azurerm v1.44.1-0.20201029183808-d721bcc1bb55
	github.com/hashicorp/terraform-provider-google v1.20.1-0.20210510171431-a764cf3da527
	github.com/hashicorp/terraform-provider-vsphere v1.26.1-0.20220510172607-30f37d268d79
	github.com/pascaldekloe/name v1.0.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.2
	github.com/vmware/govmomi v0.28.0
	github.com/zclconf/go-cty v1.10.0
	golang.org/x/sync v0.11.0
	golang.org/x/text v0.22.0
	google.golang.org/api v0.61.0
	google.golang.org/grpc v1.47.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	bitbucket.org/creachadair/stringset v0.0.8 // indirect
	cloud.google.com/go v0.97.0 // indirect
	cloud.google.com/go/bigtable v1.10.1 // indirect
	cloud.google.com/go/storage v1.16.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.5 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/GoogleCloudPlatform/declarative-resource-client-library v0.0.0-20220114025148-a9879027a727 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-versions v1.0.1 // indirect
	github.com/aws/aws-sdk-go v1.43.34 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.33 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.28.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.15 // indirect
	github.com/beevik/etree v1.1.0 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/btubbs/datetime v0.1.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/census-instrumentation/opencensus-proto v0.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cncf/udpa/go v0.0.0-20210930031921-04548b0d99d4 // indirect
	github.com/cncf/xds/go v0.0.0-20211011173535-cb28da3451f1 // indirect
	github.com/coreos/go-systemd v0.0.0-20190620071333-e64a0ec8b42a // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/dmarkham/enumer v1.5.10 // indirect
	github.com/envoyproxy/go-control-plane v0.10.2-0.20220325020618-49ff273808a1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0 // indirect
	github.com/evanphx/json-patch v4.12.0+incompatible // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gammazero/deque v0.0.0-20180920172122-f6adf94963e4 // indirect
	github.com/gammazero/workerpool v0.0.0-20181230203049-86a96b5d5d92 // indirect
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.4.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.16.0 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.14 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2 v2.0.0-beta.15 // indirect
	github.com/hashicorp/awspolicyequivalence v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-azure-sdk v0.20220824.1090858 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.5.11 // indirect
	github.com/hashicorp/go-hclog v1.2.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.4 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hc-install v0.4.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.17.2 // indirect
	github.com/hashicorp/terraform-json v0.14.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.4.1 // indirect
	github.com/hashicorp/terraform-registry-address v0.0.0-20220623143253-7d51757b572c // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/hashicorp/yamux v0.0.0-20210316155119-a95892c5f864 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.13.1 // indirect
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/manicminer/hamilton v0.44.0 // indirect
	github.com/manicminer/hamilton-autorest v0.2.0 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rickb777/date v1.12.5-0.20200422084442-6300e543c4d9 // indirect
	github.com/rickb777/plural v1.2.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tombuildsstuff/giovanni v0.20.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Force an specific version if not the AWS provider does not compile
replace (
	github.com/hashicorp/aws-sdk-go-base v0.6.0 => github.com/hashicorp/aws-sdk-go-base v0.5.0
	// If we  go to the 1.5.0 then github.com/hashicorp/terraform-plugin-test/ will break
	// as go-getter introduced a break from 1.4 -> 1.5
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform => github.com/cycloidio/terraform v1.1.9-cy
	github.com/hashicorp/terraform-provider-aws => github.com/cycloidio/terraform-provider-aws v1.60.1-0.20220513132327-e2dbdf90e533
	// Fork of Azurerm that has the V2 of the SDK
	github.com/hashicorp/terraform-provider-azurerm => github.com/cycloidio/terraform-provider-azurerm v1.44.1-0.20230517144901-90a36c6b8ed4
	github.com/hashicorp/terraform-provider-google => github.com/hashicorp/terraform-provider-google v1.20.1-0.20220201002249-bc5fcb3c89a5
)

tool (
	github.com/dmarkham/enumer
	github.com/golang/mock/mockgen
	golang.org/x/tools/cmd/goimports
)
