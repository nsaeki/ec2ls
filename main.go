package main

import (
	"os"
	"fmt"
	flag "github.com/ogier/pflag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mitchellh/go-homedir"
	"strings"
	"reflect"
)

var defaultAttributes = [...]string{
	"InstanceId",
	"Name",
	"InstanceType",
	"AvailabilityZone",
	"State",
	"PrivateIpAddress",
	"PublicIpAddress",
}

func name(i *ec2.Instance) string {
	tags := i.Tags
	for _, tag := range tags {
		if "Name" == *tag.Key {
			return *tag.Value
		}
	}
	return ""
}

func state(i *ec2.Instance) string {
	return *i.State.Name
}

func availabilityZone(i *ec2.Instance) string {
	return *i.Placement.AvailabilityZone
}

var getter = map[string](func(*ec2.Instance)string){
	"Name":  name,
	"State":  state,
	"AvailabilityZone":  availabilityZone,
}

func main() {
	config := aws.Config{}
	profileName := flag.StringP("profile", "p", "", "AWS profile")
	flag.Parse()

	if len(*profileName) > 0 {
		profileFilepath, err := homedir.Expand("~/.aws/credentials")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Using profile:", *profileName)
		creds := credentials.NewSharedCredentials(profileFilepath, *profileName)
		config.Credentials = creds
	}

	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := ec2.New(session.New(), &config)

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	// resp has all of the response data, pull out instance IDs:
	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			attrs := make([]string, 0, 10)
			for _, attr := range defaultAttributes {
				v := ""
				f, ok := getter[attr]
				if ok {
					v = f(inst)
				} else {
					field := reflect.ValueOf(*inst).FieldByName(attr)
					if !field.IsNil() {
						v = fmt.Sprintf("%v", field.Elem().Interface())
					}
				}
				attrs = append(attrs, v)
			}
			fmt.Println(strings.Join(attrs, "\t"))
		}
	}
}

