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

type VpnEndpointResponse struct {

	// IP address of the VPN gateway in your data center through which you want to connect to the Oracle Cloud VPN gateway.
	CustomerVpnGateway string `json:"customer_vpn_gateway,omitempty"`

	// <code>true</code> indicates that the VPN endpoint is enabled and a connection is established immediately, if possible.
	Enabled bool `json:"enabled,omitempty"`

	// Two-part name of the object (<code><em>/Compute-acme/object</em></code>)
	Name string `json:"name,omitempty"`

	// Pre-shared VPN key.
	Psk string `json:"psk,omitempty"`

	// List of subnets (CIDR prefixes) that are reachable through this VPN tunnel.
	ReachableRoutes []string `json:"reachable_routes,omitempty"`

	// Current status of the VPN tunnel.
	Status string `json:"status,omitempty"`

	// Describes the current status of the VPN tunnel.
	StatusDesc string `json:"status_desc,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// Statistics of VPN tunnels
	VpnStatistics map[string]interface{} `json:"vpn_statistics,omitempty"`
}
