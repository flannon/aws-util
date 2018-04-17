package awslib

import (
  //"fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/ec2"

)

// TagMap is a map of AWS tag key/values
type TagMap map[string]string

// ToEc2Tags converts a TagMap to a slice of ec2.Tag
func ToEc2Tags(m *TagMap) []*ec2.Tag {
	var result []*ec2.Tag
	for k, v := range *m {
		result = append(result, &ec2.Tag{Key: aws.String(k), Value: aws.String(v)})
	}
	return result
}

// FromEc2Tags converts a slice fo ec2.Tag to a TagMap
func FromEc2Tags(tags []*ec2.Tag) TagMap {
  result := make(TagMap)
  result["Name"] = ""
  for _, tag := range tags {
    result[*tag.Key] = *tag.Value
  }
  return result
}
