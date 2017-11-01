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

type OrchestrationResponse struct {

	// Shows the default account for your identity domain.
	Account string `json:"account,omitempty"`

	// Description of this orchestration plan.
	Description string `json:"description,omitempty"`

	// The nested parameter <code>errors</code> shows which object in the orchestration has encountered an error. Empty if there are no errors.
	Info map[string]interface{} `json:"info,omitempty"`

	// <p>The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).
	Name string `json:"name,omitempty"`

	// List of oplans. An object plan, or oplan, is a top-level orchestration attribute.
	Oplans []OPlanResponse `json:"oplans,omitempty"`

	// The relationship between the objects that are created by this orchestration.
	Relationships []map[string]interface{} `json:"relationships,omitempty"`

	// The schedule for an orchestration consists of the start and stop dates and times.
	Schedule interface{} `json:"schedule,omitempty"`

	// Shows the current status of the orchestration.
	Status string `json:"status,omitempty"`

	// This information is generally displayed at the end of the orchestration JSON. It indicates the time that the current view of the orchestration was generated. This information shows only when the orchestration is running.
	StatusTimestamp string `json:"status_timestamp,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// Two-part name of the user (<code>/Compute-<em>identity_domain</em>/<em>user</em></code>) who has most recently taken an action on the orchestration.
	User string `json:"user,omitempty"`
}