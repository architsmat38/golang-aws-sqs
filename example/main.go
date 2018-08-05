package example

import(
    "log"
    "time"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/sqs"
    SqsService "github.com/architsmat38/golang-aws-sqs/sqs"
    "github.com/architsmat38/golang-aws-sqs/poller"
)

var (
    accessKeyId     string = "xxxxxxxxxxxxxxxxx"
    secretKey       string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    region          string = "xxxxxxxxxxxxxxxxx"
    queueName       string = "xxxxxxxxxxxxxxxxx"
)

func InitializePollerSQS() {
    go poller.Start(poller.HandlerFunc(func(msg *sqs.Message) error {
        var queueMessage string = aws.StringValue(msg.Body)
        decoded, err := SqsService.Decode([]byte(queueMessage))
        if err != nil {
            return err
        }

        log.Println(string(decoded))
        return nil
    }))
}

func main() {
    // Intiialize SQS client
    SqsService.Initialize(
        SqsService.New(queueName, region, accessKeyId, secretKey, ""),
        SqsService.SetWaitSeconds(20),
    )

    // Initialize poller
    InitializePollerSQS()

    // Send
    SqsService.Send(`{"id":1,"message":"Sending data"}`)

    // Send in batches
    var data []string = []string{`{"id":1,"message":"First message"}`, `{"id":2,"message":"Second message"}`}
    SqsService.ProcessAndSendBatch(data)

    time.Sleep(1*time.Minute)
}
