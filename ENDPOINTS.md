# FlashBlade REST endpoints monitored by this exporter

This document lists the **Pure Storage FlashBlade REST API endpoints** that the exporter calls in order to produce Prometheus/OpenMetrics output.

## REST base URL and authentication flow

- **REST base URL (after login)**: `https://<flashblade>/api/<api_version>`
- **API version discovery**: `GET https://<flashblade>/api/api_version`
- **Login**: `POST https://<flashblade>/api/login` with header `api-token: <token>` → receives `x-auth-token`
- **Authenticated requests**: all subsequent calls use `x-auth-token` header and the versioned base URL (`/api/<api_version>`)
- **Logout**: `POST https://<flashblade>/api/logout` with header `x-auth-token: <token>`

## Endpoints currently used for metrics collection

**Legend**

- All paths below are **relative to** `https://<flashblade>/api/<api_version>`
- Query parameters shown reflect what the exporter actually sends today.

| Exporter scrape group | Collector / call site | HTTP | Endpoint | Query parameters | Notes |
| --- | --- | --- | --- | --- | --- |
| all / array | `ArraysCollector` | GET | `/arrays` | - | Used for `purefb_info` |
| all / array | `AlertsCollector` | GET | `/alerts` | `filter=state='open'` | Only “open” alerts are collected |
| all / array | `PerfCollector` | GET | `/arrays/performance` | `protocol=all` | Also queried with other protocol values below |
| all / array | `PerfCollector` | GET | `/arrays/performance` | `protocol=HTTP` |  |
| all / array | `PerfCollector` | GET | `/arrays/performance` | `protocol=SMB` |  |
| all / array | `PerfCollector` | GET | `/arrays/performance` | `protocol=NFS` |  |
| all / array | `PerfCollector` | GET | `/arrays/performance` | `protocol=S3` |  |
| all / array | `HttpPerfCollector` | GET | `/arrays/http-specific-performance` | - |  |
| all / array | `NfsPerfCollector` | GET | `/arrays/nfs-specific-performance` | - |  |
| all / array | `S3PerfCollector` | GET | `/arrays/s3-specific-performance` | - |  |
| all / array | `PerfReplicationCollector` | GET | `/arrays/performance/replication` | - | Used for replication throughput and `continuous.object_backlog` |
| all / array | `ArraySpaceCollector` | GET | `/arrays/space` | `type=array` | Also queried with other `type` values below |
| all / array | `ArraySpaceCollector` | GET | `/arrays/space` | `type=file-system` |  |
| all / array | `ArraySpaceCollector` | GET | `/arrays/space` | `type=object-store` |  |
| all / array | `HardwareCollector` | GET | `/hardware` | - |  |
| all / array | `HwConnectorsPerfCollector` | GET | `/hardware-connectors/performance` | - |  |
| all / filesystems | `Collector()` (preload) | GET | `/file-systems` | - | Called to preload filesystem names/ids for downstream collectors |
| all / filesystems | `FileSystemsPerfCollector` | GET | `/file-systems/performance` | `names=<csv>` + `protocol=NFS` | `names` is a comma-separated list, requested in batches of up to 5 file systems per call |
| all / filesystems | `FileSystemsSpaceCollector` | (none) | - | - | Space/quota-related filesystem metrics are derived from the `/file-systems` response (no extra REST call) |
| all / clients | `ClientsPerfCollector` | GET | `/arrays/clients/performance` | - |  |
| all / objectstore | `Collector()` (preload) | GET | `/buckets` | `destroyed=false` | Used to preload bucket names/space/quota; **note**: on a 401 retry, the current implementation re-issues the request without the `destroyed=false` param |
| all / objectstore | `BucketsPerfCollector` | GET | `/buckets/performance` | `names=<csv>` | `names` is a comma-separated list, requested in batches of up to 5 buckets per call |
| all / objectstore | `BucketsS3PerfCollector` | GET | `/buckets/s3-specific-performance` | `names=<csv>` | `names` is a comma-separated list, requested in batches of up to 5 buckets per call |
| all / objectstore | `BucketsSpaceCollector` | (none) | - | - | Bucket space/quota/object_count metrics are derived from the `/buckets` response (no extra REST call) |
| all / objectstore | `ObjectStoreAccountsCollector` | GET | `/object-store-accounts` | - |  |
| all / usage | `Collector()` (preload) | GET | `/file-systems` | - | Called to get filesystem IDs; then usage is queried per filesystem |
| all / usage | `UsageCollector` | GET | `/usage/users` | `file_system_ids=<filesystem_id>` | Called once **per filesystem** |
| all / usage | `UsageCollector` | GET | `/usage/groups` | `file_system_ids=<filesystem_id>` | Called once **per filesystem** |
| all / policies | `NfsPoliciesCollector` | GET | `/nfs-export-policies` | - |  |

## Implemented in the repo but not currently scraped/monitored

These endpoints exist in `internal/rest-client/` but are **not called by any collector** today:

- `GET /blades`




