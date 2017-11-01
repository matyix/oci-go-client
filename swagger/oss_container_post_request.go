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

// The request body contains details of the Oracle Storage Cloud Service container that you want to create. 
type OssContainerPostRequest struct {

	// The two-part name of the account (/Compute-identity_domain/cloud_storage) that contains the credentials and access details of the associated Oracle Storage Cloud Service instance.
	Account string `json:"account"`

	// The name of the container that you want to create.<p>Container names must:<p>* Contain only UTF-8 characters<p>* Be a maximum of 256 bytes<p>* Avoid using a slash (/) character because this character acts as a delimiter between the container name and the object name<p>When you specify a container name, ensure that a container of the same name doesn't already exist.
	Container string `json:"container,omitempty"`

	// Determines whether the contents of the Oracle Storage Cloud Service container is also deleted when you delete the <code>integrations/osscontainer</code> object.<p>When set to <code>true</code>, deletes the Oracle Storage Cloud Service container along with all the objects in the container when you delete the integration/osscontainer object created by this orchestration.<p>When set to <code>false</code>, only the integrations/osscontainer object created by this orchestration is deleted.
	DeleteRemote bool `json:"delete_remote"`

	// The three-part name of the integrations/osscontainer object created by this orchestration. This name is in the format /Compute-identity_domain/user/object.<p>If you don't specify a name for this object, then the name is generated automatically.<p>Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive.<p>When you specify the object name, ensure that an object of the same type and with the same name doesn't already exist. If such an object already exists, another object of the same type and with the same name won't be created and the existing object won't be updated.
	Name string `json:"name,omitempty"`
}