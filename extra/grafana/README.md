# Using Pure Storage FlashBlade OpenMetrics Exporter with Prometheus and Grafana
How to setup Prometheus to scrape metrics and display dashboards in Grafana using [Pure Storage FlashBlade OpenMetrics Exporter][1].

## Support Statement
This exporter is provided under Best Efforts support by the Pure Portfolio Solutions Group, Open Source Integrations team. For feature requests and bugs please use GitHub Issues. We will address these as soon as we can, but there are no specific SLAs.

## TL;DR
1. Configure Pure Storage OpenMetrics Exporter ([pure-fb-openmetrics-exporter][1]).
2. Deploy and configure Prometheus ([prometheus-docs][2]). [Example prometheus.yaml here](../prometheus/prometheus.yaml).
3. Deploy and configure Grafana ([grafana-docs][3]).
4. Import [grafana-purefb-flashblade-overview.json](grafana-purefb-flashblade-overview.json) into Grafana.

# Overview
Take a holistic overview of your Pure Storage FlashBlade estate on-premise with Prometheus and Grafana to summarize statistics such as:
  * FlashBlade Utilization
  * Purity OS version
  * Data Reduction Rate
  * Number and type of open alerts

Drill down into specific arrays and identify top busy hosts while correlating read and write operations and throughput to quickly highlight or eliminate investigation enquiries.
<br>
<img src="./images/grafana_purefb_overview_dash_1.png" width="66%" height="66%">
<img src="./images/grafana_purefb_overview_dash_2.png" width="33%" height="33%">
<br>

These dashboards provide an overview of your fleet to give you early indications of potential issues and a look back in time to recent history. Once you are pulling the metrics, you can create your own dashboards bespoke to your environment, even correlating metrics from other technologies. 

Would you like to collect metrics from your Pure Storage FlashBlade fleet and monitor your environment on-premise? Let's walk through it.

Included in this guide
* Overview
* How does it work?
* Setup
* Troubleshooting No Data Issues
* Troubleshooting Specific Errors
* Pure Storage FlashBlade Overview Grafana Dashboard

# How does it work?
Time-series database Prometheus polls the Pure Storage OpenMetrics exporter (OME) with details of the device endpoint (Pure Storage FlashBlade).

Pure Storage OME polls the device and relays the metrics back to Prometheus for data retention.
Now we can query the Prometheus database either in Prometheus UI or using Grafana.

Grafana can be configured to query all of the metrics available in the Prometheus database and display them in meaningful dashboards. These dashboards can be imported or developed further, not just to correlate data of the FlashBlade, but exporters for other elements of your infrastructure can be queried and correlated bespoke to your environment.
<br>
<img src="./images/pure_om_exporters_prometheus_grafana.png">
<br>

# Setup
## Prerequisites and Dependencies
This deployment assumes the [Pure Storage FlashBlade OpenMetrics Exporter][1] is previously been setup and configured.
Supported operating system platforms are available to install Prometheus and Grafana.

Prometheus and Grafana are open source platforms which are under constant development, therefore updates to our overview dashboard cannot be tested with every version. When new versions of the dashboard is released we endeavor to test between current and previous tested versions.

The Grafana dashboards have been developed and tested using the following software versions:
| OME            | Prometheus  | Grafana        |
| -------------- | ----------- |--------------- |
| v1.0.7         | 2.45.1      | 9.3.2, 10.1.15, 10.2.2 |
| v1.0.7         | 2.45.0      | 9.3.2, 10.0.1 |
| v1.0.5.hotfix1 | 2.41.1      | 9.3.2, 9.4.1  |

Purity 6.1.0 and above (REST API 2.x continues to be developed in later Purity versions)
Dashboards may have limited functionality with earlier versions and some modifications may be required.

## Prometheus
1. Install Prometheus on your chosen OS platform ([prometheus-docs][2]).

2. Generate an API token for your chosen user account.
In this case we are using a remote LDAP user account with read-only permissions.

      <details>
      <summary>Expand for CLI example</summary>

      ```console
      pureuser@arrayname01> pureadmin list
      Name          Type    Role
      pureuser      local   array_admin
      o11y-readonly remote  readonly

      pureuser@arrayname01> pureadmin create --api-token o11y-readonly
      Name           Type    API Token                               Created                  Expires
      o11y-readonly  remote  T-11111111-1111-1111-1111-111111111111  2023-01-01 12:00:00 EST  -    
      ```

      </details>

      <details>
      <summary>Expand for GUI example</summary>

      ![Alt text](images/purefb_create_api_token.png)
      </details>

