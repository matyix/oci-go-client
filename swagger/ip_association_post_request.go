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

// The request body contains details of the IP association that you want to create. 
type IpAssociationPostRequest struct {

	// <ul><li>To associate a temporary IP address from the pool, specify ippool:/oracle/public/ippool.</li><li>To associate a persistent IP address, specify ipreservation:ipreservation_name, where ipreservation_name is three-part name of an existing IP reservation in the <code>/Compute-identity_domain/user/object_name</code> format. For more information about how to create an IP reservation, see <a class=\"xref\" href=\"op-ip-reservation--post.html\">Create an IP Reservation</a>.</li></ul>
	Parentpool string `json:"parentpool"`

	// <p>The three-part name of the vcable ID of the instance that you want to associate with an IP address. The three-part name is in the format: <code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>.<p>For more information about the vcable of an instance, see <a class=\"xref\" href=\"op-instance-{name}-get.html\">Retrieve Details of an Instance</a>.
	Vcable string `json:"vcable"`
}