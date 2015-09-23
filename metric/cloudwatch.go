package metric

import (
	"fmt"
	"os"
	"time"

	"github.com/bbc/mozart-api-common/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/bbc/mozart-api-common/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/cloudwatch"
)

// Metric sends metrics to AWS Cloudwatch
func Increment(name string) {

	cosmosEnv := os.Getenv("COSMOS_ENV")

	go func() {
		if cosmosEnv == "live" {
			fmt.Println("live time")
			config := aws.NewConfig().WithRegion(os.Getenv("AWS_REGION"))
			client := cloudwatch.New(config)

			params := &cloudwatch.PutMetricDataInput{
				MetricData: []*cloudwatch.MetricDatum{
					{
						MetricName: aws.String(name),
						Timestamp:  aws.Time(time.Now()),
						Unit:       aws.String("Count"),
						Value:      aws.Float64(1.0),
					},
				},
				Namespace: aws.String(os.Getenv("CLOUDWATCH_NAMESPACE")),
			}

			_, err := client.PutMetricData(params)

			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}()
}
