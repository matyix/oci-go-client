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

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type SecIPListsApi struct {
	Configuration *Configuration
}

func NewSecIPListsApi() *SecIPListsApi {
	configuration := NewConfiguration()
	return &SecIPListsApi{
		Configuration: configuration,
	}
}

func NewSecIPListsApiWithBasePath(basePath string) *SecIPListsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &SecIPListsApi{
		Configuration: configuration,
	}
}

/**
 * Create a Security IP List
 * Creates a security IP list. Note that, after creating a security IP list, you can add additional IP addresses to the list by using the POST /seciplist/ HTTP request again with just the additional IP addresses.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param body 
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecIpListResponse
 */
func (a SecIPListsApi) AddSecIPList(body SecIpListPostRequest, cookie string) (*SecIpListResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/seciplist/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Cookie"
	localVarHeaderParams["Cookie"] = a.Configuration.APIClient.ParameterToString(cookie, "")
	// body params
	localVarPostBody = &body
	var successPayload = new(SecIpListResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "AddSecIPList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Delete a Security IP List
 * Deletes the specified security IP list. No response is returned.&lt;p&gt;You can&#39;t delete system-provided security application that are available in the &lt;code&gt;/oracle/public&lt;/code&gt; container.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return void
 */
func (a SecIPListsApi) DeleteSecIPList(name string, cookie string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/seciplist/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Cookie"
	localVarHeaderParams["Cookie"] = a.Configuration.APIClient.ParameterToString(cookie, "")
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteSecIPList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * Retrieve Names of Containers
 * Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecIpListDiscoverResponse
 */
func (a SecIPListsApi) DiscoverRootSecIPList(cookie string) (*SecIpListDiscoverResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/seciplist/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+directory+json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Cookie"
	localVarHeaderParams["Cookie"] = a.Configuration.APIClient.ParameterToString(cookie, "")
	var successPayload = new(SecIpListDiscoverResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DiscoverRootSecIPList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Retrieve Names of all Security IP Lists and Subcontainers in a Container
 * Retrieves the names of objects and subcontainers that you can access in the specified container.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param container Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. Specify &lt;code&gt;/oracle/public&lt;/code&gt; to retrieve the names of system-provided objects.
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecIpListDiscoverResponse
 */
func (a SecIPListsApi) DiscoverSecIPList(container string, cookie string) (*SecIpListDiscoverResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/seciplist/{container}"
	localVarPath = strings.Replace(localVarPath, "{"+"container"+"}", fmt.Sprintf("%v", container), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+directory+json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Cookie"
	localVarHeaderParams["Cookie"] = a.Configuration.APIClient.ParameterToString(cookie, "")
	var successPayload = new(SecIpListDiscoverResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DiscoverSecIPList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Retrieve Details of a Security IP List
 * Retrieves information about the specified security IP list. You can use this request to verify whether POST and PUT operations were completed successfully.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt; or &lt;code&gt;/oracle/public/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;.)
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecIpListResponse
 */
func (a SecIPListsApi) GetSecIPList(name string, cookie string) (*SecIpListResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/seciplist/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Cookie"
	localVarHeaderParams["Cookie"] = a.Configuration.APIClient.ParameterToString(cookie, "")
	var successPayload = new(SecIpListResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetSecIPList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Retrieve Details of all Security IP Lists in a Container
 * Retrieves details of the security IP lists that are in the specified container and match the specified query criteria. If you don&#39;t specify any query criteria, then details of all the security IP lists in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:&lt;p&gt;&lt;code&gt;?parameter1&#x3D;value1&amp;ampparameter2&#x3D;value2&amp;ampparameterN&#x3D;valueN&lt;/code&gt;&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param container &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; and &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;/code&gt;for user-defined security IP lists and &lt;code&gt;/oracle/public&lt;/code&gt; for the predefined public-internet security IP list.&lt;/p&gt;
 * @param name &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt; or &lt;code&gt;/oracle/public/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;.)
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecIpListListResponse
 */
func (a SecIPListsApi) ListSecIPList(container string, name string, cookie string) (*SecIpListListResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/seciplist/{container}/"
	localVarPath = strings.Replace(localVarPath, "{"+"container"+"}", fmt.Sprintf("%v", container), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("name", a.Configuration.APIClient.ParameterToString(name, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Cookie"
	localVarHeaderParams["Cookie"] = a.Configuration.APIClient.ParameterToString(cookie, "")
	var successPayload = new(SecIpListListResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListSecIPList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Update a Security IP List
 * Updates IP addresses and description of the specified security IP list. Note that this command replaces the values in the &lt;code&gt;secipentries&lt;/code&gt; and &lt;code&gt;description&lt;/code&gt; fields with the new values that you specify. To add one or more IP addresses to the existing list, run the add seciplist command and specify just the additional IP addresses.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param body 
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecIpListResponse
 */
func (a SecIPListsApi) UpdateSecIPList(name string, body SecIpListPutRequest, cookie string) (*SecIpListResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/seciplist/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Cookie"
	localVarHeaderParams["Cookie"] = a.Configuration.APIClient.ParameterToString(cookie, "")
	// body params
	localVarPostBody = &body
	var successPayload = new(SecIpListResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateSecIPList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

