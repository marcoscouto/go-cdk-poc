# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests


## Run Local

To run this project locally you have to install (docker)[https://www.docker.com/] and (cdklocal)[https://github.com/localstack/aws-cdk-local].

To start local stack execute:

```
docker-compose up
```

To deploy this stack locally execute:

```
cdklocal deploy
```

## Stack Components

- 2 SQS Queues
- 2 SQS DLQ Queues
- 2 SNS Topics
- 2 Subscription SQS to SNS
- 2 Buckets S3