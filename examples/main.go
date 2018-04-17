package main

import (
	"../../awslib"
	"fmt"
	//"os"

	//"github.com/aws/aws-sdk-go/service/s3"
	//"github.com/aws/aws-sdk-go/aws/awserr"

)

func main() {

	fmt.Println("Hello from main, it's cold up here")

	conf := awslib.Config{Profile: "default"}
	conf.Ec2Service()
	//fmt.Println(os.Getenv("AWS_PROFILE"))
	//conf.S3Service()
	displayName, accountId := awslib.Whoami(conf.Profile)
	fmt.Printf("displayName: %v \naccountId: %v\n", displayName, accountId)
}
