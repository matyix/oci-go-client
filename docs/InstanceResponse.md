# InstanceResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Shows the default account for your identity domain. | [optional] [default to null]
**Attributes** | [**map[string]interface{}**](interface{}.md) | A dictionary of attributes to be made available to the instance. A value with the key \&quot;userdata\&quot; will be made available in an EC2-compatible manner. | [optional] [default to null]
**AvailabilityDomain** | **string** | The availability domain the instance is in | [optional] [default to null]
**BootOrder** | **[]int32** | Boot order list. | [optional] [default to null]
**DesiredState** | **string** | Desired state for the instance. The value can be &lt;code&gt;shutdown&lt;/code&gt; or &lt;code&gt;running&lt;/code&gt; to shutdown an instance or to restart a previously shutdown instance respectively. | [optional] [default to null]
**DiskAttach** | **string** | A label assigned by the user to identify disks. | [optional] [default to null]
**Domain** | **string** | The default domain to use for the hostname and for DNS lookups. | [optional] [default to null]
**Entry** | **int32** | Optional imagelistentry number (default will be used if not specified). | [optional] [default to null]
**ErrorReason** | **string** | The reason for the instance going to error state, if available. | [optional] [default to null]
**Fingerprint** | **string** | SSH server fingerprint presented by the instance. | [optional] [default to null]
**Hostname** | **string** | The hostname for this instance. | [optional] [default to null]
**Hypervisor** | [**map[string]interface{}**](interface{}.md) | A dictionary of hypervisor-specific attributes. | [optional] [default to null]
**ImageFormat** | **string** | The format of the image. | [optional] [default to null]
**Imagelist** | **string** | Name of imagelist to be launched. | [optional] [default to null]
**Ip** | **string** | IP address of the instance. | [optional] [default to null]
**Label** | **string** | A label assigned by the user, specifically for defining inter-instance relationships. | [optional] [default to null]
**Name** | **string** | Multipart name of the instance. | [optional] [default to null]
**Networking** | [**map[string]interface{}**](interface{}.md) | Mapping of &lt;device name&gt; to network specifiers for virtual NICs to be attached to this instance. | [optional] [default to null]
**PlacementRequirements** | **[]string** | A list of strings specifying arbitrary tags on nodes to be matched on placement. | [optional] [default to null]
**Platform** | **string** | The OS platform for the instance. | [optional] [default to null]
**Priority** | **string** | The priority at which this instance will be run. | [optional] [default to null]
**Quota** | **string** | Not used | [optional] [default to null]
**QuotaReservation** | **string** | Reference to the QuotaReservation, to be destroyed with the instance. | [optional] [default to null]
**Relationships** | [**[]map[string]interface{}**](map.md) | A list of relationship specifications to be satisfied on this instance&#39;s placement | [optional] [default to null]
**Resolvers** | **[]string** | Resolvers to use instead of the default resolvers. | [optional] [default to null]
**ReverseDns** | **bool** | Add PTR records for the hostname. | [optional] [default to null]
**Shape** | **string** | A shape is a resource profile that specifies the number of CPU threads and the amount of memory (in MB) to be allocated to an instance. | [optional] [default to null]
**Site** | **string** | Site to run on. | [optional] [default to null]
**Sshkeys** | **[]string** | SSH keys that will be exposed to the instance. | [optional] [default to null]
**StartTime** | **string** | Start time of the instance. | [optional] [default to null]
**State** | **string** | State of the instance. | [optional] [default to null]
**StorageAttachments** | **[]string** | List of dictionaries containing storage attachment Information. | [optional] [default to null]
**Tags** | **[]string** | Comma-separated list of strings used to tag the instance. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]
**VcableId** | **string** | vCable for this instance. | [optional] [default to null]
**Virtio** | **bool** | Specify if the devices created for the instance are virtio devices. If not specified, the default will come from the cluster configration file. | [optional] [default to null]
**Vnc** | **string** | IP address and port of the VNC console for the instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


