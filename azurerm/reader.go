package azurerm

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
)

//go:generate go run ./cmd

// AzureReader is the middleware between TC and AzureRM
type AzureReader struct {
	subscriptionID string
	authorizer     autorest.Authorizer
	env            azure.Environment

	resourceGroup armresources.ResourceGroup
}

// NewAzureReader returns a AzureReader
func NewAzureReader(ctx context.Context, clientID, clientSecret, environment, resourceGroupName, subscriptionID, tenantID string) (*AzureReader, error) {
	env := azure.PublicCloud

	switch environment {
	case "AzureChinaCloud":
		env = azure.ChinaCloud
	case "AzureGermanCloud":
		env = azure.GermanCloud
	case "AzureUSGovernmentCloud":
		env = azure.USGovernmentCloud
	}

	// Config
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		return nil, fmt.Errorf("could not initialize 'azidentity.ClientSecretCredential' because: %s", err)
	}
	oauthConfig, err := adal.NewOAuthConfig(env.ActiveDirectoryEndpoint, tenantID)
	if err != nil {
		return nil, fmt.Errorf("could not initialize 'adal.NewOAuthConfig' because: %s", err)
	}

	token, err := adal.NewServicePrincipalToken(*oauthConfig, clientID, clientSecret, env.ResourceManagerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not initialize 'adal.NewServicePrincipalToken' because: %s", err)
	}

	clientOptions := arm.ClientOptions{}

	// Resource Group
	client, err := armresources.NewResourceGroupsClient(subscriptionID, cred, &clientOptions)
	if err != nil {
		return nil, fmt.Errorf("could not initialize 'armresources.ResourceGroupsClient' because: %s", err)
	}

	resourceGroup, err := client.Get(ctx, resourceGroupName, nil)
	if err != nil {
		return nil, fmt.Errorf("could not 'armresources.ResourceGroupsClient.Get' the resource group because: %s", err)
	}

	return &AzureReader{
		subscriptionID: subscriptionID,
		authorizer:     autorest.NewBearerAuthorizer(token),
		resourceGroup:  resourceGroup.ResourceGroup,
		env:            env,
	}, nil
}

// GetResourceGroup returns the current Resource Group resource
func (ar *AzureReader) GetResourceGroup() armresources.ResourceGroup {
	return ar.resourceGroup
}

// GetResourceGroupName returns the current Resource Group name
func (ar *AzureReader) GetResourceGroupName() string {
	return *ar.resourceGroup.Name
}

// GetLocation returns the current Resource Group location
func (ar *AzureReader) GetLocation() string {
	return *ar.resourceGroup.Location
}
