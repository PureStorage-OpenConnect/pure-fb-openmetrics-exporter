# Semantic Conventions for Pure FlashBlade Metrics <!-- omit from toc -->

This document describes the semantic conventions for Pure FlashBlade Metrics.


<!-- toc -->

- [Collections by Endpoint](#collections-by-endpoint)
- [Metric Statuses](#metric-statuses)
- [Metric Instruments](#metric-instruments)
  - [`purefb_info` - FlashBlade System Information](#purefb_info---flashblade-system-information)
  - [`purefb_alerts` - FlashBlade Alerts Information](#purefb_alerts---flashblade-alerts-information)
  - [`purefb_array` - FlashBlade Array Metrics](#purefb_array---flashblade-array-metrics)
  - [`purefb_buckets` - FlashBlade Bucket metrics](#purefb_buckets---flashblade-bucket-metrics)
  - [`purefb_clients` - FlashBlade Client metrics](#purefb_clients---flashblade-client-metrics)
  - [`purefb_file_systems` - FlashBlade File System metrics](#purefb_file_systems---flashblade-file-system-metrics)
  - [`purefb_file_system_usage` - FlashBlade File System Usage metrics](#purefb_file_system_usage---flashblade-file-system-usage-metrics)
  - [`purefb_hardware` - FlashBlade Hardware metrics](#purefb_hardware---flashblade-hardware-metrics)
  - [`purefb_nfs_export_rule` - FlashBlade NFS Export information](#purefb_nfs_export_rule---flashblade-nfs-export-information)

<!-- tocstop -->

## Collections by Endpoint

| Endpoint          | Description             | Metrics Instruments collected                                                                              |
| ----------------- | ----------------------- | ---------------------------------------------------------------------------------------------------------- |
| /metrics          | Full array metrics      | all                                                                                                        |
| /metrics/array    | Array only metrics      | `purefb_info`, `purefb_alerts`, `purefb_array`, `purefb_buckets`, `purefb_file_systems`, `purefb_hardware` |
| /metrics/clients  | Client only metrics     | `purefb_info`, `purefb_clients`                                                                            |
| /metrics/usage    | Quota only metrics      | `purefb_info`, `purefb_file_system_usage`                                                                  |
| /metrics/policies | NFS policy related info | `purefb_info`, `purefb_nfs_export_rule`                                                                    |

## Metric Statuses

| Status         | Description                                                                                                                          |
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| Available      | The metric is available for use with the OME in supported LLR(ER) Purity versions.                                                   |
| New            | The metric has been recently released within the OME and may only be available for use with the latest MRB(SR) Purity versions.      |
| In Development | The metric is available in a Purity API version and is in development for use within the OME.                                        |
| Accepted       | The metric may not be currently available in the Purity API, but has been selected for development for the OME if or when available. |
| Proposed       | The metric may not be available in the Purity API, but is not in development for the OME yet.                                        |
| Abandoned      | The metric is no longer supported, or in development. If the metric name has changed, the new metric name will be linked.            |


## Metric Instruments

### `purefb_info` - FlashBlade System Information

**Description:** FlashBlade System Information

| Status    | Name        | Description                   | Units | Metric Type ([*](https://github.com/OpenObservability/OpenMetrics/blob/main/specification/OpenMetrics.md#metric-types)) | Value Type | Attribute Key | Attribute Values       |
| --------- | ----------- | ----------------------------- | ----- | ----------------------------------------------------------------------------------------------------------------------- | ---------- | ------------- | ---------------------- |
| Available | purefb_info | FlashBlade system information |       | Gauge                                                                                                                   | Double     | `array_name`  | (name)                 |
|           |             |                               |       |                                                                                                                         |            | `os`          | (array os name)        |
|           |             |                               |       |                                                                                                                         |            | `system_id`   | (array system id)      |
|           |             |                               |       |                                                                                                                         |            | `version`     | (array purity version) |


### `purefb_alerts` - FlashBlade Alerts Information

**Description:** FlashBlade Open Alerts

| Status    | Name               | Description                  | Units | Metric Type ([*](https://github.com/OpenObservability/OpenMetrics/blob/main/specification/OpenMetrics.md#metric-types)) | Value Type | Attribute Key    | Attribute Values              |
| --------- | ------------------ | ---------------------------- | ----- | ----------------------------------------------------------------------------------------------------------------------- | ---------- | ---------------- | ----------------------------- |
| Available | purefb_alerts_open | FlashBlade open alert events |       | Gauge                                                                                                                   | Double     | `name`           | (name)                        |
|           |                    |                              |       |                                                                                                                         |            | `action`         | (action)                      |
|           |                    |                              |       | The code number of the event                                                                                            |            | `code`           | (code)                        |
|           |                    |                              |       |                                                                                                                         |            | `component_name` | (component name)              |
|           |                    |                              |       |                                                                                                                         |            | `component_type` | (component type)              |
|           |                    |                              |       | The time the alert was created in milliseconds since the UNIX epoch                                                     |            | `created`        | (created)                     |
|           |                    |                              |       | Knowledge Base URL related to the alert                                                                                 |            | `kburl`          | (kburl)                       |
|           |                    |                              |       |                                                                                                                         |            | `severity`       | `info`, `warning`, `critical` |
|           |                    |                              |       | A summary of the alert                                                                                                  |            | `summary`        | (summary)                     |

### `purefb_array` - FlashBlade Array Metrics

**Description:** TODO

### `purefb_buckets` - FlashBlade Bucket metrics

| Status                                                                                                                                                                                                                    | Name                                                   | Description                                                        | Units            | Metric Type ([*](https://github.com/OpenObservability/OpenMetrics/blob/main/specification/OpenMetrics.md#metric-types)) | Value Type | Attribute Key        | Attribute Values                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------ | ------------------------------------------------------------------ | ---------------- | ----------------------------------------------------------------------------------------------------------------------- | ---------- | -------------------- | ------------------------------------------------------------------------------------------------------------------ |
| Available                                                                                                                                                                                                                 | purefb_buckets_object_count                            | FlashBlade buckets object count                                    |                  | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
| Available                                                                                                                                                                                                                 | purefb_buckets_performance_average_bytes               | FlashBlade buckets average operations size in bytes                | byte             | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
|                                                                                                                                                                                                                           |                                                        |                                                                    |                  |                                                                                                                         |            | `dimension`          | `bytes_per_op`, `bytes_per_read`, `bytes_per_write`                                                                |
| Available                                                                                                                                                                                                                 | purefb_buckets_performance_bandwidth_bytes             | FlashBlade buckets bandwidth in bytes per second                   | byte/second      | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
|                                                                                                                                                                                                                           |                                                        |                                                                    |                  |                                                                                                                         |            | `dimension`          | `read_bytes_per_sec`, `write_bytes_per_sec`                                                                        |
| Available                                                                                                                                                                                                                 | purefb_buckets_performance_latency_usec                | FlashBlade buckets latency in microseconds                         | microsecond      | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
|                                                                                                                                                                                                                           |                                                        |                                                                    |                  |                                                                                                                         |            | `dimension`          | `usec_per_other_op`, `usec_per_read_op`, `usec_per_write_op`                                                       |
| Available                                                                                                                                                                                                                 | purefb_buckets_performance_throughput_iops             | FlashBlade buckets throughput in operations per second             | operation/second | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
|                                                                                                                                                                                                                           |                                                        |                                                                    |                  |                                                                                                                         |            | `dimension`          | `others_per_sec`, `reads_per_sec`, `writes_per_sec`                                                                |
| Available                                                                                                                                                                                                                 | purefb_buckets_quota_space_bytes                       | FlashBlade buckets quota space in bytes                            | byte             | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
|                                                                                                                                                                                                                           |                                                        |                                                                    |                  |                                                                                                                         |            | `hard_limit_enabled` | `true`,`false`                                                                                                     |
| Available                                                                                                                                                                                                                 | purefb_buckets_s3_specific_performance_latency_usec    | FlashBlade buckets S3 specific latency in microseconds             | microsecond      | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
|                                                                                                                                                                                                                           |                                                        |                                                                    |                  |                                                                                                                         |            | `dimension`          | `usec_per_other_op`, `usec_per_read_bucket_op`, `usec_per_read_object_op`, `usec_per_write_bucket_op`              |
| Available                                                                                                                                                                                                                 | purefb_buckets_s3_specific_performance_throughput_iops | FlashBlade buckets S3 specific throughput in operations per second | second           | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
|                                                                                                                                                                                                                           |                                                        |                                                                    |                  |                                                                                                                         |            | `dimension`          | `others_per_sec`, `read_buckets_per_sec`, `read_objects_per_sec`, `write_buckets_per_sec`, `write_objects_per_sec` |
| Available <br />[Note: object_count space dimension to be removed in a future release in favor of purefb_buckets_object_count released in v1.0.8](https://github.com/PureStorage-OpenConnect/pure-fb-openmetrics-exporter/pull/42) | purefb_buckets_space_bytes                             | FlashBlade buckets space in bytes                                  | byte             | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |
|                                                                                                                                                                                                                           |                                                        |                                                                    |                  |                                                                                                                         |            | `space`              | `snapshots`, `total_physical`, `unique`, `virutal`, `object_count`                                                 |
| Available                                                                                                                                                                                                                 | purefb_buckets_space_data_reduction_ratio              | FlashBlade buckets space data reduction ratio                      | ratio            | Gauge                                                                                                                   | Double     | `name`               | (name)                                                                                                             |


### `purefb_clients` - FlashBlade Client metrics

**Description:** TODO

### `purefb_file_systems` - FlashBlade File System metrics

**Description:** TODO

### `purefb_file_system_usage` - FlashBlade File System Usage metrics

**Description:** TODO

### `purefb_hardware` - FlashBlade Hardware metrics

**Description:** TODO

### `purefb_nfs_export_rule` - FlashBlade NFS Export information

**Description:** TODO
