/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UsersFollowingCreateResponse struct {
	Data *UsersFollowingCreateResponseData `json:"data,omitempty"`
	Errors []Problem `json:"errors,omitempty"`
}