# VirtualNicSetPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAcls** | **[]string** | List of ACLs applied to the VNICs in the set. | [optional] [default to null]
**Description** | **string** | Description of the object. | [optional] [default to null]
**Name** | **string** | The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the virtual NIC set. Object names can contain only alphanumeric, underscore (_), dash (-), and period (.) characters. Object names are case-sensitive. | [default to null]
**Tags** | **[]string** | Tags associated with the object. | [optional] [default to null]
**Vnics** | **[]string** | List of VNICs associated with this VNIC set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


