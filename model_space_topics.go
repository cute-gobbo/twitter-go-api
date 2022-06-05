/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The Twitter Topic object
type SpaceTopics struct {
	// An ID suitable for use in the REST API.
	Id string `json:"id"`
	// The description of the given topic.
	Description string `json:"description,omitempty"`
	// The name of the given topic.
	Name string `json:"name"`
}