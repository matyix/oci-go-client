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

type IpAddressReservationResponse struct {

	// Description of the object.
	Description string `json:"description,omitempty"`

	// Reserved NAT IPv4 address from the IP address pool.
	IpAddress string `json:"ipAddress,omitempty"`

	// Name of the IP address pool to reserve the NAT IP from.
	IpAddressPool string `json:"ipAddressPool,omitempty"`

	// Name of the IP address reservation.
	Name string `json:"name,omitempty"`

	// Tags associated with the object.
	Tags []string `json:"tags,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`
}