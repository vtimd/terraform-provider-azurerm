package dataprotection

import "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "DataProtection"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"DataProtection",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azurerm_data_protection_backup_vault":               resourceDataProtectionBackupVault(),
		"azurerm_data_protection_backup_policy_postgresql":   resourceDataProtectionBackupPolicyPostgreSQL(),
		"azurerm_data_protection_backup_instance_postgresql": resourceDataProtectionBackupInstancePostgreSQL(),
	}
}
