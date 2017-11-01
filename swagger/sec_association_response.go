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

type SecAssociationResponse struct {

	// The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).
	Name string `json:"name,omitempty"`

	// <p>Security list that you want to associate with the instance.
	Seclist string `json:"seclist,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// <p>vcable of the instance that you want to associate with the security list.<p>For more information about the vcable of an instance, see <a class=\"xref\" href=\"op-instance-{name}-get.html\">Retrieve Details of an Instance.</a>
	Vcable string `json:"vcable,omitempty"`
}