3. Configure `/etc/prometheus/prometheus.yaml` to point use the OpenMetrics exporter to query the device endpoint.

[This is an example of configuring the prometheus.yaml](../prometheus/prometheus.yaml)

Let's take a walkthrough an example of scraping the `/metrics/array` endpoint.

```yaml
# Scrape job for one Pure Storage FlashBlade scraping /metrics/array
# Each Prometheus scrape requires a job name. In this example we have structures the name `exporter_endpoint_arrayname`
  - job_name: 'purefb_array_arrayname01'
    # Specify the array endpoint from /metrics/array
    metrics_path: /metrics/array
    # Provide FlashBlade authorization API token
    authorization:
      credentials: T-11111111-1111-1111-1111-111111111111
    # Provide parameters to pass the exporter the device to connect to. Provide FQDN or IP address
    params:
      endpoint: ['arrayname01.fqdn.com']

    static_configs:
    # Tell Prometheus which exporter to make the request
    - targets:
      - purefa-openmetrics-exporter.fqdn.com:9491
      # Finally provide labels to the device.
      labels:
        # Instance should be the device name and is used to correlate metrics between different endpoints in Prometheus and Grafana. Ensure this is the same for each endpoint for the same device.
        instance: arrayname01
        # location, site and env are specific to your environment. Feel free to add more labels but maintain these three to minimize changes to Grafana which is expecting to use location, site and env as filter variables. 
        location: uk
        site: London
        env: production

# Repeat for the above for additional endpoints (if required) but note the limitations of collecting these in https://github.com/PureStorage-OpenConnect/pure-fb-openmetrics-exporter.
# /metrics/usage
# /metrics/clients 

# Repeat again for more Pure Storage FlashBlades
```
4. Test the prometheus.yaml file is valid

    ```console
    > promtool check config /etc/prometheus/prometheus.yml
    Checking prometheus.yml
    SUCCESS: prometheus.yml is valid prometheus config file syntax
    ```

5. Restart Prometheus to ensure changes take effect

6. Navigate to your Prometheus instance via web browser http://prometheus-server.fqdn.com:9090
  -  Type `purefb_info` in the query box and hit return
  
7. All going well, you will see your device listed:

    ```
    purefb_info{array_name="ARRAYNAME01", env="production", instance="arrayname01", job="purefb_array_arrayname01", location="uk", os="Purity//FB", site="London", system_id="11111111-1111-1111-1111-1111-111111111111", version="4.3.1"}
    ```

## Grafana

1. Install Grafana on your chosen OS platform ([grafana-docs][3]).
2. Point Grafana as your Prometheus data source ([grafana-datasource][4]]).
<br>
<img src="./images/grafana_add_datasource.png" width="40%" height="40%">
<br>

3. Download dashboard(s) from this folder postfixed with .json.
    Either copy the contents of the .json file or download the file and import it into Grafana.
    [Pure Storage FlashBlade Overview Grafana Dashboard](grafana-purefb-flashblade-overview.json)
4. Import the .json file into Grafana and specify the Prometheus datasource.
5. Now open the dashboard and check that data from your devices are visible.
<br>
<img src="./images/grafana_purefb_overview_dash_1.png" width="40%" height="40%">
<br>


# Troubleshooting No Data Issues
Whilst this guide does not replace Prometheus and Grafana support and community portals, it should give you an insight on where we might look first and where to perform checks.

## No data is visible in dashboard
Check the data is accessible at component in the stack. If at any on these points do not work, resolve them before moving on to the next component.
  * Check Pure OpenMetrics Exporter
  * Check Prometheus
  * Check Grafana

### Check Pure OpenMetrics Exporter
1. Run cURL against the exporter and pass is the bearer token and endpoint. 

    ```
    curl -H 'Authorization: Bearer T-11111111-1111-1111-1111-111111111111' -X GET 'http://<exporter_ip>:9491/metrics/array?endpoint=arrayname01.fqdn.com'
    ```

### Check Prometheus
2. Using the Prometheus UI, run a simple query to see if any results are returned.
<br>
<img src="./images/prometheus_simple_purefb_query.png" width="40%" height="40%">
<br>

