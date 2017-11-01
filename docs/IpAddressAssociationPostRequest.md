# IpAddressAssociationPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the IP address association. | [optional] [default to null]
**IpAddressReservation** | **string** | The three-part name of the NAT IP address reservation (/Compute-identity_domain/user/object). | [optional] [default to null]
**Name** | **string** | The three-part name of the Ip address association (/Compute-identity_domain/user/object).&lt;p&gt;Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. When you specify the object name, ensure that an object of the same type and with the same name doesn&#39;t already exist. If such an object already exists, another object of the same type and with the same name won&#39;t be created and the existing object won&#39;t be updated. | [default to null]
**Tags** | **[]string** | Description of the IP address association. | [optional] [default to null]
**Vnic** | **string** | Multipart name of the virtual NIC that you want to associate with the NAT IP reservation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


