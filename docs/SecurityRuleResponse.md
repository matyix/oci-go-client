# SecurityRuleResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | **string** | Name of the ACL that contains this rule. | [optional] [default to null]
**Description** | **string** | Description of the object. | [optional] [default to null]
**DstIpAddressPrefixSets** | **[]string** | List of IP address prefix set names to match the packet&#39;s destination IP address. | [optional] [default to null]
**DstVnicSet** | **string** | Name of virtual NIC set containing the packet&#39;s destination virtual NIC. | [optional] [default to null]
**EnabledFlag** | **bool** | Allows the security rule to be disabled. | [optional] [default to null]
**FlowDirection** | **string** | Direction of the flow; Can be \&quot;egress\&quot; or \&quot;ingress\&quot;. | [optional] [default to null]
**Name** | **string** | Name of the security rule. | [optional] [default to null]
**SecProtocols** | **[]string** | List of security protocol object names to match the packet&#39;s protocol and port. | [optional] [default to null]
**SrcIpAddressPrefixSets** | **[]string** | List of multipart names of IP address prefix set to match the packet&#39;s source IP address. | [optional] [default to null]
**SrcVnicSet** | **string** | Name of virtual NIC set containing the packet&#39;s source virtual NIC. | [optional] [default to null]
**Tags** | **[]string** | Tags associated with the object. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


