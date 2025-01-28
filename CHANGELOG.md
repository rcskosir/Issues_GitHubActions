## 4.14.0 (Unreleased)

ENHANCEMENTS:
* `breaking` - remove `updates the template and the accepted entry type` [GH-27950]

FEATURES:

BUG FIXES:

## 4.13.0 (December 05, 2024)
ENHANCEMENTS:
* `azurerm_cognitive_deployment` - support for the `dynamic_throttling_enabled` property ([#28100](https://github.com/hashicorp/terraform-provider-azurerm/issues/28100))
* `azurerm_key_vault_managed_hardware_security_module_key` - the `key_type` property now supports `oct-HSM` ([#28171](https://github.com/hashicorp/terraform-provider-azurerm/issues/28171))
* `azurerm_machine_learning_datastore_datalake_gen2` - can now be used with storage account in a different subscription ([#28123](https://github.com/hashicorp/terraform-provider-azurerm/issues/28123))
* `azurerm_network_watcher_flow_log` - `target_resource_id` supports subnets and network interfaces ([#28177](https://github.com/hashicorp/terraform-provider-azurerm/issues/28177))
BUG:
* Data Source: `azurerm_logic_app_standard` - update the `identity` property to support User Assigned Identities ([#28158](https://github.com/hashicorp/terraform-provider-azurerm/issues/28158))
* `azurerm_cdn_frontdoor_origin_group` - update validation of the `interval_in_seconds` property to match API behaviour ([#28143](https://github.com/hashicorp/terraform-provider-azurerm/issues/28143))
* `azurerm_container_group` - retrieve log analytics workspace key from config when updating resource ([#28025](https://github.com/hashicorp/terraform-provider-azurerm/issues/28025))
* `azurerm_mssql_elasticpool` - fix sku tier and family validation that prevented the creation of Hyperscale PRMS pools ([#28178](https://github.com/hashicorp/terraform-provider-azurerm/issues/28178))
* `azurerm_search_service` -  the `partition_count` property can now be up to `3` when using basic sku ([#28105](https://github.com/hashicorp/terraform-provider-azurerm/issues/28105))


## 4.12.0 (November 28, 2024)

FEATURES:

* **New Data Source**: `azurerm_mssql_managed_database` ([#27026](https://github.com/hashicorp/terraform-provider-azurerm/issues/27026))

BUG FIXES:

* `azurerm_application_insights_api_key` - fix condition that nil checks the list of available API keys to prevent an indefinate loop when keys created outside of Terraform are present ([#28037](https://github.com/hashicorp/terraform-provider-azurerm/issues/28037))
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
