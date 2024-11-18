package internal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/olekukonko/tablewriter"
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

func PrintCallerIdentityTable(identity *sts.GetCallerIdentityOutput, name string) {
	// Create a new table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Account", "Role Arn", "User ID"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Append([]string{name, *identity.Account, *identity.Arn, *identity.UserId})
	table.Render()
}

func AssertAccountAsExpected(identity *sts.GetCallerIdentityOutput, expectedAccount string) error {
	if *identity.Account != expectedAccount {
		return fmt.Errorf("Expected account %s, but got %s", expectedAccount, *identity.Account)
	}
	return nil
}
