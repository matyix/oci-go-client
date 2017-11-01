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

// The request body contains details of the virtual NIC set that you want to create. 
type VirtualNicSetPostRequest struct {

	// List of ACLs applied to the VNICs in the set.
	AppliedAcls []string `json:"appliedAcls,omitempty"`

	// Description of the object.
	Description string `json:"description,omitempty"`

	// The three-part name (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>) of the virtual NIC set. Object names can contain only alphanumeric, underscore (_), dash (-), and period (.) characters. Object names are case-sensitive.
	Name string `json:"name"`

	// Tags associated with the object.
	Tags []string `json:"tags,omitempty"`

	// List of VNICs associated with this VNIC set.
	Vnics []string `json:"vnics,omitempty"`
}