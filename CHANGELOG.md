## 4.17.0 (Unreleased)

ENHANCEMENTS:

FEATURES:

BUG FIXES:

## 4.16.0 (January 16, 2025)

**NOTE:** This release contains a breaking change reverting `redisenterprise` API version from `2024-10-01` to `2024-06-01-preview` as not all regions are currently supported in the `2024-10-01` version 

BREAKING CHANGES:

* dependencies - `redisenterprise` API version reverted from `2024-10-01` to `2024-06-01-preview` ([#28516](https://github.com/hashicorp/terraform-provider-azurerm/issues/28516))

FEATURES:

* **New Resource**: `azurerm_container_registry_credential_set` ([#27528](https://github.com/hashicorp/terraform-provider-azurerm/issues/27528))
* **New Resource**: `azurerm_mssql_job` ([#28456](https://github.com/hashicorp/terraform-provider-azurerm/issues/28456))
* **New Resource**: `azurerm_mssql_job_schedule` ([#28456](https://github.com/hashicorp/terraform-provider-azurerm/issues/28456))

ENHANCEMENTS:

* dependencies - update `hashicorp/go-azure-sdk` to `v0.20250115.1141151` ([#28519](https://github.com/hashicorp/terraform-provider-azurerm/issues/28519))
* dependencies - `costmanagement` update to use `2023-08-01` ([#27680](https://github.com/hashicorp/terraform-provider-azurerm/issues/27680))
* dependencies - `postgresql` update API version to `2024-08-01` ([#28474](https://github.com/hashicorp/terraform-provider-azurerm/issues/28474))
* `azurerm_container_app` – support for the `termination_grace_period_seconds` property ([#28307](https://github.com/hashicorp/terraform-provider-azurerm/issues/28307))
* `azurerm_cost_anomaly_alert` - add support for the `notification_email` property ([#27680](https://github.com/hashicorp/terraform-provider-azurerm/issues/27680))
* `azurerm_data_protection_backup_vault` - support for `immutability` property ([#27859](https://github.com/hashicorp/terraform-provider-azurerm/issues/27859))
* `azurerm_databricks_workspace` - fix `ignore_changes` support ([#28527](https://github.com/hashicorp/terraform-provider-azurerm/issues/28527))
* `azurerm_kubernetes_cluster_node_pool` - add support for the `temporary_name_for_rotation` property to allow node pool rotation ([#27791](https://github.com/hashicorp/terraform-provider-azurerm/issues/27791))
* `azurerm_linux_function_app` - add  support for node `22` and java `17` support for `JBOSSEAP` ([#28472](https://github.com/hashicorp/terraform-provider-azurerm/issues/28472))
* `azurerm_linux_web_app` - add  support for node `22` and java `17` support for `JBOSSEAP` ([#28472](https://github.com/hashicorp/terraform-provider-azurerm/issues/28472))
* `azurerm_windows_function_app` - add  support for node `22` and java `17` support for `JBOSSEAP` ([#28472](https://github.com/hashicorp/terraform-provider-azurerm/issues/28472))


BUG FIXES:

* `azurerm_logic_app_standard` - fix setting `public_network_access` for conflicting API properties ([#28465](https://github.com/hashicorp/terraform-provider-azurerm/issues/28465))
* `azurerm_redis_cache` - `data_persistence_authentication_method` can now be unset ([#27932](https://github.com/hashicorp/terraform-provider-azurerm/issues/27932))
* `azurerm_mssql_database` - fix bug where verifying TDE might fail to return an error on failure ([#28505](https://github.com/hashicorp/terraform-provider-azurerm/issues/28505))
* `azurerm_mssql_database` - fix several potential bugs where retry functions could return false negatives for actual errors ([#28505](https://github.com/hashicorp/terraform-provider-azurerm/issues/28505))
* `azurerm_private_endpoint` - fix a bug where reading Private DNS could error and exit the Read of the resource early without raising an error ([#28505](https://github.com/hashicorp/terraform-provider-azurerm/issues/28505))


## 4.15.0 (January 10, 2025)

FEATURES:

* **New Data Source**: `azurerm_kubernetes_fleet_manager` ([#28278](https://github.com/hashicorp/terraform-provider-azurerm/issues/28278))
* **New Resource**: `azurerm_arc_kubernetes_provisioned_cluster` ([#28216](https://github.com/hashicorp/terraform-provider-azurerm/issues/28216))
* **New Resource**: `azurerm_machine_learning_workspace_network_outbound_rule_private_endpoint` ([#27874](https://github.com/hashicorp/terraform-provider-azurerm/issues/27874))
* **New Resource** `azurerm_machine_learning_workspace_network_outbound_rule_service_tag` ([#27931](https://github.com/hashicorp/terraform-provider-azurerm/issues/27931))
* **New Resource** `azurerm_dynatrace_tag_rules` ([#27985](https://github.com/hashicorp/terraform-provider-azurerm/issues/27985))

ENHANCEMENTS:

* dependencies - update tool Go version and bump `go-git` version to `5.13.0` ([#28425](https://github.com/hashicorp/terraform-provider-azurerm/issues/28425))
* dependencies - update `hashicorp/go-azure-sdk` to `v0.20241212.1154051` ([#28360](https://github.com/hashicorp/terraform-provider-azurerm/issues/28360))
* dependencies - `frontdoor` - partial update to use `2024-02-01` API ([#28233](https://github.com/hashicorp/terraform-provider-azurerm/issues/28233))
* dependencies - `postgresql` - update to `2024-08-01` ([#28380](https://github.com/hashicorp/terraform-provider-azurerm/issues/28380))
* dependencies - `redisenterprise` - update to `2024-10-01` and support for new skus ([#28280](https://github.com/hashicorp/terraform-provider-azurerm/issues/28280))
* Data Source: `azurerm_healthcare_dicom_service` - add support for the `data_partitions_enabled`, `cors`, `encryption_key_url` and `storage` properties ([#27375](https://github.com/hashicorp/terraform-provider-azurerm/issues/27375))
* Data Source: `azurerm_nginx_deployment` - add support for the `dataplane_api_endpoint` property ([#28379](https://github.com/hashicorp/terraform-provider-azurerm/issues/28379)) 
* Data Source: `azurerm_static_web_app` - add  support for the `repository_url` and `repository_branch` properties ([#27401](https://github.com/hashicorp/terraform-provider-azurerm/issues/27401))
* `azurerm_billing_account_cost_management_export` - add support for the `file_format` property ([#27122](https://github.com/hashicorp/terraform-provider-azurerm/issues/27122))
* `azurerm_cdn_frontdoor_profile` - add support for the `identity` property ([#28281](https://github.com/hashicorp/terraform-provider-azurerm/issues/28281))
* `azurerm_cognitive_deployment` - `DataZoneProvisionedManaged` and `GlobalProvisionedManaged` skus are now supported ([#28404](https://github.com/hashicorp/terraform-provider-azurerm/issues/28404))
* `azurerm_databricks_access_connector` - `SystemAssigned,UserAssigned` identity is now supported ([#28442](https://github.com/hashicorp/terraform-provider-azurerm/issues/28442))
* `azurerm_healthcare_dicom_service` - add support for the `data_partitions_enabled`, `cors`, `encryption_key_url` and `storage` properties ([#27375](https://github.com/hashicorp/terraform-provider-azurerm/issues/27375))
* `azurerm_kubernetes_flux_configuration` - add support for the `post_build` and `wait` properties ([#25695](https://github.com/hashicorp/terraform-provider-azurerm/issues/25695))
* `azurerm_linux_virtual_machine` - export the `os_disk.0.id` attribute ([#28352](https://github.com/hashicorp/terraform-provider-azurerm/issues/28352))
* `azurerm_netapp_volume` - make the `network_features` property Optional/Computed ([#28390](https://github.com/hashicorp/terraform-provider-azurerm/issues/28390))
* `azurerm_nginx_deployment` - add support for the `dataplane_api_endpoint` property ([#28379](https://github.com/hashicorp/terraform-provider-azurerm/issues/28379)) 
* `azurerm_resource_group_cost_management_export` - add support for the `file_format` property ([#27122](https://github.com/hashicorp/terraform-provider-azurerm/issues/27122))
* `azurerm_site_recovery_replicated_vm` - support for the `network_interface.recovery_load_balancer_backend_address_pool_ids` property ([#28398](https://github.com/hashicorp/terraform-provider-azurerm/issues/28398))
* `azurerm_static_web_app` - add  support for the `repository_url`, `repository_branch` and `repository_token` properties ([#27401](https://github.com/hashicorp/terraform-provider-azurerm/issues/27401))
* `azurerm_subscription_cost_management_export` - add support for the `file_format` property ([#27122](https://github.com/hashicorp/terraform-provider-azurerm/issues/27122))
* `azurerm_virtual_network` - support for the `private_endpoint_vnet_policies` property ([#27830](https://github.com/hashicorp/terraform-provider-azurerm/issues/27830))
* `azurerm_windows_virtual_machine` - export the `os_disk.0.id` attribute ([#28352](https://github.com/hashicorp/terraform-provider-azurerm/issues/28352))
* `azurerm_mssql_managed_instance` - support for new property `azure_active_directory_administrator` ([#24801](https://github.com/hashicorp/terraform-provider-azurerm/issues/24801))

BUG FIXES:

* `azurerm_api_management` - update the `capacity` property to allow increasing the apim scalability to `31` ([#28427](https://github.com/hashicorp/terraform-provider-azurerm/issues/28427))
* `azurerm_automation_software_update_configuration` remove deprecated misspelled attribute `error_meesage` ([#28312](https://github.com/hashicorp/terraform-provider-azurerm/issues/28312))
* `azurerm_batch_pool` - support for new block `security_profile` ([#28069](https://github.com/hashicorp/terraform-provider-azurerm/issues/28069))
* `azurerm_log_analytics_data_export_rule` - now creates successfully without returning `404` ([#27876](https://github.com/hashicorp/terraform-provider-azurerm/issues/27876))
* `azurerm_mongo_cluster` - remove CustomizeDiff logic for `administrator_password` to allow the input to be generated by the `random_password` resource ([#28215](https://github.com/hashicorp/terraform-provider-azurerm/issues/28215))
* `azurerm_mongo_cluster` - valdation updated so the resource now creates successfully when using `create_mode` `GeoReplica` ([#28269](https://github.com/hashicorp/terraform-provider-azurerm/issues/28269))
* `azurerm_mssql_managed_instance` - allow system and user assigned identities, fix update failure ([#28319](https://github.com/hashicorp/terraform-provider-azurerm/issues/28319))
* `azurerm_storage_account` - fix error handling for `static_website` and `queue_properties` availability checks ([#28279](https://github.com/hashicorp/terraform-provider-azurerm/issues/28279))



## 4.14.0 (December 12, 2024)

BREAKING CHANGES:

* `nginx` - update api version to `2024-09-01-preview`, this API no longer supports certain properties which have had to be deprecated in the provider for the upgrade ([#27776](https://github.com/hashicorp/terraform-provider-azurerm/issues/27776))
* Data Source: `azurerm_nginx_configuration` - the `protected_file.content` property will not be populated and has been deprecated ([#27776](https://github.com/hashicorp/terraform-provider-azurerm/issues/27776))
* Data Source: `azurerm_nginx_deployment` - the `managed_resource_group` property will not be populated and has been deprecated ([#27776](https://github.com/hashicorp/terraform-provider-azurerm/issues/27776))
* `azurerm_network_function_collector_policy` - the API doesn't preserve the ordering of the `ipfx_ingestion.source_resource_ids` property causing non-empty plans after apply, this property's type has been changed from a list to a set to prevent Terraform from continually trying to recreate this resource. If this property is being referenced anywhere you will need to update your config to convert it to a list before referencing it ([#27915](https://github.com/hashicorp/terraform-provider-azurerm/issues/27915))
* `azurerm_nginx_deployment` - the `managed_resource_group` property is no longer supported and has been deprecated ([#27776](https://github.com/hashicorp/terraform-provider-azurerm/issues/27776))

FEATURES:

* **New Resource**: `azurerm_cognitive_account_rai_blocklist` ([#28043](https://github.com/hashicorp/terraform-provider-azurerm/issues/28043))
* **New Resource**: `azurerm_fabric_capacity` ([#28080](https://github.com/hashicorp/terraform-provider-azurerm/issues/28080))

ENHANCEMENTS:

* dependencies - update `go-azure-sdk` to `v0.20241206.1180327` ([#28211](https://github.com/hashicorp/terraform-provider-azurerm/issues/28211))
* `nginx` - update api version to `2024-11-01-preview` ([#28227](https://github.com/hashicorp/terraform-provider-azurerm/issues/28227))
* `azurerm_linux_function_app` - add support for  preview  value `21` for `java_version` ([#26304](https://github.com/hashicorp/terraform-provider-azurerm/issues/26304))
* `azurerm_linux_function_app_slot` - support `1.3` for `site_config.minimum_tls_version` and `site_config.scm_minimum_tls_version` ([#28016](https://github.com/hashicorp/terraform-provider-azurerm/issues/28016))
* `azurerm_linux_web_app` - add support for  preview  value `21` for `java_version` ([#26304](https://github.com/hashicorp/terraform-provider-azurerm/issues/26304))
* `azurerm_orchestrated_virtual_machine_scale_set` - support hot patching for `2025-datacenter-azure-edition-core-smalldisk` ([#28160](https://github.com/hashicorp/terraform-provider-azurerm/issues/28160))
* `azurerm_search_service` - add support for the `network_rule_bypass_option` property ([#28139](https://github.com/hashicorp/terraform-provider-azurerm/issues/28139))
* `azurerm_windows_function_app` - add support for  preview  value `21` for `java_version` ([#26304](https://github.com/hashicorp/terraform-provider-azurerm/issues/26304))
* `azurerm_windows_function_app_slot` - support `1.3` for `site_config.minimum_tls_version` and `site_config.scm_minimum_tls_version` ([#28016](https://github.com/hashicorp/terraform-provider-azurerm/issues/28016))
* `azurerm_windows_virtual_machine` - support hot patching for `2025-datacenter-azure-edition-core-smalldisk` ([#28160](https://github.com/hashicorp/terraform-provider-azurerm/issues/28160))
* `azurerm_windows_web_app` - add support for  preview  value `21` for `java_version` ([#26304](https://github.com/hashicorp/terraform-provider-azurerm/issues/26304))

BUG FIXES:

* `azurerm_management_group` - fix regression where subscription ID can't be parsed correctly anymore ([#28228](https://github.com/hashicorp/terraform-provider-azurerm/issues/28228))
