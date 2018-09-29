// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationautoscaling provides the client and types for making API
// requests to Application Auto Scaling.
//
// With Application Auto Scaling, you can configure automatic scaling for your
// scalable resources. You can use Application Auto Scaling to accomplish the
// following tasks:
//
//    * Define scaling policies to automatically scale your AWS or custom resources
//
//    * Scale your resources in response to CloudWatch alarms
//
//    * Schedule one-time or recurring scaling actions
//
//    * View the history of your scaling events
//
// Application Auto Scaling can scale the following resources:
//
//    * Amazon ECS services. For more information, see Service Auto Scaling
//    (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-auto-scaling.html)
//    in the Amazon Elastic Container Service Developer Guide.
//
//    * Amazon EC2 Spot fleets. For more information, see Automatic Scaling
//    for Spot Fleet (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/fleet-auto-scaling.html)
//    in the Amazon EC2 User Guide.
//
//    * Amazon EMR clusters. For more information, see Using Automatic Scaling
//    in Amazon EMR (http://docs.aws.amazon.com/ElasticMapReduce/latest/ManagementGuide/emr-automatic-scaling.html)
//    in the Amazon EMR Management Guide.
//
//    * AppStream 2.0 fleets. For more information, see Fleet Auto Scaling for
//    Amazon AppStream 2.0 (http://docs.aws.amazon.com/appstream2/latest/developerguide/autoscaling.html)
//    in the Amazon AppStream 2.0 Developer Guide.
//
//    * Provisioned read and write capacity for Amazon DynamoDB tables and global
//    secondary indexes. For more information, see Managing Throughput Capacity
//    Automatically with DynamoDB Auto Scaling (http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/AutoScaling.html)
//    in the Amazon DynamoDB Developer Guide.
//
//    * Amazon Aurora Replicas. For more information, see Using Amazon Aurora
//    Auto Scaling with Aurora Replicas (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Integrating.AutoScaling.html).
//
//    * Amazon SageMaker endpoint variants. For more information, see Automatically
//    Scaling Amazon SageMaker Models (http://docs.aws.amazon.com/sagemaker/latest/dg/endpoint-auto-scaling.html).
//
//    * Custom resources provided by your own applications or services. More
//    information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
//
//
// To learn more about Application Auto Scaling, see the Application Auto Scaling
// User Guide (http://docs.aws.amazon.com/autoscaling/application/userguide/what-is-application-auto-scaling.html).
//
// To configure automatic scaling for multiple resources across multiple services,
// use AWS Auto Scaling to create a scaling plan for your application. For more
// information, see the AWS Auto Scaling User Guide (http://docs.aws.amazon.com/autoscaling/plans/userguide/what-is-aws-auto-scaling.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06 for more information on this service.
//
// See applicationautoscaling package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/applicationautoscaling/
//
// Using the Client
//
// To Application Auto Scaling with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Application Auto Scaling client ApplicationAutoScaling for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/applicationautoscaling/#New
package applicationautoscaling
