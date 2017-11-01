# OrchestrationPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of this orchestration plan. | [optional] [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] [default to null]
**Oplans** | [**[]OPlanPostRequest**](OPlan-post-request.md) | List of oplans. An object plan, or oplan, is a top-level orchestration attribute. | [default to null]
**Relationships** | [**[]map[string]interface{}**](map.md) | A list of relationship specifications to be satisfied on this orchestration. For more information, see &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;cloud&amp;id&#x3D;STCSG-GUID-1896C799-49A6-42B8-9813-7DE5695267FE__RELATIONSHIPS-58824D2E\&quot;&gt;Relationships Between Object Plans&lt;/a&gt; in &lt;em&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/em&gt; | [optional] [default to null]
**Schedule** | [**interface{}**](interface{}.md) | The schedule for an orchestration consists of the start and stop dates and times. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


