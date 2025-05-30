// Code generated by "enumer -type ResourceType -addprefix aws_ -transform snake -linecomment"; DO NOT EDIT.

package aws

import (
	"fmt"
	"strings"
)

const _ResourceTypeName = "aws_instanceaws_albaws_api_gateway_deploymentaws_api_gateway_resourceaws_api_gateway_rest_apiaws_api_gateway_stageaws_athena_workgroupaws_autoscaling_groupaws_autoscaling_policyaws_autoscaling_scheduleaws_batch_job_definitionaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_cloudwatch_log_groupaws_dax_clusteraws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_directory_service_directoryaws_dms_replication_instanceaws_dx_gatewayaws_dynamodb_global_tableaws_dynamodb_tableaws_ebs_volumeaws_ecs_clusteraws_ecs_serviceaws_ecs_task_definitionaws_ec2_transit_gatewayaws_ec2_transit_gateway_vpc_attachmentaws_ec2_transit_gateway_route_tableaws_ec2_transit_gateway_multicast_domainaws_ec2_transit_gateway_peering_attachmentaws_ec2_transit_gateway_peering_attachment_accepteraws_ec2_transit_gateway_prefix_list_referenceaws_ec2_transit_gateway_routeaws_ec2_transit_gateway_route_table_associationaws_ec2_transit_gateway_route_table_propagationaws_ec2_transit_gateway_vpc_attachment_accepteraws_efs_file_systemaws_eipaws_eks_clusteraws_elasticache_clusteraws_elasticache_replication_groupaws_elastic_beanstalk_applicationaws_elasticsearch_domainaws_elasticsearch_domain_policyaws_elbaws_emr_clusteraws_fsx_lustre_file_systemaws_glue_catalog_databaseaws_glue_catalog_tableaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_internet_gatewayaws_key_pairaws_kinesis_streamaws_lambda_functionaws_launch_configurationaws_launch_templateaws_lbaws_lb_cookie_stickiness_policyaws_lb_listeneraws_lb_listener_certificateaws_lb_listener_ruleaws_lb_target_groupaws_lb_target_group_attachmentaws_lightsail_instanceaws_media_store_containeraws_mq_brokeraws_nat_gatewayaws_neptune_clusteraws_rds_clusteraws_rds_global_clusteraws_redshift_clusteraws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_route53_zoneaws_route53_zone_associationaws_route_tableaws_s3_bucketaws_security_groupaws_ses_active_receipt_rule_setaws_ses_configuration_setaws_ses_domain_dkimaws_ses_domain_identityaws_ses_domain_mail_fromaws_ses_identity_notification_topicaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_templateaws_sqs_queueaws_storagegateway_gatewayaws_subnetaws_volume_attachmentaws_vpcaws_vpc_endpointaws_vpc_peering_connectionaws_vpn_gateway"

var _ResourceTypeIndex = [...]uint16{0, 12, 19, 45, 69, 93, 114, 134, 155, 177, 201, 225, 252, 289, 314, 341, 365, 380, 395, 417, 436, 467, 495, 509, 534, 552, 566, 581, 596, 619, 642, 680, 715, 755, 797, 848, 893, 922, 969, 1016, 1063, 1082, 1089, 1104, 1127, 1160, 1193, 1217, 1248, 1255, 1270, 1296, 1321, 1343, 1361, 1382, 1413, 1426, 1450, 1470, 1501, 1525, 1556, 1570, 1582, 1601, 1631, 1652, 1678, 1690, 1719, 1738, 1768, 1788, 1808, 1820, 1838, 1857, 1881, 1900, 1906, 1937, 1952, 1979, 1999, 2018, 2048, 2070, 2095, 2108, 2123, 2142, 2157, 2179, 2199, 2225, 2249, 2270, 2288, 2317, 2354, 2370, 2398, 2413, 2426, 2444, 2475, 2500, 2519, 2542, 2566, 2601, 2623, 2643, 2667, 2683, 2696, 2722, 2732, 2753, 2760, 2776, 2802, 2817}

