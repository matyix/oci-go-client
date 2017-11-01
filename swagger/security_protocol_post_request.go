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

// The request body contains details of the security protocol that you want to create. 
type SecurityProtocolPostRequest struct {

	// Description of the security protocol.
	Description string `json:"description,omitempty"`

	// Enter a list of port numbers or port range strings. Traffic is enabled by a security rule when a packet's destination port matches the ports specified here.<p>For TCP, SCTP, and UDP, each port is a destination transport port, between 0 and 65535, inclusive. For ICMP, each port is an ICMP type, between 0 and 255, inclusive.<p>If no destination ports are specified, all destination ports or ICMP types are allowed.
	DstPortSet []string `json:"dstPortSet,omitempty"`

	// The protocol used in the data portion of the IP datagram. Specify one of the permitted values or enter a number in the range 0 - 254 to represent the protocol that you want to specify. See <a target=\"_blank\" href=\"http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml\">Assigned Internet Protocol Numbers</a>. Permitted values are: <code>tcp</code>, <code>udp</code>, <code>icmp</code>, <code>igmp</code>, <code>ipip</code>, <code>rdp</code>, <code>esp</code>, <code>ah</code>, <code>gre</code>, <code>icmpv6</code>, <code>ospf</code>, <code>pim</code>, <code>sctp</code>, <code>mplsip</code>, <code>all</code>.<p>Traffic is enabled by a security rule when the protocol in the packet matches the protocol specified here. If no protocol is specified, all protocols are allowed.
	IpProtocol string `json:"ipProtocol,omitempty"`

	// The three-part name of the Ip address association (/Compute-identity_domain/user/object).<p>Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. When you specify the object name, ensure that an object of the same type and with the same name doesn't already exist. If such an object already exists, another object of the same type and with the same name won't be created and the existing object won't be updated.
	Name string `json:"name"`

	// Enter a list of port numbers or port range strings. Traffic is enabled by a security rule when a packet's source port matches the ports specified here.<p>For TCP, SCTP, and UDP, each port is a source transport port, between 0 and 65535, inclusive. For ICMP, each port is an ICMP type, between 0 and 255, inclusive.<p>If no source ports are specified, all source ports or ICMP types are allowed.
	SrcPortSet []string `json:"srcPortSet,omitempty"`

	// Strings that you can use to tag the security protocol.
	Tags []string `json:"tags,omitempty"`
}