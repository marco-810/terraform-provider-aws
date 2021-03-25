#
# For future consideration, this list could be automatically generated
# via the AWS SDK service list.
#

variable "service_labels" {
  default = [
    "accessanalyzer",
    "acm",
    "acmpca",
    "alexaforbusiness",
    "amplify",
    "apigateway",
    "apigatewaymanagementapi",
    "apigatewayv2",
    "appconfig",
    "appflow",
    "applicationautoscaling",
    "applicationdiscoveryservice",
    "applicationinsights",
    "appmesh",
    "appstream",
    "appsync",
    "athena",
    "auditmanager",
    "autoscaling",
    "autoscalingplans",
    "backup",
    "batch",
    "braket",
    "budgets",
    "chime",
    "cloud9",
    "clouddirectory",
    "cloudformation",
    "cloudfront",
    "cloudhsm",
    "cloudhsmv2",
    "cloudsearch",
    "cloudtrail",
    "cloudwatch",
    "cloudwatchevents",
    "cloudwatchlogs",
    "codeartifact",
    "codebuild",
    "codecommit",
    "codedeploy",
    "codeguruprofiler",
    "codegurureviewer",
    "codepipeline",
    "codestar",
    "codestarconnections",
    "codestarnotifications",
    "cognito",
    "comprehend",
    "comprehendmedical",
    "computeoptimizer",
    "configservice",
    "connect",
    "costandusagereportservice",
    "costexplorer",
    "databasemigrationservice",
    "dataexchange",
    "datapipeline",
    "datasync",
    "dax",
    "detective",
    "devicefarm",
    "directconnect",
    "directoryservice",
    "dlm",
    "docdb",
    "dynamodb",
    "ec2-classic",
    "ec2",
    "ecr",
    "ecrpublic",
    "ecs",
    "efs",
    "eks",
    "elastic-transcoder",
    "elasticache",
    "elasticbeanstalk",
    "elasticinference",
    "elasticsearch",
    "elb",
    "elbv2",
    "emr",
    "emrcontainers",
    "eventbridge",
    "firehose",
    "fms",
    "forecastservice",
    "frauddetector",
    "fsx",
    "gamelift",
    "glacier",
    "globalaccelerator",
    "glue",
    "greengrass",
    "groundstation",
    "guardduty",
    "honeycode",
    "iam",
    "identitystore",
    "imagebuilder",
    "inspector",
    "iot",
    "iotanalytics",
    "iotevents",
    "iotsecuretunneling",
    "iotsitewise",
    "iotthingsgraph",
    "ivs",
    "kafka",
    "kendra",
    "kinesis",
    "kinesisanalytics",
    "kinesisanalyticsv2",
    "kinesisvideo",
    "kms",
    "lakeformation",
    "lambda",
    "lexmodelbuildingservice",
    "licensemanager",
    "lightsail",
    "machinelearning",
    "macie",
    "macie2",
    "managedblockchain",
    "marketplacecatalog",
    "mediaconnect",
    "mediaconvert",
    "medialive",
    "mediapackage",
    "mediapackagevod",
    "mediastore",
    "mediatailor",
    "meteringmarketplace",
    "mobile",
    "mq",
    "mwaa",
    "neptune",
    "networkfirewall",
    "networkmanager",
    "opsworks",
    "opsworkscm",
    "organizations",
    "outposts",
    "personalize",
    "pi",
    "pinpoint",
    "pinpointemail",
    "pinpointsmsvoice",
    "polly",
    "pricing",
    "prometheusservice",
    "qldb",
    "quicksight",
    "ram",
    "rds",
    "redshift",
    "resourcegroups",
    "robomaker",
    "route53",
    "route53domains",
    "route53resolver",
    "s3",
    "s3control",
    "s3outposts",
    "sagemaker",
    "savingsplans",
    "secretsmanager",
    "securityhub",
    "serverlessapplicationrepository",
    "servicecatalog",
    "servicediscovery",
    "servicequotas",
    "ses",
    "sesv2",
    "sfn",
    "shield",
    "signer",
    "simpledb",
    "sms",
    "snowball",
    "sns",
    "sqs",
    "ssm",
    "ssoadmin",
    "storagegateway",
    "sts",
    "support",
    "swf",
    "synthetics",
    "textract",
    "timestreamwrite",
    "transcribeservice",
    "timestreamwrite",
    "transfer",
    "translate",
    "waf",
    "wafv2",
    "workdocs",
    "worklink",
    "workmail",
    "workspaces",
    "xray",
  ]
  description = "Set of AWS Go SDK service labels"
  type        = set(string)
}

resource "github_issue_label" "service" {
  for_each = var.service_labels

  repository = "terraform-provider-aws"
  name       = "service/${each.value}"
  color      = "bfd4f2"
}