3. If the query does not return results, check the status of the targets for status errors.
<br>
<img src="./images/prometheus_purefb_target_status.png" width="40%" height="40%">
<br>

4. Run prometheus.yaml through the yaml checker. Check the configuration is correct and restart Prometheus.

    ```console
    > promtool check config /etc/prometheus/prometheus.yml
    Checking prometheus.yml
    SUCCESS: prometheus.yml is valid prometheus config file syntax
    ```

5. Check messages log for Prometheus errors.

### Check Grafana
6. Perform a simple test for Grafana by navigating to 'Explore' and entering a simple query.
<br>
<img src="./images/grafana_simple_purefb_query.png" width="50%" height="50%">
<br>

# Troubleshooting Specific Errors
## Some panels have errors
If the panels have errors then the query could be unsupported. Check the versions of Prometheus and Grafana are above the prerequisites versions tested above.

## Selected graphs are not loading or taking a long time to load
Filesystems panels may take longer to load when a greater number of more FlashBlade's are selected combined with longer time ranges are selected.

The more samples Prometheus is running queries against, the longer the query will take to complete and may even time out. If we are scraping data for 10x FlashBlades every 30 seconds, each with a modest 1,000 filesystems and we set a time range to 15 days, the sum is:

`2x (samples per min) X 60 min X 24 hours X 15 days X 10 FlashBlades X 1,000 filesystems = 432,000,000 sampled metrics.`

Now Prometheus is trying to calculate the top(k) of these results.

Try reducing the scope of the query and/or increase processing and memory resources to Prometheus.

# Pure Storage FlashBlade Overview Grafana Dashboard

## Dashboard Template Variables and Filters
The dashboards are fully templated which means they will work with your FlashBlade configuration and allow you to filter by environment, array and how many top(k) metrics you wish to display on your dashboard.
FlashBlade Overview allows filtering of most panels by protocol. Selecting All provides an aggregate of all protocols.
<br>
<img src="./images/grafana_purefb_dashboard_template.png">
<br>

## Average values
Many of the dashboard panels average metrics to help smooth out graphs and identify trends rather than isolated max and min values. Adjust the time range to drill down in to specific values.
## Accessibility - Dark/Light Mode
Colors have been selected to work equally in both dark and light mode. 
<br>
<img src="./images/grafana_purefb_dark_light_modes.png">
<br>

## Threshold Defaults
### Utilization
Grafana dashboards are configured with utilization thresholds to pro-actively highlight potential capacity issues with plenty of time to address. These values can be adjusted to suit your own threshold policies.

The utilization bar and value text will change color according to the array utilization.

* <70% Base is Pure Ocean
* $\ge$ 70% is Pure Yellow
* $\ge$ 80% is Pure Orange
* $\ge$ 90% is Pure Magenta
<br>
<img src="./images/grafana_utilization_thresholds.png" width="40%" height="40%">
<br>

## Alerts
Alerts are set to highlight how many and how serious the alerts across your FlashBlade fleet.
Alert panel is ordered by Critical alerts to ensure any arrays with critical alerts are moved to the top of the table for your immediate attention. 

**Total number of alerts**
* 0  background is transparent zero & replaced with value '-'
* $\gt$ 0 is above threshold

**Alert Criticality**
* $\gt$ 0 Info Alert background is Pure Ocean
* $\gt$ 0 Warning Alert background is Pure Magenta
* $\gt$ 0 Critical Alert background is Pure Magenta

## Data Reduction Rate
Data Reduction Rate also has a threshold set for informational purposes to highlight when an array might be storing incompressible or unique datasets. This can be adjusted to suit your environment policies.
* Default Data Reduction Rate <1.01:1 = Pure Yellow
<br>
<img src="./images/purefb_grafana_dashboard_drr_threshold.png" width="40%" height="40%">
<br>

[1]: https://github.com/PureStorage-OpenConnect/pure-fb-openmetrics-exporter "pure-fb-openmetrics-exporter"
[2]: https://prometheus.io/docs/introduction/overview/ "prometheus-docs"
[3]: https://grafana.com/docs/grafana/latest/ "grafana-docs"
[4]: https://grafana.com/docs/grafana/latest/administration/data-source-management/ "grafana-datasource"
