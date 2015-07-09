package main

import (
	"os"
	"fmt"
	flag "github.com/ogier/pflag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mitchellh/go-homedir"
)

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

	if len(config.Region) == 0 {
		// Sorry, this is my preference!
		config.Region = "ap-northeast-1"
	}

	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := ec2.New(&config)

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	// resp has all of the response data, pull out instance IDs:
	fmt.Println("> Number of reservation sets: ", len(resp.Reservations))
	for idx, res := range resp.Reservations {
		fmt.Println("  > Number of instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			fmt.Println("    - Instance ID: ", *inst.InstanceID)
		}
	}
}
