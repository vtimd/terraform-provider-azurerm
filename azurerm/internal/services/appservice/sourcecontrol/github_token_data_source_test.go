package sourcecontrol_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
)

type AppServiceGithubTokenDataSource struct{}

func TestAccSourceControlGitHubTokenDataSource_basic(t *testing.T) {
	if ok := os.Getenv("ARM_GITHUB_ACCESS_TOKEN"); ok == "" {
		t.Skip("Skipping as `ARM_GITHUB_ACCESS_TOKEN` is not specified")
	}

	data := acceptance.BuildTestData(t, "data.azurerm_app_service_github_token", "test")
	r := AppServiceGithubTokenDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.basic(),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("token").Exists(),
			),
		},
	})
}

func (AppServiceGithubTokenDataSource) basic() string {
	return fmt.Sprintf(`

%s

data azurerm_app_service_github_token test {}

`, AppServiceGitHubTokenResource{}.basic())
}
