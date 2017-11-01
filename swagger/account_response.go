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

type AccountResponse struct {

	// Account Type for this account
	Accounttype string `json:"accounttype,omitempty"`

	// Credentials specific to the account, which may include a username, password or certificate. The credentials are not returned.
	Credentials map[string]interface{} `json:"credentials,omitempty"`

	// Description of this account.
	Description string `json:"description,omitempty"`

	// Two-part name of the account: <code>/Compute-<em>identity_domain</em>/default</code> or <code>/Compute-<em>identity_domain</em>/cloud_storage</code>
	Name string `json:"name,omitempty"`

	// Properties to be applied to specific objects created using this account
	Objectproperties map[string]interface{} `json:"objectproperties,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`
}