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

type OPlanResponse struct {

	// 
	HaPolicy string `json:"ha_policy,omitempty"`

	// Info dictionary for the oplan.
	Info map[string]interface{} `json:"info,omitempty"`

	// Description of this object plan.
	Label string `json:"label,omitempty"`

	// Type of the object.
	ObjType string `json:"obj_type,omitempty"`

	// List of object dictionaries or object names.
	Objects []string `json:"objects,omitempty"`

	// Most recent status.
	Status string `json:"status,omitempty"`

	// Timestamp of the most-recent status change.
	StatusTimestamp string `json:"status_timestamp,omitempty"`
}
