package awslib

import (
  "fmt"
  "reflect"

  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/aws/awserr"
)

// Whoami takes a profile name and returns the account name and the
// canonical account id.
func Whoami(profile string) (string, string) {
  conf := Config{Profile: profile}
  svc := conf.S3Service()

  input := &s3.ListBucketsInput{}
  result, err := svc.ListBuckets(input)
  if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        default:
            fmt.Println(aerr.Error())
         }
    } else {
      // Print the error, cast err to awserr.Error to get the Code and
      // Message from an error.
      fmt.Println(err.Error())
    }
   return "I dont know whats going on", "this is silly"
  }
  displayName := *result.Owner.DisplayName
  fmt.Println("type of displayName:", reflect.TypeOf(displayName))
  accountId := *result.Owner.ID
  fmt.Println("type of accountId:", reflect.TypeOf(accountId))

  return displayName, accountId
  //return displayName
}
