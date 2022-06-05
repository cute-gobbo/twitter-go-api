/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RulesResponseMetadata struct {
	Sent string `json:"sent"`
	Summary *RulesRequestSummary `json:"summary,omitempty"`
	// This parameter is used to get the next 'page' of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified.
	NextToken string `json:"next_token,omitempty"`
	// Number of Rules in result set
	ResultCount int32 `json:"result_count,omitempty"`
}