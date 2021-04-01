package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
	"time"
)

func main() {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()
	conf := &aws.Config{
		Region: aws.String("eu-west-1"),
	}
	sess := session.Must(session.NewSession())
	svc := configservice.New(sess, conf)
	params := &configservice.DescribeConfigurationRecordersInput{}
	output, err := svc.DescribeConfigurationRecordersWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
	output2, err := svc.DescribeDeliveryChannelStatus(&configservice.DescribeDeliveryChannelStatusInput{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output2)

	input2 := configservice.PutConfigRuleInput{
		ConfigRule: &configservice.ConfigRule{
			Source: &configservice.Source{
				Owner:            aws.String("AWS"),
				SourceIdentifier: aws.String("IAM_POLICY_IN_USE"),
			},
		},
	}
	output4, err := svc.PutConfigRule(&input2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output4)

}
