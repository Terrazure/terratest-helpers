package terratest_helpers //import github.com/Terrazure/terratest-helpers

import (
	"context"
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gruntwork-io/terratest/modules/azure"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/otiai10/copy"
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

var copyOptions = copy.Options{
	Skip: func(src string) (bool, error) {
		return filepath.Ext(src) != ".tf", nil
	},
}

func PrepareTerraformParallelTestingDir(source string, contextName string, testCaseIndex int) string {
	originalTerraformDir, _ := filepath.Abs(fmt.Sprintf("%s", source))
	parallelTerraformDir, _ := filepath.Abs(fmt.Sprintf("%s-%s-%v", source, strings.ToLower(contextName), testCaseIndex))
	copy.Copy(originalTerraformDir, parallelTerraformDir, copyOptions)
	return parallelTerraformDir
}
