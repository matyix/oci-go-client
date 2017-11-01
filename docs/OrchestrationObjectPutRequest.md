# OrchestrationObjectPutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/default&lt;/code&gt; | [optional] [default to null]
**Description** | **string** |  A text string describing the object. | [optional] [default to null]
**DesiredState** | **string** | The desired state of this object | [optional] [default to null]
**Label** | **string** | A text string describing the object. Labels can&#39;t include spaces. In an orchestration, the label for each object must be unique. Maximum length is 256 characters. | [default to null]
**Name** | **string** |  The four-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;orchestration&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). If you don&#39;t specify a name for this object, the name is generated automatically. Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. When you specify the object name, ensure that an object of the same type and with the same name doesn&#39;t already exist. If such a object already exists, then another object of the same type and with the same name won&#39;t be created and the existing object won&#39;t be updated. | [optional] [default to null]
**Orchestration** | **string** | The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the orchestration to which the object belongs. | [default to null]
**Persistent** | **bool** | Specifies whether the object should persist when the orchestration is suspended. Specify one of the following:&lt;p&gt;* &lt;code&gt;true&lt;/code&gt;: The object persists when the orchestration is suspended.&lt;p&gt;*  &lt;code&gt;false&lt;/code&gt;: The object is deleted when the orchestration is suspended.&lt;p&gt;By default, &lt;code&gt;persistent&lt;/code&gt; is set to &lt;code&gt;false&lt;/code&gt;. It is recommended that you specify true for storage volumes and other critical objects. Persistence applies only when you&#39;re suspending an orchestration. When you terminate an orchestration, all the objects defined in it are deleted. | [optional] [default to null]
**Relationships** | [**[]map[string]interface{}**](map.md) | The relationship between the objects that are created by this orchestration. The only supported relationship is depends, indicating that the specified target objects must be created first.&lt;p&gt;Note that when recovering from a failure, the orchestration doesn&#39;t consider object relationships. Orchestrations v2 use object references to recover interdependent objects to a healthy state. See&lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;STCSG-GUID-B3732A3E-2E16-4643-9C28-648F3BB09FE9\&quot;&gt;Object References and Relationships&lt;/a&gt; in &lt;i&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/i&gt;. | [optional] [default to null]
**Template** | [**map[string]interface{}**](interface{}.md) |  The template attribute defines the properties or characteristics of the Oracle Compute Cloud Service object that you want to create, as specified by the &lt;code&gt;type&lt;/code&gt; attribute.&lt;p&gt;The fields in the template section vary depending on the specified type. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;STCSG-GUID-B5D59289-8892-4C4C-8DA6-33ACE963E235\&quot;&gt;Orchestration v2 Attributes Specific to Each Object Type&lt;/a&gt; in &lt;i&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/i&gt; to determine the parameters that are specific to each object &lt;code&gt;type&lt;/code&gt; that you want to create.&lt;p&gt;For example, if you want to create a storage volume, the &lt;code&gt;type&lt;/code&gt; would be &lt;code&gt;StorageVolume&lt;/code&gt;, and the &lt;code&gt;template&lt;/code&gt; would include &lt;code&gt;size&lt;/code&gt; and &lt;code&gt;bootable&lt;/code&gt;. If you want to create an instance, the type would be &lt;code&gt;Instance&lt;/code&gt;, and the template would include instance-specific attributes, such as &lt;code&gt;imagelist&lt;/code&gt; and &lt;code&gt;shape&lt;/code&gt;. | [default to null]
**Type_** | **string** | Specify one of the following object types that you want to create.&lt;p&gt;* Acl &lt;p&gt;* Backup &lt;p&gt;* BackupConfiguration &lt;p&gt;* Instance &lt;p&gt;* IpAddressAssociation &lt;p&gt;* IpAddressPrefixSet &lt;p&gt;* IpAddressReservation &lt;p&gt;* IpNetwork &lt;p&gt;* IpNetworkExchange &lt;p&gt;* IPReservation &lt;p&gt;* OSSContainer &lt;p&gt;* Restore &lt;p&gt;* Route &lt;p&gt;* SecApplication &lt;p&gt;* SecIPList &lt;p&gt;* SecList &lt;p&gt;* SecRule &lt;p&gt;* SecurityProtocol &lt;p&gt;* SecurityRule &lt;p&gt;* SSHKey &lt;p&gt;* StorageAttachment &lt;p&gt;* StorageSnapshot &lt;p&gt;* StorageSnapshot &lt;p&gt;* StorageVolume &lt;p&gt;* VirtualNicSet &lt;p&gt;For the most up-to-date information about the supported object types, see &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;STCSG-GUID-2D49EB53-324B-4B5F-B7F7-61E64D62D235\&quot;&gt;Object Types in Orchestgations v2&lt;/a&gt; in &lt;i&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/i&gt;. | [default to null]
**Version** | **int32** | Specify the latest version of the orchestration object in the request body. Before updating the orchestration object, retrieve its details. This ensures that you only update the latest version of the orchestration object. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

