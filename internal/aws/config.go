package aws

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/spf13/cobra"
)

type Config struct {
	KeyID     string
	SecretKey string
	Region    string
}

func NewFromCommand(cmd *cobra.Command, interactive bool) aws.Config {
	fmt.Printf("is interactive? %v\n", interactive)
	accessKeyId, _ := cmd.PersistentFlags().GetString("access-key-id")
	secretAccessKey, _ := cmd.PersistentFlags().GetString("secret-access-key")
	region, _ := cmd.PersistentFlags().GetString("region")

	customConfig := &Config{
		KeyID:     accessKeyId,
		SecretKey: secretAccessKey,
		Region:    region,
	}

	var cfg aws.Config
	var err error

	if interactive {
		//if customConfig.KeyID == "" {
		// get aws key id from user
		customConfig.KeyID = strings.TrimSpace(readUserInput("AWS Access Key ID: "))
		//}

		//if customConfig.SecretKey == "" {
		// get aws secret key id from user
		customConfig.SecretKey = strings.TrimSpace(readUserInput("AWS Secret Key: "))
		//}

		//if customConfig.Region == "" {
		customConfig.Region = strings.TrimSpace(readUserInput("AWS Region (default `us-east-1`): "))

		if customConfig.Region == "" {
			customConfig.Region = "us-east-1"
		}
		//}

		// https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/configure-gosdk.html
		cfg, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(customConfig.KeyID, customConfig.SecretKey, ""),
			),
			config.WithRegion(customConfig.Region),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO())
	}
	cobra.CheckErr(err)

	return cfg
}

func readUserInput(txt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(txt)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading user input: %s\n", err)
		os.Exit(1)
	}

	return strings.TrimSpace(response)
}
