# SecurityRulePostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | **string** | Select the three-part name of the access control list (ACL) that you want to add this security rule to. Security rules are applied to vNIC sets by using ACLs. | [optional] [default to null]
**Description** | **string** | Description of the security rule. | [optional] [default to null]
**DstIpAddressPrefixSets** | **[]string** | A list of IP address prefix sets to which you want to permit traffic. Only packets to IP addresses in the specified IP address prefix sets are permitted. When no destination IP address prefix sets are specified, traffic to any IP address is permitted. | [optional] [default to null]
**DstVnicSet** | **string** | The vNICset to which you want to permit traffic. Only packets to vNICs in the specified vNICset are permitted. When no destination vNICset is specified, traffic to any vNIC is permitted. | [optional] [default to null]
**EnabledFlag** | **bool** | Allows the security rule to be enabled or disabled. This parameter is set to true by default. Specify false to disable the security rule. | [optional] [default to null]
**FlowDirection** | **string** |  Specify the direction of flow of traffic, which is relative to the instances, for this security rule. Allowed values are &lt;code&gt;ingress&lt;/code&gt; or &lt;code&gt;egress&lt;/code&gt;.&lt;p&gt;An ingress packet is a packet received by a virtual NIC, for example from another virtual NIC or from the public Internet.&lt;p&gt;An egress packet is a packet sent by a virtual NIC, for example to another virtual NIC or to the public Internet. | [default to null]
**Name** | **string** | The three-part name of the Ip address association (/Compute-identity_domain/user/object).&lt;p&gt;Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. When you specify the object name, ensure that an object of the same type and with the same name doesn&#39;t already exist. If such an object already exists, another object of the same type and with the same name won&#39;t be created and the existing object won&#39;t be updated. | [default to null]
**SecProtocols** | **[]string** | A list of security protocols for which you want to permit traffic. Only packets that match the specified protocols and ports are permitted. When no security protocols are specified, traffic using any protocol over any port is permitted. | [optional] [default to null]
**SrcIpAddressPrefixSets** | **[]string** | A list of IP address prefix sets from which you want to permit traffic. Only packets from IP addresses in the specified IP address prefix sets are permitted. When no source IP address prefix sets are specified, traffic from any IP address is permitted. | [optional] [default to null]
**SrcVnicSet** | **string** | The vNICset from which you want to permit traffic. Only packets from vNICs in the specified vNICset are permitted. When no source vNICset is specified, traffic from any vNIC is permitted. | [optional] [default to null]
**Tags** | **[]string** | Strings that you can use to tag the security rule. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


