# OssContainerPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | The two-part name of the account (/Compute-identity_domain/cloud_storage) that contains the credentials and access details of the associated Oracle Storage Cloud Service instance. | [default to null]
**Container** | **string** | The name of the container that you want to create.&lt;p&gt;Container names must:&lt;p&gt;* Contain only UTF-8 characters&lt;p&gt;* Be a maximum of 256 bytes&lt;p&gt;* Avoid using a slash (/) character because this character acts as a delimiter between the container name and the object name&lt;p&gt;When you specify a container name, ensure that a container of the same name doesn&#39;t already exist. | [optional] [default to null]
**DeleteRemote** | **bool** | Determines whether the contents of the Oracle Storage Cloud Service container is also deleted when you delete the &lt;code&gt;integrations/osscontainer&lt;/code&gt; object.&lt;p&gt;When set to &lt;code&gt;true&lt;/code&gt;, deletes the Oracle Storage Cloud Service container along with all the objects in the container when you delete the integration/osscontainer object created by this orchestration.&lt;p&gt;When set to &lt;code&gt;false&lt;/code&gt;, only the integrations/osscontainer object created by this orchestration is deleted. | [default to null]
**Name** | **string** | The three-part name of the integrations/osscontainer object created by this orchestration. This name is in the format /Compute-identity_domain/user/object.&lt;p&gt;If you don&#39;t specify a name for this object, then the name is generated automatically.&lt;p&gt;Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive.&lt;p&gt;When you specify the object name, ensure that an object of the same type and with the same name doesn&#39;t already exist. If such an object already exists, another object of the same type and with the same name won&#39;t be created and the existing object won&#39;t be updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


