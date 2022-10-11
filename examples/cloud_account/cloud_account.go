package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/vmware/vra-sdk-go/pkg/client"
	"github.com/vmware/vra-sdk-go/pkg/client/cloud_account"
	"github.com/vmware/vra-sdk-go/pkg/client/login"
	"github.com/vmware/vra-sdk-go/pkg/models"

	httptransport "github.com/go-openapi/runtime/client"
)

var apiHost = "api.mgmt.cloud.vmware.com"
var debug = true

func setupVRAEnvironment() string {
	// Assume a valid refresh token is passed in via an environment variable
	vraRefreshToken := os.Getenv("VRA_REFRESH_TOKEN")
	if vraRefreshToken == "" {
		fmt.Printf("Need to set VRA_REFRESH_TOKEN\n")
		os.Exit(1)
	}

	envAPIHost := os.Getenv("VRA_API_HOST")
	if apiHost != "" {
		apiHost = envAPIHost
	}

	if os.Getenv("VRA_DEBUG") != "" {
		debug = true
	}

	return vraRefreshToken
}

func setupAWSEnvironment() (string, string) {
	accessKeyID := os.Getenv("VRA_AWS_ACCESS_KEY_ID")
	if accessKeyID == "" {
		fmt.Printf("Need to set VRA_AWS_ACCESS_KEY_ID\n")
		os.Exit(1)
	}

	secretAccessKey := os.Getenv("VRA_AWS_SECRET_ACCESS_KEY")
	if secretAccessKey == "" {
		fmt.Printf("Need to set VRA_AWS_SECRET_ACCESS_KEY\n")
		os.Exit(1)
	}
	return accessKeyID, secretAccessKey
}

func getToken(apiToken string) (string, error) {
	transport := httptransport.New(apiHost, "", nil)
	transport.SetDebug(debug)
	apiclient := client.New(transport, strfmt.Default)

	params := login.NewRetrieveAuthTokenParams().WithBody(
		&models.CspLoginSpecification{
			RefreshToken: &apiToken,
		},
	)
	authTokenResponse, err := apiclient.Login.RetrieveAuthToken(params)
	if err != nil || *authTokenResponse.Payload.TokenType != "Bearer" {
		return "", err
	}

	return *authTokenResponse.Payload.Token, nil
}

func withString(s string) *string {
	return &s
}

func main() {
	flag.BoolVar(&debug, "debug", false, "output verbose http debug information")
	deletePtr := flag.Bool("delete", false, "delete cloud account")
	cloudAccountNamePtr := flag.String("name", "", "Name of test cloud account")
	flag.Parse()

	if *cloudAccountNamePtr == "" {
		fmt.Printf("Must specify a name for the new cloud account\n")
		os.Exit(1)
	}

	vraRefreshToken := setupVRAEnvironment()
	accessKeyID, secretAccessKey := setupAWSEnvironment()

	fmt.Printf("Getting bearer token\n")
	bearerToken, err := getToken(vraRefreshToken)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
		os.Exit(1)
	}
	// bearerAuth := httptransport.BearerToken(bearerToken)

	transport := httptransport.New(apiHost, "", nil)
	transport.SetDebug(debug)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Bearer "+bearerToken)
	apiclient := client.New(transport, strfmt.Default)

	fmt.Printf("Creating cloud account %s\n", *cloudAccountNamePtr)
	createResp, err := apiclient.CloudAccount.CreateAwsCloudAccountAsync(cloud_account.NewCreateAwsCloudAccountAsyncParams().WithBody(&models.CloudAccountAwsSpecification{
		Name:            cloudAccountNamePtr,
		AccessKeyID:     withString(accessKeyID),
		SecretAccessKey: withString(secretAccessKey),
		Regions: []*models.RegionSpecification{
			&models.RegionSpecification{
				ExternalRegionID: withString("us-west-2"),
				Name:             withString("us-west-2"),
			},
		},
		/* Tags: []*models.Tag{&models.Tag{
			Key: withString("key"),
			Value: withString("value"),
		},}, */
	}))

	if err != nil {
		if _, ok := err.(*cloud_account.CreateAwsCloudAccountAsyncBadRequest); ok {
			fmt.Printf("Cloud account '%s' already exists\n", *cloudAccountNamePtr)
		} else {
			fmt.Printf("Unknown error: %+v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Created cloud account: %s, id: %s\n", createResp.Payload.Name, *createResp.Payload.ID)
	}

	id := ""
	ret, err := apiclient.CloudAccount.GetCloudAccounts(cloud_account.NewGetCloudAccountsParams())
	for i := int64(0); i < ret.Payload.TotalElements; i++ {
		cloudAccount := ret.Payload.Content[i]
		fmt.Printf("Element %d: name: %s, type: %s\n", i, cloudAccount.Name, *cloudAccount.CloudAccountType)
		if cloudAccount.Name == *cloudAccountNamePtr {
			id = *cloudAccount.ID
		}
	}

	if *deletePtr && id != "" {
		_, _, err := apiclient.CloudAccount.DeleteAwsCloudAccount(cloud_account.NewDeleteAwsCloudAccountParams().WithID(id))
		if err != nil {
			fmt.Printf("Error with delete: %s %+v\n", *cloudAccountNamePtr, err)
		} else {
			fmt.Printf("Deleted cloud account: %s\n", *cloudAccountNamePtr)
		}
	}
}
