package awslib

import (
  "fmt"
  "log"
  "github.com/aws/aws-sdk-go/aws"
  //"github.com/aws/aws-sdk-go/aws/awserr"
  ////"github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
  //"github.com/mitchellh/mapstructure"
)

//var Svc *ec2.EC2

// Get list of tagged instances
// !+getTaggedInstances()
func getTaggedInstances(svc *ec2.EC2, t string) []*ec2.Instance {

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:" + t),
				Values: []*string{
					aws.String("daily"), aws.String("/dev/sda"), aws.String("monthly")},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("error listing instances in", err.Error())
		log.Fatal(err.Error())
	}

	var taggedInstances []*ec2.Instance

	// Reservations can have one or more taggedInstances,
	// so we need to loop through it twice, first getting a Reservation
	// then getting instances for the reservation.
	for r := range resp.Reservations {
		for _, inst := range resp.Reservations[r].Instances {
			if inst != nil {
				taggedInstances = append(taggedInstances, inst)
			}
		}
	}
	return taggedInstances
} // !-getTaggedInstances()
