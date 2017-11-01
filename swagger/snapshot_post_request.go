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

// The request body contains details of the instance snapshot that you want to create. 
type SnapshotPostRequest struct {

	// Two-part name of the account (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>) that contains the credentials and access details of Oracle Storage Cloud Service. The machine image file is uploaded to the Oracle Storage Cloud Service account that you specify.
	Account string `json:"account,omitempty"`

	// Use this option when you want to take preserve the custom changes you have made to an instance before deleting the instance. The only permitted value is shutdown. Snapshot of the instance is not taken immediately. It creates a machine image which preserves the changes you have made to the instance, and then the instance is deleted.<p><b>Note:</b> This option has no effect if you shutdown the instance from inside it. Any pending snapshot request on that instance goes into error state. You must delete the instance (<a class=\"xref\" href=\"op-instance-{name}-delete.html\">DELETE /instance/{name}</a>).
	Delay string `json:"delay,omitempty"`

	// Multipart name of the instance that you want to clone.
	Instance string `json:"instance"`

	// Specify the three-part name (<code>/Compute-identity_domain/user/object</code>) of the machine image created by the snapshot request.<p>Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive.<p>If you don't specify a name for this object, then the name is generated automatically.
	Machineimage string `json:"machineimage,omitempty"`
}
