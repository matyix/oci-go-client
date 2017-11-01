# VpnConnectionResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerVpnGateway** | **string** | IP address of the VPN gateway in your data center through which you want to connect to the Oracle Cloud VPN gateway. | [optional] [default to null]
**IkeIdentifier** | **string** | The Internet Key Exchange (IKE) ID that you have specified. If you don&#39;t specify a value, the default value is the public IP address of the cloud gateway. | [optional] [default to null]
**IpNetwork** | **string** |  The name of the IP network on which the cloud gateway is created by VPNaaS. | [optional] [default to null]
**LocalGatewayAddress** | **string** | IP address of your cloud gateway. An IP address that is available in the IP network you specify is assigned to the cloud gateway. | [optional] [default to null]
**Name** | **string** | Name that you have specified for this VPN connection. | [optional] [default to null]
**PfsFlag** | **bool** | &lt;code&gt;True&lt;/code&gt; indicates that Perfect Forward Secrecy (PFS) is required and your third-party device supports PFS. | [optional] [default to null]
**Phase1Settings** | [**map[string]interface{}**](interface{}.md) | Settings for Phase 1 of protocol (IKE). | [optional] [default to null]
**Phase2Settings** | [**map[string]interface{}**](interface{}.md) | Settings for Phase 2 of protocol (IPSEC). | [optional] [default to null]
**Psk** | **string** | The pre-shared VPN key. | [optional] [default to null]
**ReachableRoutes** | **[]string** | List of subnets (CIDR prefixes) that are reachable through this VPN tunnel. | [optional] [default to null]
**TunnelStatus** | **string** | Current status of the tunnel. The tunnel can be in one of the following states:&lt;p&gt;* PENDING: indicates that your VPN connection is being set up.&lt;p&gt;* UP: indicates that your VPN connection is established.&lt;p&gt;* DOWN: indicates that your VPN connection is down.&lt;p&gt;* ERROR: indicates that your VPN connection is in the error state. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier. | [optional] [default to null]
**VnicSets** | **[]string** | Comma-separated list of vNIC sets. Traffic is allowed to and from these vNIC sets to the cloud gateway&#39;s vNIC set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


