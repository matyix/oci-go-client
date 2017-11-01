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

// The request body contains details of the instance that you want to update. 
type InstancePutRequest struct {

	// Desired state for the instance.<p>* Set the value of the <code>desired_state</code> parameter as <code>shutdown</code> to shut down the instance.<p>* Set the value of the <code>desired_state</code> parameter as <code>running</code> to restart an instance that you had previously shutdown. The instance is restarted without losing any of the instance data or configuration.
	DesiredState string `json:"desired_state,omitempty"`

	// A shape is a resource profile that specifies the number of CPU threads and the amount of memory (in MB) to be allocated to an instance.<p>If the instance is already shut down, you can change the shape that is associated with the instance.
	Shape string `json:"shape,omitempty"`

	// Comma-separated list of strings that you can use to tag the instance. Later on you can use these tags to retrieve the instance.
	Tags []string `json:"tags,omitempty"`
}
