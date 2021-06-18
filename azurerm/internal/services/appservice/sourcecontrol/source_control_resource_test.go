package sourcecontrol_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/appservice/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type AppServiceSourceControlResource struct{}

func TestAccSourceControlResource_windowsExternalGit(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_app_service_source_control", "test")
	r := AppServiceSourceControlResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.windowsExternalGit(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("scm_type").HasValue("ExternalGit"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSourceControlResource_windowsLocalGit(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_app_service_source_control", "test")
	r := AppServiceSourceControlResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.windowsLocalGit(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("scm_type").HasValue("LocalGit"),
				check.That(data.ResourceName).Key("repo_url").HasValue(fmt.Sprintf("https://acctestwa-%d.scm.azurewebsites.net", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSourceControlResource_windowsGitHubAction(t *testing.T) {
	if ok := os.Getenv("ARM_GITHUB_ACCESS_TOKEN"); ok == "" {
		t.Skip("Skipping as `ARM_GITHUB_ACCESS_TOKEN` is not specified")
	}

	data := acceptance.BuildTestData(t, "azurerm_app_service_source_control", "test")
	r := AppServiceSourceControlResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.windowsGitHubAction(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("scm_type").HasValue("LocalGit"),
				check.That(data.ResourceName).Key("repo_url").HasValue(fmt.Sprintf("https://acctestwa-%d.scm.azurewebsites.net", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSourceControlResource_linuxExternalGit(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_app_service_source_control", "test")
	r := AppServiceSourceControlResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.linuxExternalGit(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("scm_type").HasValue("ExternalGit"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSourceControlResource_linuxLocalGit(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_app_service_source_control", "test")
	r := AppServiceSourceControlResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.linuxLocalGit(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("scm_type").HasValue("LocalGit"),
				check.That(data.ResourceName).Key("repo_url").HasValue(fmt.Sprintf("https://acctestwa-%d.scm.azurewebsites.net", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSourceControlResource_linuxGitHubAction(t *testing.T) {
	if ok := os.Getenv("ARM_GITHUB_ACCESS_TOKEN"); ok == "" {
		t.Skip("Skipping as `ARM_GITHUB_ACCESS_TOKEN` is not specified")
	}

	data := acceptance.BuildTestData(t, "azurerm_app_service_source_control", "test")
	r := AppServiceSourceControlResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.linuxGitHubAction(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r AppServiceSourceControlResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := parse.WebAppID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.AppService.WebAppsClient.GetSourceControl(ctx, id.ResourceGroup, id.SiteName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving Source Control for %s: %v", id, err)
	}
	if resp.SiteSourceControlProperties == nil || resp.SiteSourceControlProperties.RepoURL == nil {
		return utils.Bool(false), nil
	}

	return utils.Bool(true), nil
}

func (r AppServiceSourceControlResource) windowsExternalGit(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_app_service_source_control" "test" {
  app_id   = azurerm_windows_web_app.test.id
  repo_url = "https://github.com/Azure-Samples/app-service-web-dotnet-get-started"
  branch   = "master"
}

`, baseWindowsAppTemplate(data))
}

func (r AppServiceSourceControlResource) linuxExternalGit(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_app_service_source_control" "test" {
  app_id   = azurerm_linux_web_app.test.id
  repo_url = "https://github.com/Azure-Samples/python-docs-hello-world"
  branch   = "master"
}

`, baseLinuxAppTemplate(data))
}

func (r AppServiceSourceControlResource) windowsLocalGit(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_app_service_source_control" "test" {
  app_id   = azurerm_windows_web_app.test.id
  acm_type = "LocalGit"
}

`, baseWindowsAppTemplate(data))
}

func (r AppServiceSourceControlResource) linuxLocalGit(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_app_service_source_control" "test" {
  app_id   = azurerm_linux_web_app.test.id
  scm_type = "LocalGit"
}

`, baseLinuxAppTemplate(data))
}

// TODO - Figure this out, GHA items here are not well documented
func (r AppServiceSourceControlResource) windowsGitHubAction(data acceptance.TestData) string {
	token := os.Getenv("ARM_GITHUB_ACCESS_TOKEN")
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource azurerm_app_service_github_token test {
  token = "%s"
}

resource "azurerm_app_service_source_control" "test" {
  app_id   = azurerm_windows_web_app.test.id
  repo_url = "https://github.com/jackofallops/azure-app-service-static-site-tests.git"
  scm_type = "GitHub"

  github_action_configuration {
    linux_action           = true
    generate_workflow_file = true

    container_configuration {
      registry_url = ""
      image_name   = ""
    }

    code_configuration {
      runtime_stack   = ""
      runtime_version = ""
    }
  }

}

`, baseWindowsAppTemplate(data), token)
}

// TODO - Figure this out, GHA items here are not well documented
func (r AppServiceSourceControlResource) linuxGitHubAction(data acceptance.TestData) string {
	token := os.Getenv("ARM_GITHUB_ACCESS_TOKEN")
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource azurerm_app_service_github_token test {
  token = "%s"
}

resource "azurerm_app_service_source_control" "test" {
  app_id   = azurerm_linux_web_app.test.id
  repo_url = "https://github.com/jackofallops/azure-app-service-static-site-tests.git"
  scm_type = "GitHub"

  github_action_configuration {
    linux_action           = true
    generate_workflow_file = true

    container_configuration {
      registry_url = ""
      image_name   = ""
    }

    code_configuration {
      runtime_stack   = ""
      runtime_version = ""
    }
  }
}

`, baseLinuxAppTemplate(data), token)
}

func baseWindowsAppTemplate(data acceptance.TestData) string {
	return fmt.Sprintf(`

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-ASSC-%[1]d"
  location = "%s"
}

resource "azurerm_app_service_plan" "test" {
  name                = "acctestASSC-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  kind                = "Windows"

  sku {
    tier = "Standard"
    size = "S1"
  }
}

resource "azurerm_windows_web_app" "test" {
  name                = "acctestWA-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_app_service_plan.test.id
}

`, data.RandomInteger, data.Locations.Primary)
}

func baseLinuxAppTemplate(data acceptance.TestData) string {
	return fmt.Sprintf(`

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-ASSC-%[1]d"
  location = "%[2]s"
}

resource "azurerm_app_service_plan" "test" {
  name                = "acctestASP-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  kind                = "Linux"
  reserved            = true

  sku {
    tier = "Standard"
    size = "S1"
  }
}

resource "azurerm_linux_web_app" "test" {
  name                = "acctestWA-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_app_service_plan.test.id

  site_config {
    application_stack {
      python_version = "3.8"
    }
  }
}
`, data.RandomInteger, data.Locations.Primary)
}
