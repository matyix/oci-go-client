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

// The request body contains details of the security rule that you want to create. 
type SecurityRulePostRequest struct {

	// Select the three-part name of the access control list (ACL) that you want to add this security rule to. Security rules are applied to vNIC sets by using ACLs.
	Acl string `json:"acl,omitempty"`

	// Description of the security rule.
	Description string `json:"description,omitempty"`

	// A list of IP address prefix sets to which you want to permit traffic. Only packets to IP addresses in the specified IP address prefix sets are permitted. When no destination IP address prefix sets are specified, traffic to any IP address is permitted.
	DstIpAddressPrefixSets []string `json:"dstIpAddressPrefixSets,omitempty"`

	// The vNICset to which you want to permit traffic. Only packets to vNICs in the specified vNICset are permitted. When no destination vNICset is specified, traffic to any vNIC is permitted.
	DstVnicSet string `json:"dstVnicSet,omitempty"`

	// Allows the security rule to be enabled or disabled. This parameter is set to true by default. Specify false to disable the security rule.
	EnabledFlag bool `json:"enabledFlag,omitempty"`

	//  Specify the direction of flow of traffic, which is relative to the instances, for this security rule. Allowed values are <code>ingress</code> or <code>egress</code>.<p>An ingress packet is a packet received by a virtual NIC, for example from another virtual NIC or from the public Internet.<p>An egress packet is a packet sent by a virtual NIC, for example to another virtual NIC or to the public Internet.
	FlowDirection string `json:"flowDirection"`

	// The three-part name of the Ip address association (/Compute-identity_domain/user/object).<p>Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. When you specify the object name, ensure that an object of the same type and with the same name doesn't already exist. If such an object already exists, another object of the same type and with the same name won't be created and the existing object won't be updated.
	Name string `json:"name"`

	// A list of security protocols for which you want to permit traffic. Only packets that match the specified protocols and ports are permitted. When no security protocols are specified, traffic using any protocol over any port is permitted.
	SecProtocols []string `json:"secProtocols,omitempty"`

	// A list of IP address prefix sets from which you want to permit traffic. Only packets from IP addresses in the specified IP address prefix sets are permitted. When no source IP address prefix sets are specified, traffic from any IP address is permitted.
	SrcIpAddressPrefixSets []string `json:"srcIpAddressPrefixSets,omitempty"`

	// The vNICset from which you want to permit traffic. Only packets from vNICs in the specified vNICset are permitted. When no source vNICset is specified, traffic from any vNIC is permitted.
	SrcVnicSet string `json:"srcVnicSet,omitempty"`

	// Strings that you can use to tag the security rule.
	Tags []string `json:"tags,omitempty"`
}