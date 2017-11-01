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

type IpReservationResponse struct {

	// Shows the default account for your identity domain.
	Account string `json:"account,omitempty"`

	// Public IP address.<p>An IP reservation is a public IP address that you can attach to an Oracle Compute Cloud Service instance that requires access to or from the Internet.
	Ip string `json:"ip,omitempty"`

	// <p>The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).
	Name string `json:"name,omitempty"`

	// <code>/oracle/public/ippool</code><p>Pool of public IP addresses
	Parentpool string `json:"parentpool,omitempty"`

	// <code>true</code> indicates that the IP reservation has a persistent public IP address. You can associate either a temporary or a persistent public IP address with an instance when you create the instance.<p>Temporary public IP addresses are assigned dynamically from a pool of public IP addresses. When you associate a temporary public IP address with an instance, if the instance is restarted or is deleted and created again later, its public IP address might change.
	Permanent bool `json:"permanent,omitempty"`

	// Not used
	Quota string `json:"quota,omitempty"`

	// A comma-separated list of strings which helps you to identify IP reservation.
	Tags []string `json:"tags,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// <code>true</code> indicates that the IP reservation is associated with an instance.
	Used bool `json:"used,omitempty"`
}