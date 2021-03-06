/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A problem that indicates that you are not allowed to see a particular field on a Tweet, User, etc.
type FieldUnauthorizedProblem struct {
	Type_ string `json:"type"`
	Title string `json:"title"`
	Detail string `json:"detail,omitempty"`
	Status int32 `json:"status,omitempty"`
	Section string `json:"section"`
	ResourceType string `json:"resource_type"`
	Field string `json:"field"`
}
