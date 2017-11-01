# InstancePutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredState** | **string** | Desired state for the instance.&lt;p&gt;* Set the value of the &lt;code&gt;desired_state&lt;/code&gt; parameter as &lt;code&gt;shutdown&lt;/code&gt; to shut down the instance.&lt;p&gt;* Set the value of the &lt;code&gt;desired_state&lt;/code&gt; parameter as &lt;code&gt;running&lt;/code&gt; to restart an instance that you had previously shutdown. The instance is restarted without losing any of the instance data or configuration. | [optional] [default to null]
**Shape** | **string** | A shape is a resource profile that specifies the number of CPU threads and the amount of memory (in MB) to be allocated to an instance.&lt;p&gt;If the instance is already shut down, you can change the shape that is associated with the instance. | [optional] [default to null]
**Tags** | **[]string** | Comma-separated list of strings that you can use to tag the instance. Later on you can use these tags to retrieve the instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


