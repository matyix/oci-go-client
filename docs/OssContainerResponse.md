# OssContainerResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | The two-part name of the account (&lt;code&gt;/Compute-&lt;i&gt;identity_domain&lt;/i&gt;/cloud_storage&lt;/code&gt;) that contains the credentials and access details of the associated Oracle Storage Cloud Service instance. | [optional] [default to null]
**Container** | **string** | The name of the container. | [optional] [default to null]
**CreationTime** | **string** | Provides the time when the OSS container was created. | [optional] [default to null]
**DeleteRemote** | **bool** | When set to &lt;code&gt;true&lt;/code&gt;, deletes the Oracle Storage Cloud Service container along with all the objects in the container when you delete the integration/osscontainer object created by this orchestration. When set to &lt;code&gt;false&lt;/code&gt;, only the integrations/osscontainer object created by this orchestration is deleted. By default, this option is set to &lt;code&gt;true&lt;/code&gt;. | [optional] [default to null]
**DeletionProgress** | **string** | If delete operation is being carried out, this attribute indicates the current progress of the delete operation. Ignore this attribute value for other operations. | [optional] [default to null]
**ErrorReason** | **string** | Describes the cause for the OSS container object to go into error state. | [optional] [default to null]
**GeoreplicationDetails** | [**map[string]interface{}**](interface{}.md) | Details applying to the georeplication policy | [optional] [default to null]
**GeoreplicationPolicy** | **string** | Optional georeplication policy string | [optional] [default to null]
**ModificationTime** | **string** | Time when the OSS container object was last updated. | [optional] [default to null]
**Name** | **string** | Multipart name of the object. | [optional] [default to null]
**State** | **string** | Current state of the OSS container object. If the object is in the &lt;code&gt;error&lt;/code&gt; state, look at &lt;code&gt;error_reason&lt;/code&gt; to identify the cause of this error. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


