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

type SecurityProtocolsApi struct {
	Configuration *Configuration
}

func NewSecurityProtocolsApi() *SecurityProtocolsApi {
	configuration := NewConfiguration()
	return &SecurityProtocolsApi{
		Configuration: configuration,
	}
}

func NewSecurityProtocolsApiWithBasePath(basePath string) *SecurityProtocolsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &SecurityProtocolsApi{
		Configuration: configuration,
	}
}

/**
 * Create a Security Protocol for IP Networks
 * Creates a security protocol. A security protocol allows you to specify a transport protocol and the source and destination ports to be used with the specified protocol. When you create a security rule, the protocols and ports of the specified security protocols are used to determine the type of traffic that is permitted by that security rule. If you don&#39;t specify protocols and ports in a security protocol, traffic is permitted over all protocols and ports.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param body 
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecurityProtocolResponse
 */
func (a SecurityProtocolsApi) AddSecurityProtocol(body SecurityProtocolPostRequest, cookie string) (*SecurityProtocolResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/network/v1/secprotocol/"

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
	var successPayload = new(SecurityProtocolResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "AddSecurityProtocol", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Delete a Security Protocol Used in IP Networks
 * Deletes a security protocol. Ensure that the security protocol is not being used before deleting it.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return void
 */
func (a SecurityProtocolsApi) DeleteSecurityProtocol(name string, cookie string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/network/v1/secprotocol/{name}"
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
	var localVarAPIResponse = &APIResponse{Operation: "DeleteSecurityProtocol", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Retrieve Details of a Security Protocol Used in IP Networks
 * Retrieves details of the specified security protocol.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecurityProtocolResponse
 */
func (a SecurityProtocolsApi) GetSecurityProtocol(name string, cookie string) (*SecurityProtocolResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/network/v1/secprotocol/{name}"
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
	var successPayload = new(SecurityProtocolResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetSecurityProtocol", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Retrieve Details of all Security Protocols in a Container
 * Retrieve details of all security protocols in the specified container. This request is for security protocols used in IP networks.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param container &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt;
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecurityProtocolListResponse
 */
func (a SecurityProtocolsApi) ListSecurityProtocol(container string, cookie string) (*SecurityProtocolListResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/network/v1/secprotocol/{container}/"
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
		"application/oracle-compute-v3+json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Cookie"
	localVarHeaderParams["Cookie"] = a.Configuration.APIClient.ParameterToString(cookie, "")
	var successPayload = new(SecurityProtocolListResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListSecurityProtocol", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Update a Security Protocol Used in IP Networks
 * Update an existing security protocol. You can update values of the &lt;code&gt;description&lt;/code&gt;, &lt;code&gt;ipProtocol&lt;/code&gt;, &lt;code&gt;srcPortSet&lt;/code&gt;, &lt;code&gt;dstPortSet&lt;/code&gt;, and &lt;code&gt;tags&lt;/code&gt; parameters of a security protocol.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param body 
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *SecurityProtocolResponse
 */
func (a SecurityProtocolsApi) UpdateSecurityProtocol(name string, body SecurityProtocolPutRequest, cookie string) (*SecurityProtocolResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/network/v1/secprotocol/{name}"
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
	var successPayload = new(SecurityProtocolResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateSecurityProtocol", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
