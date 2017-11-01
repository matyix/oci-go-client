# PrivateGatewayPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the object. | [optional] [default to null]
**IpNetworks** | **[]string** | List of IP networks that you want associate with the private gateway. When you set up Oracle Network Cloud Service - FastConnect private peering, you can directly access the instances on these IP networks from your premises without using NAT or setting up an IPSec tunnel. | [optional] [default to null]
**Name** | **string** | The three-part name of the private gateway (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). Object names can contain only alphanumeric, underscore (_), dash (-), and period (.) characters. Object names are case-sensitive. | [default to null]
**Tags** | **[]string** | Tags associated with the object. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


