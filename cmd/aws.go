package cmd

import (
	"fmt"

	"github.com/arielril/t/internal"
	"github.com/arielril/t/internal/aws"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS Commands",
}

var awsS3UploadCmd = &cobra.Command{
	Use:   "s3-put <file_names>",
	Short: "Upload files to AWS S3",
	Run:   awsS3UploadFiles,
}

func init() {
	awsCmd.PersistentFlags().BoolP("interactive", "i", false, "Interactive mode")
	awsCmd.PersistentFlags().String("region", "us-east-1", "AWS Region")
	awsCmd.PersistentFlags().String("access-key-id", "", "AWS Access Key ID")
	awsCmd.PersistentFlags().String("secret-access-key", "", "AWS Secret Access Key")

	awsS3UploadCmd.Flags().String("bucket", "", "AWS S3 Bucket")
	_ = awsS3UploadCmd.MarkFlagRequired("bucket")

	awsCmd.AddCommand(awsS3UploadCmd)
	rootCmd.AddCommand(awsCmd)
}

func awsS3UploadFiles(cmd *cobra.Command, args []string) {
	interactive, _ := cmd.Flags().GetBool("interactive")
	awsConfig := aws.NewFromCommand(cmd, interactive)

	fileNames := internal.GetCmdPositionalArgs(cmd, args)
	fmt.Printf("uploading files: %v\n", fileNames)

	bucketName, _ := cmd.Flags().GetString("bucket")
	aws.UploadFiles(awsConfig, bucketName, fileNames)
}