const _ResourceTypeLowerName = "aws_instanceaws_albaws_api_gateway_deploymentaws_api_gateway_resourceaws_api_gateway_rest_apiaws_api_gateway_stageaws_athena_workgroupaws_autoscaling_groupaws_autoscaling_policyaws_autoscaling_scheduleaws_batch_job_definitionaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_cloudwatch_log_groupaws_dax_clusteraws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_directory_service_directoryaws_dms_replication_instanceaws_dx_gatewayaws_dynamodb_global_tableaws_dynamodb_tableaws_ebs_volumeaws_ecs_clusteraws_ecs_serviceaws_ecs_task_definitionaws_ec2_transit_gatewayaws_ec2_transit_gateway_vpc_attachmentaws_ec2_transit_gateway_route_tableaws_ec2_transit_gateway_multicast_domainaws_ec2_transit_gateway_peering_attachmentaws_ec2_transit_gateway_peering_attachment_accepteraws_ec2_transit_gateway_prefix_list_referenceaws_ec2_transit_gateway_routeaws_ec2_transit_gateway_route_table_associationaws_ec2_transit_gateway_route_table_propagationaws_ec2_transit_gateway_vpc_attachment_accepteraws_efs_file_systemaws_eipaws_eks_clusteraws_elasticache_clusteraws_elasticache_replication_groupaws_elastic_beanstalk_applicationaws_elasticsearch_domainaws_elasticsearch_domain_policyaws_elbaws_emr_clusteraws_fsx_lustre_file_systemaws_glue_catalog_databaseaws_glue_catalog_tableaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_internet_gatewayaws_key_pairaws_kinesis_streamaws_lambda_functionaws_launch_configurationaws_launch_templateaws_lbaws_lb_cookie_stickiness_policyaws_lb_listeneraws_lb_listener_certificateaws_lb_listener_ruleaws_lb_target_groupaws_lb_target_group_attachmentaws_lightsail_instanceaws_media_store_containeraws_mq_brokeraws_nat_gatewayaws_neptune_clusteraws_rds_clusteraws_rds_global_clusteraws_redshift_clusteraws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_route53_zoneaws_route53_zone_associationaws_route_tableaws_s3_bucketaws_security_groupaws_ses_active_receipt_rule_setaws_ses_configuration_setaws_ses_domain_dkimaws_ses_domain_identityaws_ses_domain_mail_fromaws_ses_identity_notification_topicaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_templateaws_sqs_queueaws_storagegateway_gatewayaws_subnetaws_volume_attachmentaws_vpcaws_vpc_endpointaws_vpc_peering_connectionaws_vpn_gateway"

func (i ResourceType) String() string {
	i -= 1
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i+1)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ResourceTypeNoOp() {
	var x [1]struct{}
	_ = x[Instance-(1)]
	_ = x[ALB-(2)]
	_ = x[APIGatewayDeployment-(3)]
	_ = x[APIGatewayResource-(4)]
	_ = x[APIGatewayRestAPI-(5)]
	_ = x[APIGatewayStage-(6)]
	_ = x[AthenaWorkgroup-(7)]
	_ = x[AutoscalingGroup-(8)]
	_ = x[AutoscalingPolicy-(9)]
	_ = x[AutoscalingSchedule-(10)]
	_ = x[BatchJobDefinition-(11)]
	_ = x[CloudfrontDistribution-(12)]
	_ = x[CloudfrontOriginAccessIdentity-(13)]
	_ = x[CloudfrontPublicKey-(14)]
	_ = x[CloudwatchMetricAlarm-(15)]
	_ = x[CloudwatchLogGroup-(16)]
	_ = x[DaxCluster-(17)]
	_ = x[DBInstance-(18)]
	_ = x[DBParameterGroup-(19)]
	_ = x[DBSubnetGroup-(20)]
	_ = x[DirectoryServiceDirectory-(21)]
	_ = x[DmsReplicationInstance-(22)]
	_ = x[DXGateway-(23)]
	_ = x[DynamodbGlobalTable-(24)]
	_ = x[DynamodbTable-(25)]
	_ = x[EBSVolume-(26)]
	_ = x[ECSCluster-(27)]
	_ = x[ECSService-(28)]
	_ = x[ECSTaskDefinition-(29)]
	_ = x[EC2TransitGateway-(30)]
	_ = x[EC2TransitGatewayVPCAttachment-(31)]
	_ = x[EC2TransitGatewayRouteTable-(32)]
	_ = x[EC2TransitGatewayMulticastDomain-(33)]
	_ = x[EC2TransitGatewayPeeringAttachment-(34)]
	_ = x[EC2TransitGatewayPeeringAttachmentAccepter-(35)]
	_ = x[EC2TransitGatewayPrefixListReference-(36)]
	_ = x[EC2TransitGatewayRoute-(37)]
	_ = x[EC2TransitGatewayRouteTableAssociation-(38)]
	_ = x[EC2TransitGatewayRouteTablePropagation-(39)]
	_ = x[EC2TransitGatewayVPCAttachmentAccepter-(40)]
	_ = x[EFSFileSystem-(41)]
	_ = x[EIP-(42)]
	_ = x[EKSCluster-(43)]
	_ = x[ElasticacheCluster-(44)]
	_ = x[ElasticacheReplicationGroup-(45)]
	_ = x[ElasticBeanstalkApplication-(46)]
	_ = x[ElasticsearchDomain-(47)]
	_ = x[ElasticsearchDomainPolicy-(48)]
	_ = x[ELB-(49)]
	_ = x[EMRCluster-(50)]
	_ = x[FsxLustreFileSystem-(51)]
	_ = x[GlueCatalogDatabase-(52)]
	_ = x[GlueCatalogTable-(53)]
	_ = x[IAMAccessKey-(54)]
	_ = x[IAMAccountAlias-(55)]
	_ = x[IAMAccountPasswordPolicy-(56)]
	_ = x[IAMGroup-(57)]
	_ = x[IAMGroupMembership-(58)]
	_ = x[IAMGroupPolicy-(59)]
	_ = x[IAMGroupPolicyAttachment-(60)]
	_ = x[IAMInstanceProfile-(61)]
	_ = x[IAMOpenidConnectProvider-(62)]
	_ = x[IAMPolicy-(63)]
	_ = x[IAMRole-(64)]
	_ = x[IAMRolePolicy-(65)]
	_ = x[IAMRolePolicyAttachment-(66)]
	_ = x[IAMSAMLProvider-(67)]
	_ = x[IAMServerCertificate-(68)]
	_ = x[IAMUser-(69)]
	_ = x[IAMUserGroupMembership-(70)]
	_ = x[IAMUserPolicy-(71)]
	_ = x[IAMUserPolicyAttachment-(72)]
	_ = x[IAMUserSSHKey-(73)]
	_ = x[InternetGateway-(74)]
	_ = x[KeyPair-(75)]
	_ = x[KinesisStream-(76)]
	_ = x[LambdaFunction-(77)]
	_ = x[LaunchConfiguration-(78)]
	_ = x[LaunchTemplate-(79)]
	_ = x[LB-(80)]
	_ = x[LBCookieStickinessPolicy-(81)]
	_ = x[LBListener-(82)]
	_ = x[LBListenerCertificate-(83)]
	_ = x[LBListenerRule-(84)]
	_ = x[LBTargetGroup-(85)]
	_ = x[LBTargetGroupAttachment-(86)]
	_ = x[LightsailInstance-(87)]
	_ = x[MediaStoreContainer-(88)]
	_ = x[MQBroker-(89)]
	_ = x[NatGateway-(90)]
	_ = x[NeptuneCluster-(91)]
	_ = x[RDSCluster-(92)]
	_ = x[RDSGlobalCluster-(93)]
	_ = x[RedshiftCluster-(94)]
	_ = x[Route53DelegationSet-(95)]
	_ = x[Route53HealthCheck-(96)]
	_ = x[Route53QueryLog-(97)]
	_ = x[Route53Record-(98)]
	_ = x[Route53ResolverEndpoint-(99)]
	_ = x[Route53ResolverRuleAssociation-(100)]
	_ = x[Route53Zone-(101)]
	_ = x[Route53ZoneAssociation-(102)]
	_ = x[RouteTable-(103)]
	_ = x[S3Bucket-(104)]
	_ = x[SecurityGroup-(105)]
	_ = x[SESActiveReceiptRuleSet-(106)]
	_ = x[SESConfigurationSet-(107)]
	_ = x[SESDomainDKIM-(108)]
	_ = x[SESDomainIdentity-(109)]
	_ = x[SESDomainMailFrom-(110)]
	_ = x[SESIdentityNotificationTopic-(111)]
	_ = x[SESReceiptFilter-(112)]
	_ = x[SESReceiptRule-(113)]
	_ = x[SESReceiptRuleSet-(114)]
	_ = x[SESTemplate-(115)]
	_ = x[SQSQueue-(116)]
	_ = x[StoragegatewayGateway-(117)]
	_ = x[Subnet-(118)]
	_ = x[VolumeAttachment-(119)]
	_ = x[VPC-(120)]
	_ = x[VPCEndpoint-(121)]
	_ = x[VPCPeeringConnection-(122)]
	_ = x[VPNGateway-(123)]
}

