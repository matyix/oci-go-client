# OssContainerPutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | The two-part name of the account (/Compute-identity_domain/cloud_storage) that contains the credentials and access details of the associated Oracle Storage Cloud Service instance. | [default to null]
**DeleteRemote** | **bool** | Determines whether the contents of the Oracle Storage Cloud Service container is also deleted when you delete the &lt;code&gt;integrations/osscontainer&lt;/code&gt; object.&lt;p&gt;When set to &lt;code&gt;true&lt;/code&gt;, deletes the Oracle Storage Cloud Service container along with all the objects in the container when you delete the integration/osscontainer object created by this orchestration.&lt;p&gt;When set to &lt;code&gt;false&lt;/code&gt;, only the integrations/osscontainer object created by this orchestration is deleted. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


