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

type InstanceConsoleResponse struct {

	// Name of the instance for which the console output is displayed.
	Name string `json:"name,omitempty"`

	// Serial console output of the instance.
	Output string `json:"output,omitempty"`

	// Time when the console output was created.
	Timestamp string `json:"timestamp,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`
}
