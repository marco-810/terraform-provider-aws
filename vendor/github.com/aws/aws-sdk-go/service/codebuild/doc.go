// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codebuild provides the client and types for making API
// requests to AWS CodeBuild.
//
// AWS CodeBuild is a fully managed build service in the cloud. AWS CodeBuild
// compiles your source code, runs unit tests, and produces artifacts that are
// ready to deploy. AWS CodeBuild eliminates the need to provision, manage,
// and scale your own build servers. It provides prepackaged build environments
// for the most popular programming languages and build tools, such as Apache
// Maven, Gradle, and more. You can also fully customize build environments
// in AWS CodeBuild to use your own build tools. AWS CodeBuild scales automatically
// to meet peak build requests. You pay only for the build time you consume.
// For more information about AWS CodeBuild, see the AWS CodeBuild User Guide
// (https://docs.aws.amazon.com/codebuild/latest/userguide/welcome.html).
//
// AWS CodeBuild supports these operations:
//
//    * BatchDeleteBuilds: Deletes one or more builds.
//
//    * BatchGetBuilds: Gets information about one or more builds.
//
//    * BatchGetProjects: Gets information about one or more build projects.
//    A build project defines how AWS CodeBuild runs a build. This includes
//    information such as where to get the source code to build, the build environment
//    to use, the build commands to run, and where to store the build output.
//    A build environment is a representation of operating system, programming
//    language runtime, and tools that AWS CodeBuild uses to run a build. You
//    can add tags to build projects to help manage your resources and costs.
//
//    * CreateProject: Creates a build project.
//
//    * CreateReportGroup: Creates a report group. A report group contains a
//    collection of reports.
//
//    * CreateWebhook: For an existing AWS CodeBuild build project that has
//    its source code stored in a GitHub or Bitbucket repository, enables AWS
//    CodeBuild to start rebuilding the source code every time a code change
//    is pushed to the repository.
//
//    * DeleteProject: Deletes a build project.
//
//    * DeleteReport: Deletes a report.
//
//    * DeleteReportGroup: Deletes a report group.
//
//    * DeleteResourcePolicy: Deletes a resource policy that is identified by
//    its resource ARN.
//
//    * DeleteSourceCredentials: Deletes a set of GitHub, GitHub Enterprise,
//    or Bitbucket source credentials.
//
//    * DeleteWebhook: For an existing AWS CodeBuild build project that has
//    its source code stored in a GitHub or Bitbucket repository, stops AWS
//    CodeBuild from rebuilding the source code every time a code change is
//    pushed to the repository.
//
//    * DescribeTestCases: Returns a list of details about test cases for a
//    report.
//
//    * GetResourcePolicy: Gets a resource policy that is identified by its
//    resource ARN.
//
//    * ImportSourceCredentials: Imports the source repository credentials for
//    an AWS CodeBuild project that has its source code stored in a GitHub,
//    GitHub Enterprise, or Bitbucket repository.
//
//    * InvalidateProjectCache: Resets the cache for a project.
//
//    * ListBuilds: Gets a list of build IDs, with each build ID representing
//    a single build.
//
//    * ListBuildsForProject: Gets a list of build IDs for the specified build
//    project, with each build ID representing a single build.
//
//    * ListCuratedEnvironmentImages: Gets information about Docker images that
//    are managed by AWS CodeBuild.
//
//    * ListProjects: Gets a list of build project names, with each build project
//    name representing a single build project.
//
//    * ListReportGroups: Gets a list ARNs for the report groups in the current
//    AWS account.
//
//    * ListReports: Gets a list ARNs for the reports in the current AWS account.
//
//    * ListReportsForReportGroup: Returns a list of ARNs for the reports that
//    belong to a ReportGroup.
//
//    * ListSharedProjects: Gets a list of ARNs associated with projects shared
//    with the current AWS account or user.
//
//    * ListSharedReportGroups: Gets a list of ARNs associated with report groups
//    shared with the current AWS account or user
//
//    * ListSourceCredentials: Returns a list of SourceCredentialsInfo objects.
//    Each SourceCredentialsInfo object includes the authentication type, token
//    ARN, and type of source provider for one set of credentials.
//
//    * PutResourcePolicy: Stores a resource policy for the ARN of a Project
//    or ReportGroup object.
//
//    * StartBuild: Starts running a build.
//
//    * StopBuild: Attempts to stop running a build.
//
//    * UpdateProject: Changes the settings of an existing build project.
//
//    * UpdateReportGroup: Changes a report group.
//
//    * UpdateWebhook: Changes the settings of an existing webhook.
//
// See https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06 for more information on this service.
//
// See codebuild package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codebuild/
//
// Using the Client
//
// To contact AWS CodeBuild with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CodeBuild client CodeBuild for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codebuild/#New
package codebuild