var _ResourceTypeValues = []ResourceType{Instance, ALB, APIGatewayDeployment, APIGatewayResource, APIGatewayRestAPI, APIGatewayStage, AthenaWorkgroup, AutoscalingGroup, AutoscalingPolicy, AutoscalingSchedule, BatchJobDefinition, CloudfrontDistribution, CloudfrontOriginAccessIdentity, CloudfrontPublicKey, CloudwatchMetricAlarm, CloudwatchLogGroup, DaxCluster, DBInstance, DBParameterGroup, DBSubnetGroup, DirectoryServiceDirectory, DmsReplicationInstance, DXGateway, DynamodbGlobalTable, DynamodbTable, EBSVolume, ECSCluster, ECSService, ECSTaskDefinition, EC2TransitGateway, EC2TransitGatewayVPCAttachment, EC2TransitGatewayRouteTable, EC2TransitGatewayMulticastDomain, EC2TransitGatewayPeeringAttachment, EC2TransitGatewayPeeringAttachmentAccepter, EC2TransitGatewayPrefixListReference, EC2TransitGatewayRoute, EC2TransitGatewayRouteTableAssociation, EC2TransitGatewayRouteTablePropagation, EC2TransitGatewayVPCAttachmentAccepter, EFSFileSystem, EIP, EKSCluster, ElasticacheCluster, ElasticacheReplicationGroup, ElasticBeanstalkApplication, ElasticsearchDomain, ElasticsearchDomainPolicy, ELB, EMRCluster, FsxLustreFileSystem, GlueCatalogDatabase, GlueCatalogTable, IAMAccessKey, IAMAccountAlias, IAMAccountPasswordPolicy, IAMGroup, IAMGroupMembership, IAMGroupPolicy, IAMGroupPolicyAttachment, IAMInstanceProfile, IAMOpenidConnectProvider, IAMPolicy, IAMRole, IAMRolePolicy, IAMRolePolicyAttachment, IAMSAMLProvider, IAMServerCertificate, IAMUser, IAMUserGroupMembership, IAMUserPolicy, IAMUserPolicyAttachment, IAMUserSSHKey, InternetGateway, KeyPair, KinesisStream, LambdaFunction, LaunchConfiguration, LaunchTemplate, LB, LBCookieStickinessPolicy, LBListener, LBListenerCertificate, LBListenerRule, LBTargetGroup, LBTargetGroupAttachment, LightsailInstance, MediaStoreContainer, MQBroker, NatGateway, NeptuneCluster, RDSCluster, RDSGlobalCluster, RedshiftCluster, Route53DelegationSet, Route53HealthCheck, Route53QueryLog, Route53Record, Route53ResolverEndpoint, Route53ResolverRuleAssociation, Route53Zone, Route53ZoneAssociation, RouteTable, S3Bucket, SecurityGroup, SESActiveReceiptRuleSet, SESConfigurationSet, SESDomainDKIM, SESDomainIdentity, SESDomainMailFrom, SESIdentityNotificationTopic, SESReceiptFilter, SESReceiptRule, SESReceiptRuleSet, SESTemplate, SQSQueue, StoragegatewayGateway, Subnet, VolumeAttachment, VPC, VPCEndpoint, VPCPeeringConnection, VPNGateway}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:12]:           Instance,
	_ResourceTypeLowerName[0:12]:      Instance,
	_ResourceTypeName[12:19]:          ALB,
	_ResourceTypeLowerName[12:19]:     ALB,
	_ResourceTypeName[19:45]:          APIGatewayDeployment,
	_ResourceTypeLowerName[19:45]:     APIGatewayDeployment,
	_ResourceTypeName[45:69]:          APIGatewayResource,
	_ResourceTypeLowerName[45:69]:     APIGatewayResource,
	_ResourceTypeName[69:93]:          APIGatewayRestAPI,
	_ResourceTypeLowerName[69:93]:     APIGatewayRestAPI,
	_ResourceTypeName[93:114]:         APIGatewayStage,
	_ResourceTypeLowerName[93:114]:    APIGatewayStage,
	_ResourceTypeName[114:134]:        AthenaWorkgroup,
	_ResourceTypeLowerName[114:134]:   AthenaWorkgroup,
	_ResourceTypeName[134:155]:        AutoscalingGroup,
	_ResourceTypeLowerName[134:155]:   AutoscalingGroup,
	_ResourceTypeName[155:177]:        AutoscalingPolicy,
	_ResourceTypeLowerName[155:177]:   AutoscalingPolicy,
	_ResourceTypeName[177:201]:        AutoscalingSchedule,
	_ResourceTypeLowerName[177:201]:   AutoscalingSchedule,
	_ResourceTypeName[201:225]:        BatchJobDefinition,
	_ResourceTypeLowerName[201:225]:   BatchJobDefinition,
	_ResourceTypeName[225:252]:        CloudfrontDistribution,
	_ResourceTypeLowerName[225:252]:   CloudfrontDistribution,
	_ResourceTypeName[252:289]:        CloudfrontOriginAccessIdentity,
	_ResourceTypeLowerName[252:289]:   CloudfrontOriginAccessIdentity,
	_ResourceTypeName[289:314]:        CloudfrontPublicKey,
	_ResourceTypeLowerName[289:314]:   CloudfrontPublicKey,
	_ResourceTypeName[314:341]:        CloudwatchMetricAlarm,
	_ResourceTypeLowerName[314:341]:   CloudwatchMetricAlarm,
	_ResourceTypeName[341:365]:        CloudwatchLogGroup,
	_ResourceTypeLowerName[341:365]:   CloudwatchLogGroup,
	_ResourceTypeName[365:380]:        DaxCluster,
	_ResourceTypeLowerName[365:380]:   DaxCluster,
	_ResourceTypeName[380:395]:        DBInstance,
	_ResourceTypeLowerName[380:395]:   DBInstance,
	_ResourceTypeName[395:417]:        DBParameterGroup,
	_ResourceTypeLowerName[395:417]:   DBParameterGroup,
	_ResourceTypeName[417:436]:        DBSubnetGroup,
	_ResourceTypeLowerName[417:436]:   DBSubnetGroup,
	_ResourceTypeName[436:467]:        DirectoryServiceDirectory,
	_ResourceTypeLowerName[436:467]:   DirectoryServiceDirectory,
	_ResourceTypeName[467:495]:        DmsReplicationInstance,
	_ResourceTypeLowerName[467:495]:   DmsReplicationInstance,
	_ResourceTypeName[495:509]:        DXGateway,
	_ResourceTypeLowerName[495:509]:   DXGateway,
	_ResourceTypeName[509:534]:        DynamodbGlobalTable,
	_ResourceTypeLowerName[509:534]:   DynamodbGlobalTable,
	_ResourceTypeName[534:552]:        DynamodbTable,
	_ResourceTypeLowerName[534:552]:   DynamodbTable,
	_ResourceTypeName[552:566]:        EBSVolume,
	_ResourceTypeLowerName[552:566]:   EBSVolume,
	_ResourceTypeName[566:581]:        ECSCluster,
	_ResourceTypeLowerName[566:581]:   ECSCluster,
	_ResourceTypeName[581:596]:        ECSService,
	_ResourceTypeLowerName[581:596]:   ECSService,
	_ResourceTypeName[596:619]:        ECSTaskDefinition,
	_ResourceTypeLowerName[596:619]:   ECSTaskDefinition,
	_ResourceTypeName[619:642]:        EC2TransitGateway,
	_ResourceTypeLowerName[619:642]:   EC2TransitGateway,
	_ResourceTypeName[642:680]:        EC2TransitGatewayVPCAttachment,
	_ResourceTypeLowerName[642:680]:   EC2TransitGatewayVPCAttachment,
	_ResourceTypeName[680:715]:        EC2TransitGatewayRouteTable,
	_ResourceTypeLowerName[680:715]:   EC2TransitGatewayRouteTable,
	_ResourceTypeName[715:755]:        EC2TransitGatewayMulticastDomain,
	_ResourceTypeLowerName[715:755]:   EC2TransitGatewayMulticastDomain,
	_ResourceTypeName[755:797]:        EC2TransitGatewayPeeringAttachment,
	_ResourceTypeLowerName[755:797]:   EC2TransitGatewayPeeringAttachment,
	_ResourceTypeName[797:848]:        EC2TransitGatewayPeeringAttachmentAccepter,
	_ResourceTypeLowerName[797:848]:   EC2TransitGatewayPeeringAttachmentAccepter,
	_ResourceTypeName[848:893]:        EC2TransitGatewayPrefixListReference,
	_ResourceTypeLowerName[848:893]:   EC2TransitGatewayPrefixListReference,
	_ResourceTypeName[893:922]:        EC2TransitGatewayRoute,
	_ResourceTypeLowerName[893:922]:   EC2TransitGatewayRoute,
	_ResourceTypeName[922:969]:        EC2TransitGatewayRouteTableAssociation,
	_ResourceTypeLowerName[922:969]:   EC2TransitGatewayRouteTableAssociation,
	_ResourceTypeName[969:1016]:       EC2TransitGatewayRouteTablePropagation,
	_ResourceTypeLowerName[969:1016]:  EC2TransitGatewayRouteTablePropagation,
	_ResourceTypeName[1016:1063]:      EC2TransitGatewayVPCAttachmentAccepter,
	_ResourceTypeLowerName[1016:1063]: EC2TransitGatewayVPCAttachmentAccepter,
	_ResourceTypeName[1063:1082]:      EFSFileSystem,
	_ResourceTypeLowerName[1063:1082]: EFSFileSystem,
	_ResourceTypeName[1082:1089]:      EIP,
	_ResourceTypeLowerName[1082:1089]: EIP,
	_ResourceTypeName[1089:1104]:      EKSCluster,
	_ResourceTypeLowerName[1089:1104]: EKSCluster,
	_ResourceTypeName[1104:1127]:      ElasticacheCluster,
	_ResourceTypeLowerName[1104:1127]: ElasticacheCluster,
	_ResourceTypeName[1127:1160]:      ElasticacheReplicationGroup,
	_ResourceTypeLowerName[1127:1160]: ElasticacheReplicationGroup,
	_ResourceTypeName[1160:1193]:      ElasticBeanstalkApplication,
	_ResourceTypeLowerName[1160:1193]: ElasticBeanstalkApplication,
	_ResourceTypeName[1193:1217]:      ElasticsearchDomain,
	_ResourceTypeLowerName[1193:1217]: ElasticsearchDomain,
	_ResourceTypeName[1217:1248]:      ElasticsearchDomainPolicy,
	_ResourceTypeLowerName[1217:1248]: ElasticsearchDomainPolicy,
	_ResourceTypeName[1248:1255]:      ELB,
	_ResourceTypeLowerName[1248:1255]: ELB,
	_ResourceTypeName[1255:1270]:      EMRCluster,
	_ResourceTypeLowerName[1255:1270]: EMRCluster,
	_ResourceTypeName[1270:1296]:      FsxLustreFileSystem,
	_ResourceTypeLowerName[1270:1296]: FsxLustreFileSystem,
	_ResourceTypeName[1296:1321]:      GlueCatalogDatabase,
	_ResourceTypeLowerName[1296:1321]: GlueCatalogDatabase,
	_ResourceTypeName[1321:1343]:      GlueCatalogTable,
	_ResourceTypeLowerName[1321:1343]: GlueCatalogTable,
	_ResourceTypeName[1343:1361]:      IAMAccessKey,
	_ResourceTypeLowerName[1343:1361]: IAMAccessKey,
	_ResourceTypeName[1361:1382]:      IAMAccountAlias,
	_ResourceTypeLowerName[1361:1382]: IAMAccountAlias,
	_ResourceTypeName[1382:1413]:      IAMAccountPasswordPolicy,
	_ResourceTypeLowerName[1382:1413]: IAMAccountPasswordPolicy,
	_ResourceTypeName[1413:1426]:      IAMGroup,
	_ResourceTypeLowerName[1413:1426]: IAMGroup,
	_ResourceTypeName[1426:1450]:      IAMGroupMembership,
	_ResourceTypeLowerName[1426:1450]: IAMGroupMembership,
	_ResourceTypeName[1450:1470]:      IAMGroupPolicy,
	_ResourceTypeLowerName[1450:1470]: IAMGroupPolicy,
	_ResourceTypeName[1470:1501]:      IAMGroupPolicyAttachment,
	_ResourceTypeLowerName[1470:1501]: IAMGroupPolicyAttachment,
	_ResourceTypeName[1501:1525]:      IAMInstanceProfile,
	_ResourceTypeLowerName[1501:1525]: IAMInstanceProfile,
	_ResourceTypeName[1525:1556]:      IAMOpenidConnectProvider,
	_ResourceTypeLowerName[1525:1556]: IAMOpenidConnectProvider,
	_ResourceTypeName[1556:1570]:      IAMPolicy,
	_ResourceTypeLowerName[1556:1570]: IAMPolicy,
	_ResourceTypeName[1570:1582]:      IAMRole,
	_ResourceTypeLowerName[1570:1582]: IAMRole,
	_ResourceTypeName[1582:1601]:      IAMRolePolicy,
	_ResourceTypeLowerName[1582:1601]: IAMRolePolicy,
	_ResourceTypeName[1601:1631]:      IAMRolePolicyAttachment,
	_ResourceTypeLowerName[1601:1631]: IAMRolePolicyAttachment,
	_ResourceTypeName[1631:1652]:      IAMSAMLProvider,
	_ResourceTypeLowerName[1631:1652]: IAMSAMLProvider,
	_ResourceTypeName[1652:1678]:      IAMServerCertificate,
	_ResourceTypeLowerName[1652:1678]: IAMServerCertificate,
	_ResourceTypeName[1678:1690]:      IAMUser,
	_ResourceTypeLowerName[1678:1690]: IAMUser,
	_ResourceTypeName[1690:1719]:      IAMUserGroupMembership,
	_ResourceTypeLowerName[1690:1719]: IAMUserGroupMembership,
	_ResourceTypeName[1719:1738]:      IAMUserPolicy,
	_ResourceTypeLowerName[1719:1738]: IAMUserPolicy,
	_ResourceTypeName[1738:1768]:      IAMUserPolicyAttachment,
	_ResourceTypeLowerName[1738:1768]: IAMUserPolicyAttachment,
	_ResourceTypeName[1768:1788]:      IAMUserSSHKey,
	_ResourceTypeLowerName[1768:1788]: IAMUserSSHKey,
	_ResourceTypeName[1788:1808]:      InternetGateway,
	_ResourceTypeLowerName[1788:1808]: InternetGateway,
	_ResourceTypeName[1808:1820]:      KeyPair,
	_ResourceTypeLowerName[1808:1820]: KeyPair,
	_ResourceTypeName[1820:1838]:      KinesisStream,
	_ResourceTypeLowerName[1820:1838]: KinesisStream,
	_ResourceTypeName[1838:1857]:      LambdaFunction,
	_ResourceTypeLowerName[1838:1857]: LambdaFunction,
	_ResourceTypeName[1857:1881]:      LaunchConfiguration,
	_ResourceTypeLowerName[1857:1881]: LaunchConfiguration,
	_ResourceTypeName[1881:1900]:      LaunchTemplate,
	_ResourceTypeLowerName[1881:1900]: LaunchTemplate,
	_ResourceTypeName[1900:1906]:      LB,
	_ResourceTypeLowerName[1900:1906]: LB,
	_ResourceTypeName[1906:1937]:      LBCookieStickinessPolicy,
	_ResourceTypeLowerName[1906:1937]: LBCookieStickinessPolicy,
	_ResourceTypeName[1937:1952]:      LBListener,
	_ResourceTypeLowerName[1937:1952]: LBListener,
	_ResourceTypeName[1952:1979]:      LBListenerCertificate,
	_ResourceTypeLowerName[1952:1979]: LBListenerCertificate,
	_ResourceTypeName[1979:1999]:      LBListenerRule,
	_ResourceTypeLowerName[1979:1999]: LBListenerRule,
	_ResourceTypeName[1999:2018]:      LBTargetGroup,
	_ResourceTypeLowerName[1999:2018]: LBTargetGroup,
	_ResourceTypeName[2018:2048]:      LBTargetGroupAttachment,
	_ResourceTypeLowerName[2018:2048]: LBTargetGroupAttachment,
	_ResourceTypeName[2048:2070]:      LightsailInstance,
	_ResourceTypeLowerName[2048:2070]: LightsailInstance,
	_ResourceTypeName[2070:2095]:      MediaStoreContainer,
	_ResourceTypeLowerName[2070:2095]: MediaStoreContainer,
	_ResourceTypeName[2095:2108]:      MQBroker,
	_ResourceTypeLowerName[2095:2108]: MQBroker,
	_ResourceTypeName[2108:2123]:      NatGateway,
	_ResourceTypeLowerName[2108:2123]: NatGateway,
	_ResourceTypeName[2123:2142]:      NeptuneCluster,
	_ResourceTypeLowerName[2123:2142]: NeptuneCluster,
	_ResourceTypeName[2142:2157]:      RDSCluster,
	_ResourceTypeLowerName[2142:2157]: RDSCluster,
	_ResourceTypeName[2157:2179]:      RDSGlobalCluster,
	_ResourceTypeLowerName[2157:2179]: RDSGlobalCluster,
	_ResourceTypeName[2179:2199]:      RedshiftCluster,
	_ResourceTypeLowerName[2179:2199]: RedshiftCluster,
	_ResourceTypeName[2199:2225]:      Route53DelegationSet,
	_ResourceTypeLowerName[2199:2225]: Route53DelegationSet,
	_ResourceTypeName[2225:2249]:      Route53HealthCheck,
	_ResourceTypeLowerName[2225:2249]: Route53HealthCheck,
	_ResourceTypeName[2249:2270]:      Route53QueryLog,
	_ResourceTypeLowerName[2249:2270]: Route53QueryLog,
	_ResourceTypeName[2270:2288]:      Route53Record,
	_ResourceTypeLowerName[2270:2288]: Route53Record,
	_ResourceTypeName[2288:2317]:      Route53ResolverEndpoint,
	_ResourceTypeLowerName[2288:2317]: Route53ResolverEndpoint,
	_ResourceTypeName[2317:2354]:      Route53ResolverRuleAssociation,
	_ResourceTypeLowerName[2317:2354]: Route53ResolverRuleAssociation,
	_ResourceTypeName[2354:2370]:      Route53Zone,
	_ResourceTypeLowerName[2354:2370]: Route53Zone,
	_ResourceTypeName[2370:2398]:      Route53ZoneAssociation,
	_ResourceTypeLowerName[2370:2398]: Route53ZoneAssociation,
	_ResourceTypeName[2398:2413]:      RouteTable,
	_ResourceTypeLowerName[2398:2413]: RouteTable,
	_ResourceTypeName[2413:2426]:      S3Bucket,
	_ResourceTypeLowerName[2413:2426]: S3Bucket,
	_ResourceTypeName[2426:2444]:      SecurityGroup,
	_ResourceTypeLowerName[2426:2444]: SecurityGroup,
	_ResourceTypeName[2444:2475]:      SESActiveReceiptRuleSet,
	_ResourceTypeLowerName[2444:2475]: SESActiveReceiptRuleSet,
	_ResourceTypeName[2475:2500]:      SESConfigurationSet,
	_ResourceTypeLowerName[2475:2500]: SESConfigurationSet,
	_ResourceTypeName[2500:2519]:      SESDomainDKIM,
	_ResourceTypeLowerName[2500:2519]: SESDomainDKIM,
	_ResourceTypeName[2519:2542]:      SESDomainIdentity,
	_ResourceTypeLowerName[2519:2542]: SESDomainIdentity,
	_ResourceTypeName[2542:2566]:      SESDomainMailFrom,
	_ResourceTypeLowerName[2542:2566]: SESDomainMailFrom,
	_ResourceTypeName[2566:2601]:      SESIdentityNotificationTopic,
	_ResourceTypeLowerName[2566:2601]: SESIdentityNotificationTopic,
	_ResourceTypeName[2601:2623]:      SESReceiptFilter,
	_ResourceTypeLowerName[2601:2623]: SESReceiptFilter,
	_ResourceTypeName[2623:2643]:      SESReceiptRule,
	_ResourceTypeLowerName[2623:2643]: SESReceiptRule,
	_ResourceTypeName[2643:2667]:      SESReceiptRuleSet,
	_ResourceTypeLowerName[2643:2667]: SESReceiptRuleSet,
	_ResourceTypeName[2667:2683]:      SESTemplate,
	_ResourceTypeLowerName[2667:2683]: SESTemplate,
	_ResourceTypeName[2683:2696]:      SQSQueue,
	_ResourceTypeLowerName[2683:2696]: SQSQueue,
	_ResourceTypeName[2696:2722]:      StoragegatewayGateway,
	_ResourceTypeLowerName[2696:2722]: StoragegatewayGateway,
	_ResourceTypeName[2722:2732]:      Subnet,
	_ResourceTypeLowerName[2722:2732]: Subnet,
	_ResourceTypeName[2732:2753]:      VolumeAttachment,
	_ResourceTypeLowerName[2732:2753]: VolumeAttachment,
	_ResourceTypeName[2753:2760]:      VPC,
	_ResourceTypeLowerName[2753:2760]: VPC,
	_ResourceTypeName[2760:2776]:      VPCEndpoint,
	_ResourceTypeLowerName[2760:2776]: VPCEndpoint,
	_ResourceTypeName[2776:2802]:      VPCPeeringConnection,
	_ResourceTypeLowerName[2776:2802]: VPCPeeringConnection,
	_ResourceTypeName[2802:2817]:      VPNGateway,
	_ResourceTypeLowerName[2802:2817]: VPNGateway,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:12],
	_ResourceTypeName[12:19],
	_ResourceTypeName[19:45],
	_ResourceTypeName[45:69],
	_ResourceTypeName[69:93],
	_ResourceTypeName[93:114],
	_ResourceTypeName[114:134],
	_ResourceTypeName[134:155],
	_ResourceTypeName[155:177],
	_ResourceTypeName[177:201],
	_ResourceTypeName[201:225],
	_ResourceTypeName[225:252],
	_ResourceTypeName[252:289],
	_ResourceTypeName[289:314],
	_ResourceTypeName[314:341],
	_ResourceTypeName[341:365],
	_ResourceTypeName[365:380],
	_ResourceTypeName[380:395],
	_ResourceTypeName[395:417],
	_ResourceTypeName[417:436],
	_ResourceTypeName[436:467],
	_ResourceTypeName[467:495],
	_ResourceTypeName[495:509],
	_ResourceTypeName[509:534],
	_ResourceTypeName[534:552],
	_ResourceTypeName[552:566],
	_ResourceTypeName[566:581],
	_ResourceTypeName[581:596],
	_ResourceTypeName[596:619],
	_ResourceTypeName[619:642],
	_ResourceTypeName[642:680],
	_ResourceTypeName[680:715],
	_ResourceTypeName[715:755],
	_ResourceTypeName[755:797],
	_ResourceTypeName[797:848],
	_ResourceTypeName[848:893],
	_ResourceTypeName[893:922],
	_ResourceTypeName[922:969],
	_ResourceTypeName[969:1016],
	_ResourceTypeName[1016:1063],
	_ResourceTypeName[1063:1082],
	_ResourceTypeName[1082:1089],
	_ResourceTypeName[1089:1104],
	_ResourceTypeName[1104:1127],
	_ResourceTypeName[1127:1160],
	_ResourceTypeName[1160:1193],
	_ResourceTypeName[1193:1217],
	_ResourceTypeName[1217:1248],
	_ResourceTypeName[1248:1255],
	_ResourceTypeName[1255:1270],
	_ResourceTypeName[1270:1296],
	_ResourceTypeName[1296:1321],
	_ResourceTypeName[1321:1343],
	_ResourceTypeName[1343:1361],
	_ResourceTypeName[1361:1382],
	_ResourceTypeName[1382:1413],
	_ResourceTypeName[1413:1426],
	_ResourceTypeName[1426:1450],
	_ResourceTypeName[1450:1470],
	_ResourceTypeName[1470:1501],
	_ResourceTypeName[1501:1525],
	_ResourceTypeName[1525:1556],
	_ResourceTypeName[1556:1570],
	_ResourceTypeName[1570:1582],
	_ResourceTypeName[1582:1601],
	_ResourceTypeName[1601:1631],
	_ResourceTypeName[1631:1652],
	_ResourceTypeName[1652:1678],
	_ResourceTypeName[1678:1690],
	_ResourceTypeName[1690:1719],
	_ResourceTypeName[1719:1738],
	_ResourceTypeName[1738:1768],
	_ResourceTypeName[1768:1788],
	_ResourceTypeName[1788:1808],
	_ResourceTypeName[1808:1820],
	_ResourceTypeName[1820:1838],
	_ResourceTypeName[1838:1857],
	_ResourceTypeName[1857:1881],
	_ResourceTypeName[1881:1900],
	_ResourceTypeName[1900:1906],
	_ResourceTypeName[1906:1937],
	_ResourceTypeName[1937:1952],
	_ResourceTypeName[1952:1979],
	_ResourceTypeName[1979:1999],
	_ResourceTypeName[1999:2018],
	_ResourceTypeName[2018:2048],
	_ResourceTypeName[2048:2070],
	_ResourceTypeName[2070:2095],
	_ResourceTypeName[2095:2108],
	_ResourceTypeName[2108:2123],
	_ResourceTypeName[2123:2142],
	_ResourceTypeName[2142:2157],
	_ResourceTypeName[2157:2179],
	_ResourceTypeName[2179:2199],
	_ResourceTypeName[2199:2225],
	_ResourceTypeName[2225:2249],
	_ResourceTypeName[2249:2270],
	_ResourceTypeName[2270:2288],
	_ResourceTypeName[2288:2317],
	_ResourceTypeName[2317:2354],
	_ResourceTypeName[2354:2370],
	_ResourceTypeName[2370:2398],
	_ResourceTypeName[2398:2413],
	_ResourceTypeName[2413:2426],
	_ResourceTypeName[2426:2444],
	_ResourceTypeName[2444:2475],
	_ResourceTypeName[2475:2500],
	_ResourceTypeName[2500:2519],
	_ResourceTypeName[2519:2542],
	_ResourceTypeName[2542:2566],
	_ResourceTypeName[2566:2601],
	_ResourceTypeName[2601:2623],
	_ResourceTypeName[2623:2643],
	_ResourceTypeName[2643:2667],
	_ResourceTypeName[2667:2683],
	_ResourceTypeName[2683:2696],
	_ResourceTypeName[2696:2722],
	_ResourceTypeName[2722:2732],
	_ResourceTypeName[2732:2753],
	_ResourceTypeName[2753:2760],
	_ResourceTypeName[2760:2776],
	_ResourceTypeName[2776:2802],
	_ResourceTypeName[2802:2817],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ResourceTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
