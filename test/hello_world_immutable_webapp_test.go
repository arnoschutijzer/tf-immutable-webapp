package test

import (
	"fmt"
	"os"
	"testing"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	awsRegion := "eu-west-1"

	bucketName := "arnoschutijzer-immutable-webapp-test"
	configurationPath := "./configuration"
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"bucket_name":              bucketName,
			"configuration_files_path": configurationPath,
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	}

	terraform.InitAndApply(t, terraformOptions)

	applicationFileName := "./example_app/index.html"
	DeployApplication(t, awsRegion, bucketName, applicationFileName)

	s3uri := terraform.Output(t, terraformOptions, "s3_uri")
	expectedS3Uri := fmt.Sprintf("s3://%s", bucketName)

	assert.Equal(t, s3uri, expectedS3Uri)

	defer CleanUpState(t, awsRegion, bucketName, terraformOptions)
}

func DeployApplication(t *testing.T, awsRegion string, bucketName string, fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Errorf("failed to open file %q, %v", fileName, err).Error())
	}

	s3Uploader := aws.NewS3Uploader(t, awsRegion)
	result, err := s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: awssdk.String(bucketName),
		Key:    awssdk.String(fileName),
		Body:   f,
	})
	if err != nil {
		panic(fmt.Errorf("failed to upload file, %v", err))
	}

	fmt.Printf("file uploaded to, %s\n", awssdk.StringValue(&result.Location))
}

func CleanUpState(t *testing.T, awsRegion string, bucket_name string, terraformOptions *terraform.Options) {
	fmt.Print("cleaning up state")
	aws.EmptyS3Bucket(t, awsRegion, bucket_name)

	terraform.Destroy(t, terraformOptions)
}
