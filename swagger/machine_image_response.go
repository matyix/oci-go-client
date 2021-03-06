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

type MachineImageResponse struct {

	// Two-part name of the account (<code>/Compute-<em>identity_domain</em>/cloud_storage</code> that contains the credentials and access details of the associated Oracle Storage Cloud Service instance.
	Account string `json:"account,omitempty"`

	// <p>An optional JSON object or dictionary of arbitrary attributes to be made available to the instance. These are user-defined tags. After defining attributes, you can view them from within an instance at http://192.0.0.192/. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=cloud&id=STCSG-GUID-79AE6910-E6E3-4960-A9A7-368E26D894F6\">Retrieving User-Defined Instance Attributes</a> in <em>Using Oracle Compute Cloud Service (IaaS)</em>.
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// Last time when this image was audited
	Audited string `json:"audited,omitempty"`

	// Not used
	Checksums map[string]interface{} `json:"checksums,omitempty"`

	// Arbitrary text describing the image.
	Description string `json:"description,omitempty"`

	// Description of the state of the machine image if there is an error.
	ErrorReason string `json:"error_reason,omitempty"`

	// Name of the machine image file.
	File string `json:"file,omitempty"`

	// A dictionary of hypervisor-specific attributes.
	Hypervisor map[string]interface{} `json:"hypervisor,omitempty"`

	// The format of the image.
	ImageFormat string `json:"image_format,omitempty"`

	// <p>The three-part name of the object.
	Name string `json:"name,omitempty"`

	// <code>true</code> indicates that the image file is available in Oracle Cloud Storage Service.
	NoUpload bool `json:"no_upload,omitempty"`

	// The OS platform of the image.
	Platform string `json:"platform,omitempty"`

	// Not used
	Quota string `json:"quota,omitempty"`

	// Not used
	Signature string `json:"signature,omitempty"`

	// Not used
	SignedBy string `json:"signed_by,omitempty"`

	// Size values of the image file.
	Sizes interface{} `json:"sizes,omitempty"`

	// The state of the uploaded machine image.
	State string `json:"state,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`
}
