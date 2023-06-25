package terratest_helpers

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gruntwork-io/terratest/modules/azure"
	"testing"
	"time"
)

func GetAuthorizer(t *testing.T) autorest.Authorizer {
	auth, err := azure.NewAuthorizer()
	if err != nil {
		t.Fatal("Error when trying to get authorization token")
	}
	return *auth
}

func ConfigureAzureResourceClient(t *testing.T, client *autorest.Client) {
	authorizer := GetAuthorizer(t)
	client.Authorizer = authorizer
	err := client.AddToUserAgent("testing-agent")
	if err != nil {
		t.Fatalf("Failed to add user agent to HTTP client")
	}
}

func BuildDefaultHttpContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Minute)
}
