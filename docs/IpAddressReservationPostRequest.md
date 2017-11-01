# IpAddressReservationPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the IP address reservation. | [optional] [default to null]
**IpAddressPool** | **string** | The IP address pool from which you want to reserve an IP address. Specify one of the following values:&lt;p&gt;* &lt;code&gt;/oracle/public/public-ippool&lt;/code&gt;: A pool of public IP addresses. An IP address from this pool is accessible over the public Internet.&lt;p&gt;* &lt;code&gt;/oracle/public/cloud-ippool&lt;/code&gt;: A pool of cloud IP addresses. An IP address from this pool is accessible to other IP networks in the Oracle cloud. You can use these IP addresses to communicate with other Oracle services. | [optional] [default to null]
**Name** | **string** | The three-part name of the IP address reservation(/Compute-identity_domain/user/object).&lt;p&gt;Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. When you specify the object name, ensure that an object of the same type and with the same name doesn&#39;t already exist. If such an object already exists, another object of the same type and with the same name won&#39;t be created and the existing object won&#39;t be updated. | [default to null]
**Tags** | **[]string** | Strings that you can use to tag the IP address reservation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


