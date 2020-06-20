![Test](https://github.com/daviddykeuk/template_go/workflows/Test/badge.svg?branch=master)

# Go Template
This is a Go template repo, designed to work with TDD/BDD and Functional Programming, ready to be built as a container or serverless function.

## Handlers
The repo has multiple different types of handlers, found in `./src/handlers/`

* Command Line or Container
  * API
  * AMQP
  * Kafka
  * MQTT
* AWS Lambda
  * API Gateway
  * DynamoDB Streams
  * Event Bridge
  * Kinesis
  * S3
  * SQS
  * SNS

## Databases
The repo has implementations of object storage using different databases
* CosmosDB
* DynamoDB
* MongoDB

## Event publishers
The repo has implementations of different event publishers
* AWS SQS
* AWS SNS
* AWS Event Bridge
* AWS Kinesis
* AMQP
* Kafka
* MQTT

## Build pipelines
Using Docker and Make as build tools, the build pipelines can be run the same way on multiple build services
* AWS CodeBuild
* Azure DevOps
* Github Actions
* CircleCI
* BuildKite
* Travis CI
