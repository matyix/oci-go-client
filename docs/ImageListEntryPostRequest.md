# ImageListEntryPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](interface{}.md) | &lt;p&gt;User-defined parameters, in JSON format, that can be passed to an instance of this machine image when it is launched. This field can be used, for example, to specify a user-friendly name for the image list in the displayName attribute. The name you specify is displayed as the name of the image list in the Oracle Compute Cloud Service web console.  Instance metadata, including user-defined data is available at http://192.0.0.192/ within an instance. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;cloud&amp;id&#x3D;STCSG-GUID-79AE6910-E6E3-4960-A9A7-368E26D894F6\&quot;&gt;Retrieving User-Defined Instance Attributes&lt;/a&gt; in &lt;em&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/em&gt;. | [optional] [default to null]
**Machineimages** | **[]string** | &lt;p&gt;A list of machine images. Specify the three-part name of each machine image. | [default to null]
**Version** | **int32** | The unique version of the entry in the image list. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


