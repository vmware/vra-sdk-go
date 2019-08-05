package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/vmware/cas-sdk-go/pkg/client"
	"github.com/vmware/cas-sdk-go/pkg/client/cloud_account"
	"github.com/vmware/cas-sdk-go/pkg/client/login"
	"github.com/vmware/cas-sdk-go/pkg/models"

	httptransport "github.com/go-openapi/runtime/client"
)

var apiHost = "api.mgmt.cloud.vmware.com"
var debug = true

func setupCASEnvironment() string {
	// Assume a valid refresh token is passed in via an environment variable
	casRefreshToken := os.Getenv("CAS_REFRESH_TOKEN")
	if casRefreshToken == "" {
		fmt.Printf("Need to set CAS_REFRESH_TOKEN\n")
		os.Exit(1)
	}

	envAPIHost := os.Getenv("CAS_API_HOST")
	if apiHost != "" {
		apiHost = envAPIHost
	}

	if os.Getenv("CAS_DEBUG") != "" {
		debug = true
	}

	return casRefreshToken
}

func setupAWSEnvironment() (string, string) {
	accessKeyID := os.Getenv("CAS_AWS_ACCESS_KEY_ID")
	if accessKeyID == "" {
		fmt.Printf("Need to set CAS_AWS_ACCESS_KEY_ID\n")
		os.Exit(1)
	}

	secretAccessKey := os.Getenv("CAS_AWS_SECRET_ACCESS_KEY")
	if secretAccessKey == "" {
		fmt.Printf("Need to set CAS_AWS_SECRET_ACCESS_KEY\n")
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

	casRefreshToken := setupCASEnvironment()
	accessKeyID, secretAccessKey := setupAWSEnvironment()

	fmt.Printf("Getting bearer token\n")
	bearerToken, err := getToken(casRefreshToken)
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
	createResp, err := apiclient.CloudAccount.CreateAwsCloudAccount(cloud_account.NewCreateAwsCloudAccountParams().WithBody(&models.CloudAccountAwsSpecification{
		Name:            cloudAccountNamePtr,
		AccessKeyID:     withString(accessKeyID),
		SecretAccessKey: withString(secretAccessKey),
		RegionIds:       []string{"us-west-2"},
		/* Tags: []*models.Tag{&models.Tag{
			Key: withString("key"),
			Value: withString("value"),
		},}, */
	}))

	if err != nil {
		if _, ok := err.(*cloud_account.CreateAwsCloudAccountBadRequest); ok {
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
		resp, err := apiclient.CloudAccount.DeleteAwsCloudAccount(cloud_account.NewDeleteAwsCloudAccountParams().WithID(id))
		if err != nil {
			fmt.Printf("Error with delete: %s %+v %+v\n", *cloudAccountNamePtr, resp, err)
		} else {
			fmt.Printf("Deleted cloud account: %s\n", *cloudAccountNamePtr)
		}
	}
}
