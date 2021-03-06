/* 
 * REST API for Oracle Cloud Infrastructure Compute Classic
 *
 * Use the Oracle Cloud Infrastructure Compute Classic REST API to provision and manage instances and the associated resources. <p>All documentation is applicable to using the API on Oracle Cloud and Oracle Cloud Machine, unless otherwise indicated.
 *
 * OpenAPI spec version: 2017.09.21
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// The request body contains details of the orchestration object that you want to update. 
type OrchestrationObjectPutRequest struct {

	// <code>/Compute-<em>identity_domain</em>/default</code>
	Account string `json:"account,omitempty"`

	//  A text string describing the object.
	Description string `json:"description,omitempty"`

	// The desired state of this object
	DesiredState string `json:"desired_state,omitempty"`

	// A text string describing the object. Labels can't include spaces. In an orchestration, the label for each object must be unique. Maximum length is 256 characters.
	Label string `json:"label"`

	//  The four-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>orchestration</em>/<em>object</em></code>). If you don't specify a name for this object, the name is generated automatically. Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. When you specify the object name, ensure that an object of the same type and with the same name doesn't already exist. If such a object already exists, then another object of the same type and with the same name won't be created and the existing object won't be updated.
	Name string `json:"name,omitempty"`

	// The three-part name (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>) of the orchestration to which the object belongs.
	Orchestration string `json:"orchestration"`

	// Specifies whether the object should persist when the orchestration is suspended. Specify one of the following:<p>* <code>true</code>: The object persists when the orchestration is suspended.<p>*  <code>false</code>: The object is deleted when the orchestration is suspended.<p>By default, <code>persistent</code> is set to <code>false</code>. It is recommended that you specify true for storage volumes and other critical objects. Persistence applies only when you're suspending an orchestration. When you terminate an orchestration, all the objects defined in it are deleted.
	Persistent bool `json:"persistent,omitempty"`

	// The relationship between the objects that are created by this orchestration. The only supported relationship is depends, indicating that the specified target objects must be created first.<p>Note that when recovering from a failure, the orchestration doesn't consider object relationships. Orchestrations v2 use object references to recover interdependent objects to a healthy state. See<a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-B3732A3E-2E16-4643-9C28-648F3BB09FE9\">Object References and Relationships</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i>.
	Relationships []map[string]interface{} `json:"relationships,omitempty"`

	//  The template attribute defines the properties or characteristics of the Oracle Compute Cloud Service object that you want to create, as specified by the <code>type</code> attribute.<p>The fields in the template section vary depending on the specified type. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-B5D59289-8892-4C4C-8DA6-33ACE963E235\">Orchestration v2 Attributes Specific to Each Object Type</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i> to determine the parameters that are specific to each object <code>type</code> that you want to create.<p>For example, if you want to create a storage volume, the <code>type</code> would be <code>StorageVolume</code>, and the <code>template</code> would include <code>size</code> and <code>bootable</code>. If you want to create an instance, the type would be <code>Instance</code>, and the template would include instance-specific attributes, such as <code>imagelist</code> and <code>shape</code>.
	Template map[string]interface{} `json:"template"`

	// Specify one of the following object types that you want to create.<p>* Acl <p>* Backup <p>* BackupConfiguration <p>* Instance <p>* IpAddressAssociation <p>* IpAddressPrefixSet <p>* IpAddressReservation <p>* IpNetwork <p>* IpNetworkExchange <p>* IPReservation <p>* OSSContainer <p>* Restore <p>* Route <p>* SecApplication <p>* SecIPList <p>* SecList <p>* SecRule <p>* SecurityProtocol <p>* SecurityRule <p>* SSHKey <p>* StorageAttachment <p>* StorageSnapshot <p>* StorageSnapshot <p>* StorageVolume <p>* VirtualNicSet <p>For the most up-to-date information about the supported object types, see <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-2D49EB53-324B-4B5F-B7F7-61E64D62D235\">Object Types in Orchestgations v2</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i>.
	Type_ string `json:"type"`

	// Specify the latest version of the orchestration object in the request body. Before updating the orchestration object, retrieve its details. This ensures that you only update the latest version of the orchestration object.
	Version int32 `json:"version"`
}
