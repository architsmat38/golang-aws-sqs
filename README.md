# golang-aws-sqs

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/architsmat38/golang-aws-sqs/master/LICENSE)

AWS SQS Helper - Provides functionality to send, receive, delete, purge messages.
It also helps in processing messages and send them in batches to SQS.

Refer to [example/main.go](https://github.com/architsmat38/golang-aws-sqs/blob/master/example/main.go), in order to get a better understanding of the usage of this helper.

### AWS SQS Config
You will need to specify the following keys in order to connect to an AWS SQS queue.
```
var (
    accessKeyId     string = "xxxxxxxxxxxxxxxxx"
    secretKey       string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    region          string = "xxxxxxxxxxxxxxxxx"
    queueName       string = "xxxxxxxxxxxxxxxxx"
)
```

### Reference
* [AWS SDK](https://github.com/aws/aws-sdk-go) with its [documentation](https://docs.aws.amazon.com/sdk-for-go/api/)
* [Simpleaws](https://github.com/toomore/simpleaws)
* [SQS Poller](https://github.com/h2ik/go-sqs-poller)
