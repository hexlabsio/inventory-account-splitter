package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type region struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

var awsRegions = [18]region{
	region{"us-east-1", "US East (N. Virginia)"},
	region{"us-east-2", "US East (Ohio)"},
	region{"us-west-1", "US West (N. California)"},
	region{"us-west-2", "US West (Oregon)"},
	region{"ca-central-1", "Canada (Central)"},
	region{"eu-central-1", "EU (Frankfurt)"},
	region{"eu-west-1", "EU (Ireland)"},
	region{"eu-west-2", "EU (London)"},
	region{"eu-west-3", "EU (Paris)"},
	region{"eu-north-1", "EU (Stockholm)"},
	region{"ap-east-1", "Asia Pacific (Hong Kong)"},
	region{"ap-northeast-1", "Asia Pacific (Tokyo)"},
	region{"ap-northeast-2", "Asia Pacific (Seoul)"},
	region{"ap-northeast-3", "Asia Pacific (Osaka-Local)"},
	region{"ap-southeast-1", "Asia Pacific (Singapore)"},
	region{"ap-southeast-2", "Asia Pacific (Sydney)"},
	region{"ap-south-1", "Asia Pacific (Mumbai)"},
	region{"sa-east-1", "South America (São Paulo)"},
}

//HandleRequest Handles API Gateway Request
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		outputJSON, _ := json.Marshal(awsRegions)
		return events.APIGatewayProxyResponse{Body: string(outputJSON[:]), StatusCode: 200}, nil
	}

	output := ""
	for i := 0; i < len(awsRegions); i++ {
		fmt.Println(awsRegions[i].Name)
		output += awsRegions[i].Name
	}
	return events.APIGatewayProxyResponse{Body: "Accepted", StatusCode: 202}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
