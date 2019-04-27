package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type region struct {
	name, label string
}

var awsregions = [18]region{
	region{name: "us-east-1", label: "US East (N. Virginia)"},
	region{name: "us-east-2", label: "US East (Ohio)"},
	region{name: "us-west-1", label: "US West (N. California)"},
	region{name: "us-west-2", label: "US West (Oregon)"},
	region{name: "ca-central-1", label: "Canada (Central)"},
	region{name: "eu-central-1", label: "EU (Frankfurt)"},
	region{name: "eu-west-1", label: "EU (Ireland)"},
	region{name: "eu-west-2", label: "EU (London)"},
	region{name: "eu-west-3", label: "EU (Paris)"},
	region{name: "eu-north-1", label: "EU (Stockholm)"},
	region{name: "ap-east-1", label: "Asia Pacific (Hong Kong)"},
	region{name: "ap-northeast-1", label: "Asia Pacific (Tokyo)"},
	region{name: "ap-northeast-2", label: "Asia Pacific (Seoul)"},
	region{name: "ap-northeast-3", label: "Asia Pacific (Osaka-Local)"},
	region{name: "ap-southeast-1", label: "Asia Pacific (Singapore)"},
	region{name: "ap-southeast-2", label: "Asia Pacific (Sydney)"},
	region{name: "ap-south-1", label: "Asia Pacific (Mumbai)"},
	region{name: "sa-east-1", label: "South America (São Paulo)"},
}

//HandleRequest Handles API Gateway Request
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	output := ""
	for i := 0; i < len(awsregions); i++ {
		fmt.Println(awsregions[i].name)
		output += awsregions[i].name
	}
	return events.APIGatewayProxyResponse{Body: output, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
