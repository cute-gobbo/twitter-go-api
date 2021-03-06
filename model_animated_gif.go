/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AnimatedGif struct {
	Type_ string `json:"type"`
	MediaKey string `json:"media_key,omitempty"`
	Height int32 `json:"height,omitempty"`
	Width int32 `json:"width,omitempty"`
	PreviewImageUrl string `json:"preview_image_url,omitempty"`
	// An array of all available variants of the media
	Variants []VideoVariants `json:"variants,omitempty"`
}
