
---
title: "base.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `solo.io.envoy.config.core.v3` 
#### Types:


- [Locality](#locality)
- [BuildVersion](#buildversion)
- [Extension](#extension)
- [Node](#node)
- [Metadata](#metadata)
- [RuntimeUInt32](#runtimeuint32)
- [RuntimeDouble](#runtimedouble)
- [RuntimeFeatureFlag](#runtimefeatureflag)
- [HeaderValue](#headervalue)
- [HeaderValueOption](#headervalueoption)
- [HeaderMap](#headermap)
- [DataSource](#datasource)
- [RetryPolicy](#retrypolicy)
- [RemoteDataSource](#remotedatasource)
- [AsyncDataSource](#asyncdatasource)
- [TransportSocket](#transportsocket)
- [RuntimeFractionalPercent](#runtimefractionalpercent)
- [ControlPlane](#controlplane)
  

 

##### Enums:


	- [RoutingPriority](#routingpriority)
	- [RequestMethod](#requestmethod)
	- [TrafficDirection](#trafficdirection)



##### Source File: [github.com/solo-io/gloo/projects/controller/api/external/envoy/config/core/v3/base.proto](https://github.com/solo-io/gloo/blob/main/projects/controller/api/external/envoy/config/core/v3/base.proto)





---
### Locality

 
Identifies location of where either Envoy runs or where upstream hosts run.

```yaml
"region": string
"zone": string
"subZone": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `region` | `string` | Region this zone belongs to. |
| `zone` | `string` | Defines the local service zone where Envoy is running. Though optional, it should be set if discovery service routing is used and the discovery service exposes zone data, either in this message or via `--service-zone`. The meaning of zone is context dependent, e.g. [Availability Zone (AZ)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html) on AWS, [Zone](https://cloud.google.com/compute/docs/regions-zones/) on GCP, etc. |
| `subZone` | `string` | When used for locality of upstream hosts, this field further splits zone into smaller chunks of sub-zones so they can be load balanced independently. |




---
### BuildVersion

 
BuildVersion combines SemVer version of extension with free-form build information
(i.e. 'alpha', 'private-build') as a set of strings.

```yaml
"version": .solo.io.envoy.type.v3.SemanticVersion
"metadata": .google.protobuf.Struct

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `version` | [.solo.io.envoy.type.v3.SemanticVersion](../../../../type/v3/semantic_version.proto.sk/#semanticversion) | SemVer version of extension. |
| `metadata` | [.google.protobuf.Struct](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/struct) | Free-form build information. Envoy defines several well known keys in the source/common/common/version.h file. |




---
### Extension

 
Version and identification for an Envoy extension.
[#next-free-field: 6]

```yaml
"name": string
"category": string
"typeDescriptor": string
"version": .solo.io.envoy.config.core.v3.BuildVersion
"disabled": bool

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `name` | `string` | This is the name of the Envoy filter as specified in the Envoy configuration, e.g. envoy.filters.http.router, com.acme.widget. |
| `category` | `string` | Category of the extension. Extension category names use reverse DNS notation. For instance "envoy.filters.listener" for Envoy's built-in listener filters or "com.acme.filters.http" for HTTP filters from acme.com vendor. [#comment:TODO(yanavlasov): Link to the doc with existing envoy category names.]. |
| `typeDescriptor` | `string` | [#not-implemented-hide:] Type descriptor of extension configuration proto. [#comment:TODO(yanavlasov): Link to the doc with existing configuration protos.] [#comment:TODO(yanavlasov): Add tests when PR #9391 lands.]. |
| `version` | [.solo.io.envoy.config.core.v3.BuildVersion](../base.proto.sk/#buildversion) | The version is a property of the extension and maintained independently of other extensions and the Envoy API. This field is not set when extension did not provide version information. |
| `disabled` | `bool` | Indicates that the extension is present but was disabled via dynamic configuration. |




---
### Node

 
Identifies a specific Envoy instance. The node identifier is presented to the
management server, which may use this identifier to distinguish per Envoy
configuration for serving.
[#next-free-field: 12]

```yaml
"id": string
"cluster": string
"metadata": .google.protobuf.Struct
"locality": .solo.io.envoy.config.core.v3.Locality
"userAgentName": string
"userAgentVersion": string
"userAgentBuildVersion": .solo.io.envoy.config.core.v3.BuildVersion
"extensions": []solo.io.envoy.config.core.v3.Extension
"clientFeatures": []string
"listeningAddresses": []solo.io.envoy.config.core.v3.Address

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `id` | `string` | An opaque node identifier for the Envoy node. This also provides the local service node name. It should be set if any of the following features are used: statsd, :ref:`CDS <config_cluster_manager_cds>`, and :ref:`HTTP tracing <arch_overview_tracing>`, either in this message or via `--service-node`. |
| `cluster` | `string` | Defines the local service cluster name where Envoy is running. Though optional, it should be set if any of the following features are used: statsd, :ref:`health check cluster verification <envoy_api_field_config.core.v3.HealthCheck.HttpHealthCheck.service_name_matcher>`, runtime override directory, :ref:`user agent addition <envoy_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.add_user_agent>`, HTTP global rate limiting, CDS, and :ref:`HTTP tracing <arch_overview_tracing>`, either in this message or via `--service-cluster`. |
| `metadata` | [.google.protobuf.Struct](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/struct) | Opaque metadata extending the node identifier. Envoy will pass this directly to the management server. |
| `locality` | [.solo.io.envoy.config.core.v3.Locality](../base.proto.sk/#locality) | Locality specifying where the Envoy instance is running. |
| `userAgentName` | `string` | Free-form string that identifies the entity requesting config. E.g. "envoy" or "grpc". |
| `userAgentVersion` | `string` | Free-form string that identifies the version of the entity requesting config. E.g. "1.12.2" or "abcd1234", or "SpecialEnvoyBuild". Only one of `userAgentVersion` or `userAgentBuildVersion` can be set. |
| `userAgentBuildVersion` | [.solo.io.envoy.config.core.v3.BuildVersion](../base.proto.sk/#buildversion) | Structured version of the entity requesting config. Only one of `userAgentBuildVersion` or `userAgentVersion` can be set. |
| `extensions` | [[]solo.io.envoy.config.core.v3.Extension](../base.proto.sk/#extension) | List of extensions and their versions supported by the node. |
| `clientFeatures` | `[]string` | Client feature support list. These are well known features described in the Envoy API repository for a given major version of an API. Client features use reverse DNS naming scheme, for example `com.acme.feature`. See the list of features that xDS client may support. |
| `listeningAddresses` | [[]solo.io.envoy.config.core.v3.Address](../address.proto.sk/#address) | Known listening ports on the node as a generic hint to the management server for filtering listeners to be returned. For example, if there is a listener bound to port 80, the list can optionally contain the SocketAddress `(0.0.0.0,80)`. The field is optional and just a hint. |




---
### Metadata

 
Metadata provides additional inputs to filters based on matched listeners,
filter chains, routes and endpoints. It is structured as a map, usually from
filter name (in reverse DNS format) to metadata specific to the filter. Metadata
key-values for a filter are merged as connection and request handling occurs,
with later values for the same key overriding earlier values.

An example use of metadata is providing additional values to
http_connection_manager in the envoy.http_connection_manager.access_log
namespace.

Another example use of metadata is to per service config info in cluster metadata, which may get
consumed by multiple filters.

For load balancing, Metadata provides a means to subset cluster endpoints.
Endpoints have a Metadata object associated and routes contain a Metadata
object to match against. There are some well defined metadata used today for
this purpose:

* `{"envoy.lb": {"canary": <bool> }}` This indicates the canary status of an
  endpoint and is also used during header processing
  (x-envoy-upstream-canary) and for stats purposes.
[#next-major-version: move to type/metadata/v2]

```yaml
"filterMetadata": map<string, .google.protobuf.Struct>

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `filterMetadata` | `map<string, .google.protobuf.Struct>` | Key is the reverse DNS filter name, e.g. com.acme.widget. The envoy.* namespace is reserved for Envoy's built-in filters. |




---
### RuntimeUInt32

 
Runtime derived uint32 with a default when not specified.

```yaml
"defaultValue": int
"runtimeKey": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `defaultValue` | `int` | Default value if runtime value is not available. |
| `runtimeKey` | `string` | Runtime key to get value for comparison. This value is used if defined. |




---
### RuntimeDouble

 
Runtime derived double with a default when not specified.

```yaml
"defaultValue": float
"runtimeKey": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `defaultValue` | `float` | Default value if runtime value is not available. |
| `runtimeKey` | `string` | Runtime key to get value for comparison. This value is used if defined. |




---
### RuntimeFeatureFlag

 
Runtime derived bool with a default when not specified.

```yaml
"defaultValue": .google.protobuf.BoolValue
"runtimeKey": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `defaultValue` | [.google.protobuf.BoolValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/bool-value) | Default value if runtime value is not available. |
| `runtimeKey` | `string` | Runtime key to get value for comparison. This value is used if defined. The boolean value must be represented via its [canonical JSON encoding](https://developers.google.com/protocol-buffers/docs/proto3#json). |




---
### HeaderValue

 
Header name/value pair.

```yaml
"key": string
"value": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `key` | `string` | Header name. |
| `value` | `string` | Header value. The same format specifier as used for HTTP access logging applies here, however unknown header values are replaced with the empty string instead of `-`. |




---
### HeaderValueOption

 
Header name/value pair plus option to control append behavior.

```yaml
"header": .solo.io.envoy.config.core.v3.HeaderValue
"append": .google.protobuf.BoolValue

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `header` | [.solo.io.envoy.config.core.v3.HeaderValue](../base.proto.sk/#headervalue) | Header name/value pair that this option applies to. |
| `append` | [.google.protobuf.BoolValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/bool-value) | Should the value be appended? If true (default), the value is appended to existing values. |




---
### HeaderMap

 
Wrapper for a set of headers.

```yaml
"headers": []solo.io.envoy.config.core.v3.HeaderValue

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `headers` | [[]solo.io.envoy.config.core.v3.HeaderValue](../base.proto.sk/#headervalue) |  |




---
### DataSource

 
Data source consisting of either a file or an inline value.

```yaml
"filename": string
"inlineBytes": bytes
"inlineString": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `filename` | `string` | Local filesystem data source. Only one of `filename`, `inlineBytes`, or `inlineString` can be set. |
| `inlineBytes` | `bytes` | Bytes inlined in the configuration. Only one of `inlineBytes`, `filename`, or `inlineString` can be set. |
| `inlineString` | `string` | String inlined in the configuration. Only one of `inlineString`, `filename`, or `inlineBytes` can be set. |




---
### RetryPolicy

 
The message specifies the retry policy of remote data source when fetching fails.

```yaml
"retryBackOff": .solo.io.envoy.config.core.v3.BackoffStrategy
"numRetries": .google.protobuf.UInt32Value

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `retryBackOff` | [.solo.io.envoy.config.core.v3.BackoffStrategy](../backoff.proto.sk/#backoffstrategy) | Specifies parameters that control retry backoff strategy. This parameter is optional, in which case the default base interval is 1000 milliseconds. The default maximum interval is 10 times the base interval. |
| `numRetries` | [.google.protobuf.UInt32Value](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/u-int-32-value) | Specifies the allowed number of retries. This parameter is optional and defaults to 1. |




---
### RemoteDataSource

 
The message specifies how to fetch data from remote and how to verify it.

```yaml
"httpUri": .solo.io.envoy.config.core.v3.HttpUri
"sha256": string
"retryPolicy": .solo.io.envoy.config.core.v3.RetryPolicy

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `httpUri` | [.solo.io.envoy.config.core.v3.HttpUri](../http_uri.proto.sk/#httpuri) | The HTTP URI to fetch the remote data. |
| `sha256` | `string` | SHA256 string for verifying data. |
| `retryPolicy` | [.solo.io.envoy.config.core.v3.RetryPolicy](../base.proto.sk/#retrypolicy) | Retry policy for fetching remote data. |




---
### AsyncDataSource

 
Async data source which support async data fetch.

```yaml
"local": .solo.io.envoy.config.core.v3.DataSource
"remote": .solo.io.envoy.config.core.v3.RemoteDataSource

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `local` | [.solo.io.envoy.config.core.v3.DataSource](../base.proto.sk/#datasource) | Local async data source. Only one of `local` or `remote` can be set. |
| `remote` | [.solo.io.envoy.config.core.v3.RemoteDataSource](../base.proto.sk/#remotedatasource) | Remote async data source. Only one of `remote` or `local` can be set. |




---
### TransportSocket

 
Configuration for transport socket in listeners and
clusters. If the configuration is
empty, a default transport socket implementation and configuration will be
chosen based on the platform and existence of tls_context.

```yaml
"name": string
"typedConfig": .google.protobuf.Any

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `name` | `string` | The name of the transport socket to instantiate. The name must match a supported transport socket implementation. |
| `typedConfig` | [.google.protobuf.Any](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any) |  |




---
### RuntimeFractionalPercent

 
Runtime derived FractionalPercent with defaults for when the numerator or denominator is not
specified via a runtime key.

**Note**:

  Parsing of the runtime key's data is implemented such that it may be represented as a
  FractionalPercent proto represented as JSON/YAML
  and may also be represented as an integer with the assumption that the value is an integral
  percentage out of 100. For instance, a runtime key lookup returning the value "42" would parse
  as a `FractionalPercent` whose numerator is 42 and denominator is HUNDRED.

```yaml
"defaultValue": .solo.io.envoy.type.v3.FractionalPercent
"runtimeKey": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `defaultValue` | [.solo.io.envoy.type.v3.FractionalPercent](../../../../type/v3/percent.proto.sk/#fractionalpercent) | Default value if the runtime value's for the numerator/denominator keys are not available. |
| `runtimeKey` | `string` | Runtime key for a YAML representation of a FractionalPercent. |




---
### ControlPlane

 
Identifies a specific ControlPlane instance that Envoy is connected to.

```yaml
"identifier": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `identifier` | `string` | An opaque control plane identifier that uniquely identifies an instance of control plane. This can be used to identify which control plane instance, the Envoy is connected to. |



  
### RoutingPriority

Description: Envoy supports :ref:`upstream priority routing
<arch_overview_http_routing_priority>` both at the route and the virtual
cluster level. The current priority implementation uses different connection
pool and circuit breaking settings for each priority level. This means that
even for HTTP/2 requests, two physical connections will be used to an
upstream host. In the future Envoy will likely support true HTTP/2 priority
over a single upstream connection.

| Name | Description |
| ----- | ----------- | 
| DEFAULT |  |
| HIGH |  |
  
### RequestMethod

Description: HTTP request method.

| Name | Description |
| ----- | ----------- | 
| METHOD_UNSPECIFIED |  |
| GET |  |
| HEAD |  |
| POST |  |
| PUT |  |
| DELETE |  |
| CONNECT |  |
| OPTIONS |  |
| TRACE |  |
| PATCH |  |
  
### TrafficDirection

Description: Identifies the direction of the traffic relative to the local Envoy.

| Name | Description |
| ----- | ----------- | 
| UNSPECIFIED | Default option is unspecified. |
| INBOUND | The transport is used for incoming traffic. |
| OUTBOUND | The transport is used for outgoing traffic. |


<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
