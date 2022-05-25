/*
Tweets and Users

API Reference — Labs v2

API version: 2.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package twitterapi

import (
	"encoding/json"
)

// TweetReferencedTweets struct for TweetReferencedTweets
type TweetReferencedTweets struct {
	Type string `json:"type"`
	// Unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers.
	Id string `json:"id"`
}

// NewTweetReferencedTweets instantiates a new TweetReferencedTweets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetReferencedTweets(type_ string, id string) *TweetReferencedTweets {
	this := TweetReferencedTweets{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewTweetReferencedTweetsWithDefaults instantiates a new TweetReferencedTweets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetReferencedTweetsWithDefaults() *TweetReferencedTweets {
	this := TweetReferencedTweets{}
	return &this
}

// GetType returns the Type field value
func (o *TweetReferencedTweets) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TweetReferencedTweets) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TweetReferencedTweets) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TweetReferencedTweets) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TweetReferencedTweets) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TweetReferencedTweets) SetId(v string) {
	o.Id = v
}

func (o TweetReferencedTweets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableTweetReferencedTweets struct {
	value *TweetReferencedTweets
	isSet bool
}

func (v NullableTweetReferencedTweets) Get() *TweetReferencedTweets {
	return v.value
}

func (v *NullableTweetReferencedTweets) Set(val *TweetReferencedTweets) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetReferencedTweets) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetReferencedTweets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetReferencedTweets(val *TweetReferencedTweets) *NullableTweetReferencedTweets {
	return &NullableTweetReferencedTweets{value: val, isSet: true}
}

func (v NullableTweetReferencedTweets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetReferencedTweets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

