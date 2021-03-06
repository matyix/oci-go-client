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

type SshKeyResponse struct {

	// Indicates whether the key is enabled (<code>true</code>) or disabled.
	Enabled bool `json:"enabled,omitempty"`

	// <p>The SSH public key value.
	Key string `json:"key,omitempty"`

	// <p>The three-part name of the object
	Name string `json:"name,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`
}
