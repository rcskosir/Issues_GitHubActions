## 4.12.0 (November 28, 2024)

FEATURES:

* **New Data Source**: `azurerm_mssql_managed_database` ([#27026](https://github.com/hashicorp/terraform-provider-azurerm/issues/27026))

BUG FIXES:

* `azurerm_application_insights_api_keyy` - fix condition that nil checks the list of available API keys to prevent an indefinate loop when keys created outside of Terraform are present ([#28037](https://github.com/hashicorp/terraform-provider-azurerm/issues/28037))
* `azurerm_data_factory_linked_service_azure_sql_database` - send `tenant_id` only if it has been specified ([#28120](https://github.com/hashicorp/terraform-provider-azurerm/issues/28120))
* `azurerm_eventgrid_event_subscription` - fix crash when flattening `advanced_filter` ([#28110](https://github.com/hashicorp/terraform-provider-azurerm/issues/28110))
* `azurerm_virtual_network_gateway` - fix crash issue when specifying `root_certificate ` or `revoked_certificate` ([#28099](https://github.com/hashicorp/terraform-provider-azurerm/issues/28099))

ENHANCEMENTS:

* dependencies - update `go-azure-sdk` to `v0.20241128.1112539` ([#28137](https://github.com/hashicorp/terraform-provider-azurerm/issues/28137))
* `containerapps` - update api version to `2024-03-01` ([#28074](https://github.com/hashicorp/terraform-provider-azurerm/issues/28074))
* `Search` - update api version to `2024-06-01-preview` ([#27803](https://github.com/hashicorp/terraform-provider-azurerm/issues/27803))
* Data Source: `azurerm_logic_app_standard` - add support for the `public_network_access` property ([#27913](https://github.com/hashicorp/terraform-provider-azurerm/issues/27913))
* Data Source: `azurerm_search_service` - add support for the `customer_managed_key_encryption_compliance_status` property ([#27478](https://github.com/hashicorp/terraform-provider-azurerm/issues/27478))
* `azurerm_container_registry_task` - add validation on `cpu` as well as on `agent_pool_name`and `agent_setting` ([#28098](https://github.com/hashicorp/terraform-provider-azurerm/issues/28098))
* `azurerm_databricks_workspace` - add support for the `enhanced_security_compliance` block ([#26606](https://github.com/hashicorp/terraform-provider-azurerm/issues/26606))
* `azurerm_eventhub` - deprecate `namespace_name` and `resource_group_name` in favour of `namespace_id` ([#28055](https://github.com/hashicorp/terraform-provider-azurerm/issues/28055))
* `azurerm_logic_app_standard` - add support for the `public_network_access` property ([#27913](https://github.com/hashicorp/terraform-provider-azurerm/issues/27913))
* `azurerm_search_service` - add support for the `customer_managed_key_encryption_compliance_status` property ([#27478](https://github.com/hashicorp/terraform-provider-azurerm/issues/27478))
* `azurerm_cosmosdb_account` - add support for value `EnableNoSQLFullTextSearch` in the  `capabilities.name` property ([#28114](https://github.com/hashicorp/terraform-provider-azurerm/issues/28114))
[BUG] fixing missing quote

date (unreleased) header, Bug, Enhancement, Breaking Change, Features headers
date (unreleased) header, Bug, Enhancement, Breaking Change, Features headers
