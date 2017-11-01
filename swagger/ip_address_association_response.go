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

type IpAddressAssociationResponse struct {

	// Description of the object.
	Description string `json:"description,omitempty"`

	// The three-part name of the NAT IP address reservation (/Compute-identity_domain/user/object).
	IpAddressReservation string `json:"ipAddressReservation,omitempty"`

	// The three-part name of the NAT IP address association (/Compute-identity_domain/user/object).
	Name string `json:"name,omitempty"`

	// Tags associated with the object.
	Tags []string `json:"tags,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// Multipart name of the virtual NIC associated with this NAT IP reservation.
	Vnic string `json:"vnic,omitempty"`
}
