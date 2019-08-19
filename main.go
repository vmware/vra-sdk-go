package main

import (
	"fmt"
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/vmware/vra-sdk-go/pkg/client"
	"github.com/vmware/vra-sdk-go/pkg/client/cloud_account"
	"github.com/vmware/vra-sdk-go/pkg/client/login"
	"github.com/vmware/vra-sdk-go/pkg/client/project"
	"github.com/vmware/vra-sdk-go/pkg/models"

	httptransport "github.com/go-openapi/runtime/client"
)

var apiHost = "api.mgmt.cloud.vmware.com"

func getToken(apiToken string) (string, error) {
	transport := httptransport.New(apiHost, "", nil)
	transport.SetDebug(true)
	fmt.Printf("transport: %+v\n", transport)
	apiclient := client.New(transport, strfmt.Default)

	fmt.Printf("apiclient: %+v\n", apiclient)
	fmt.Printf("transport: %+v\n", apiclient.Transport)
	fmt.Printf("Login: %+v\n", apiclient.Login)
	params := login.NewRetrieveAuthTokenParams().WithBody(
		&models.CspLoginSpecification{
			RefreshToken: &apiToken,
		},
	)
	authTokenResponse, err := apiclient.Login.RetrieveAuthToken(params)
	if err != nil || *authTokenResponse.Payload.TokenType != "bearer" {
		return "", err
	}

	return *authTokenResponse.Payload.Token, nil
}

func main() {
	// Assume a valid refresh token is passed in via an environment variable
	vraRefreshToken := os.Getenv("VRA_REFRESH_TOKEN")
	if vraRefreshToken == "" {
		fmt.Printf("Need to set VRA_REFRESH_TOKEN\n")
		os.Exit(1)
	}

	bearerToken, err := getToken(vraRefreshToken)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
		os.Exit(1)
	}

	transport := httptransport.New(apiHost, "", nil)
	transport.SetDebug(true)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Bearer "+bearerToken)
	apiclient := client.New(transport, strfmt.Default)
	fmt.Printf("apiclient: %+v\n", apiclient)
	ret, err := apiclient.CloudAccount.GetCloudAccounts(cloud_account.NewGetCloudAccountsParams())
	fmt.Printf("%+v %+v\n", ret, err)

	projectReturn, err := apiclient.Project.GetProjects(project.NewGetProjectsParams())
	fmt.Printf("%+v %+v\n", projectReturn, err)
}
