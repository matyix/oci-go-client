# SecIpListPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | &lt;p&gt;A description of the security IP list. | [optional] [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).&lt;p&gt;Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. | [default to null]
**Secipentries** | **[]string** | &lt;p&gt;A comma-separated list of the subnets (in CIDR format) or IPv4 addresses for which you want to create this security IP list.&lt;p&gt;For example, to create a security IP list containing the IP addresses 203.0.113.1 and 203.0.113.2, enter one of the following:&lt;p&gt;&lt;code&gt;203.0.113.0/30&lt;/code&gt;&lt;p&gt;&lt;code&gt;203.0.113.1, 203.0.113.2&lt;/code&gt; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

