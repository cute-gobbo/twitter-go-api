/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A problem that indicates you are not allowed to see a particular Tweet, User, etc.
type ResourceUnauthorizedProblem struct {
	Type_ string `json:"type"`
	Title string `json:"title"`
	Detail string `json:"detail,omitempty"`
	Status int32 `json:"status,omitempty"`
	Value string `json:"value"`
	Parameter string `json:"parameter"`
	Section string `json:"section"`
	ResourceId string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}
