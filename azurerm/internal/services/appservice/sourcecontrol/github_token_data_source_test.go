package sourcecontrol_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
)

type AppServiceGithubTokenDataSource struct{}

func TestAccSourceControlGitHubTokenDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_app_service_github_token", "test")
	r := AppServiceGithubTokenDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.basic(),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("token").IsSet(),
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
