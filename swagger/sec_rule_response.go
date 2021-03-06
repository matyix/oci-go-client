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

type SecRuleResponse struct {

	// <p>Set this parameter to <code>PERMIT</code>.
	Action string `json:"action,omitempty"`

	// <p>The three-part name of the security application: (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object_name</em></code>) for user-defined security applications and <code>/oracle/public/<em>object_name</em></code> for predefined security applications.
	Application string `json:"application,omitempty"`

	// <p>A description of the security rule.
	Description string `json:"description,omitempty"`

	// <p>Indicates whether the security rule is enabled (set to <code>false</code>) or disabled (<code>true</code>). The default setting is <code>false</code>.
	Disabled bool `json:"disabled,omitempty"`

	// <p>The three-part name (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object_name</em></code>) of the destination security list or security IP list.<p>You must use the prefix <code>seclist</code>: or <code>seciplist</code>: to identify the list type.<p><b>Note:</b> You can specify a security IP list as the destination in a secrule, provided <code>src_list</code> is a security list that has DENY as its outbound policy.<p>You cannot specify any of the security IP lists in the <code>/oracle/public</code> container as a destination in a secrule.
	DstList string `json:"dst_list,omitempty"`

	// The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).
	Name string `json:"name,omitempty"`

	// <p>The three-part name (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object_name</em></code>) of the source security list or security IP list.<p>You must use the prefix <code>seclist</code>: or <code>seciplist</code>: to identify the list type.
	SrcList string `json:"src_list,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`
}
