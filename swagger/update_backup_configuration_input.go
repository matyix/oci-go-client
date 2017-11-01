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

type UpdateBackupConfigurationInput struct {

	// The complete URI of the storage volume that you want to backup.
	VolumeUri string `json:"volumeUri"`

	// Any actions on this model will be performed as this user.
	RunAsUser string `json:"runAsUser"`

	// Description of this Backup Configuration
	Description string `json:"description,omitempty"`

	// Backup scheduler interval
	Interval Interval `json:"interval"`

	// When true, backups will automatically be generated based on the interval.
	Enabled bool `json:"enabled,omitempty"`

	// How many backups to retain
	BackupRetentionCount int32 `json:"backupRetentionCount,omitempty"`

	// The multi-part name of the object. Object names can contain only alphanumeric characters, hyphens, underscores, @ and periods. Object names are case-sensitive.
	Name string `json:"name"`
}
