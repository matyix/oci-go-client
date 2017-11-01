# OrchestrationV2PutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/default&lt;/code&gt; | [optional] [default to null]
**Description** | **string** | Description of this orchestration | [optional] [default to null]
**DesiredState** | **string** | Specify the desired state of this orchestration: &lt;code&gt;active&lt;/code&gt;, &lt;code&gt;inactive&lt;/code&gt;, or &lt;code&gt;suspend&lt;/code&gt;. Deletes only the non-persistent objects; persistent objects are not deleted when you change the state of an active orchestration to suspend. | [default to null]
**Name** | **string** | The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the orchestration. | [default to null]
**Objects** | [**[]OrchestrationObjectPostRequest**](OrchestrationObject-post-request.md) | The list of objects in the orchestration. An object is the primary building block of an orchestration. An orchestration can contain up to 100 objects. | [default to null]
**Tags** | **[]string** | Strings that describe the orchestration and help you identify it. | [optional] [default to null]
**Version** | **int32** | Specify the latest version of the orchestration in the request body. Before updating the orchestration, retrieve its details. This ensures that you only update the latest version of the orchestration. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


