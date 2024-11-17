package internal

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
	"time"
)

package main

import (
"context"
"fmt"
"log"

"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/config"
"github.com/aws/aws-sdk-go-v2/service/sts"
)

type StsClient struct {
	Client *sts.Client
}

func NewStsClient() (*StsClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &StsClient{
		Client: sts.NewFromConfig(cfg),
	}, nil
}

func (s *StsClient) GetCallerIdentity() (*sts.GetCallerIdentityOutput, error) {
	input := &sts.GetCallerIdentityInput{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Client.GetCallerIdentity(ctx, input)
}

func StsGetCallerIdentity() {
	// Load the shared AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an STS client
	client := sts.NewFromConfig(cfg)

	// Call the GetCallerIdentity API
	result, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("unable to get caller identity, %v", err)
	}

	// Print the caller identity
	fmt.Printf("Account: %s\n", *result.Account)
	fmt.Printf("ARN: %s\n", *result.Arn)
	fmt.Printf("UserId: %s\n", *result.UserId)
}