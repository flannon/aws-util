package awslib

import (
	//"fmt"
	"os"
	//"reflect"

	_ "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//type Profile struct {
//	Name string
//}
type Config struct {
	Profile string
	Region string
}

//var Svc *ec2.EC2

// Ec2Service establishes an EC2 session, returning a *ec2.EC2 service.
//func (p Profile) Ec2Service() *ec2.EC2 {
func (c Config) Ec2Service() *ec2.EC2 {

	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
	os.Setenv("AWS_PROFILE", c.Profile)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		//Config: aws.Config{Region: aws.String(s.Region)},
		Profile: c.Profile,
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)

	return svc

}

func (c Config) S3Service() *s3.S3 {
	//sess := session.Must(session.NewSession())
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		}))

	svc := s3.New(sess)

	return svc
}
