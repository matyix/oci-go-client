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

type IPAssociationsApi struct {
	Configuration *Configuration
}

func NewIPAssociationsApi() *IPAssociationsApi {
	configuration := NewConfiguration()
	return &IPAssociationsApi{
		Configuration: configuration,
	}
}

func NewIPAssociationsApiWithBasePath(basePath string) *IPAssociationsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &IPAssociationsApi{
		Configuration: configuration,
	}
}

/**
 * Create an IP Association
 * Creates an association between an IP address and the vcable ID of an instance.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param body 
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *IpAssociationResponse
 */
func (a IPAssociationsApi) AddIPAssociation(body IpAssociationPostRequest, cookie string) (*IpAssociationResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ip/association/"

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
	var successPayload = new(IpAssociationResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "AddIPAssociation", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Delete an IP Association
 * Deletes the specified IP association. Use this HTTP request when you want to change the public IP address of an instance or if you no longer need a public IP address for the instance.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return void
 */
func (a IPAssociationsApi) DeleteIPAssociation(name string, cookie string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ip/association/{name}"
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
	var localVarAPIResponse = &APIResponse{Operation: "DeleteIPAssociation", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Retrieve Names of all IP Associations and Subcontainers in a Container
 * Retrieves the names of objects and subcontainers that you can access in the specified container.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param container Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access.
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *IpAssociationDiscoverResponse
 */
func (a IPAssociationsApi) DiscoverIPAssociation(container string, cookie string) (*IpAssociationDiscoverResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ip/association/{container}"
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
	var successPayload = new(IpAssociationDiscoverResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DiscoverIPAssociation", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Retrieve Names of Containers
 * Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *IpAssociationDiscoverResponse
 */
func (a IPAssociationsApi) DiscoverRootIPAssociation(cookie string) (*IpAssociationDiscoverResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ip/association/"

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
	var successPayload = new(IpAssociationDiscoverResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DiscoverRootIPAssociation", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Retrieve Details of an IP Association
 * &lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *IpAssociationResponse
 */
func (a IPAssociationsApi) GetIPAssociation(name string, cookie string) (*IpAssociationResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ip/association/{name}"
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
	var successPayload = new(IpAssociationResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetIPAssociation", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Retrieve Details of all IP Associations in a Container
 * Retrieves details of the IP associations that are available in the specified container and match the specified query criteria. If you don&#39;t specify any query criteria, then details of all the IP associations in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:&lt;p&gt;&lt;code&gt;?parameter1&#x3D;value1&amp;ampparameter2&#x3D;value2&amp;ampparameterN&#x3D;valueN&lt;/code&gt;&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param container &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt;
 * @param parentpool Use this option if you want to retrieve details of temporary IP addresses from the pool. Specify &lt;code&gt;ippool:/oracle/public/ippool&lt;/code&gt; as the value.
 * @param reservation Use this option if you want to retrieve details of a specific persistent IP address. Specify the name of the reservation in the format, &lt;code&gt;ipreservation:&lt;em&gt;ipreservation_name&lt;/em&gt;&lt;/code&gt;, where &lt;code&gt;&lt;em&gt;ipreservation_name&lt;/em&gt;&lt;/code&gt; is three-part name of an existing IP reservation in the &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt; format.
 * @param vcable vcable ID of the instance that you want to associate with the IP reservation.&lt;p&gt;For more information about the vcable of an instance, see &lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-instance-{name}-get.html\&quot;&gt;Retrieve Details of an Instance&lt;/a&gt;.
 * @param enabled Start the NAT Connection in enabled state. In disabled state networking will be torn down
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *IpAssociationListResponse
 */
func (a IPAssociationsApi) ListIPAssociation(container string, parentpool string, reservation string, vcable string, enabled bool, cookie string) (*IpAssociationListResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ip/association/{container}/"
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
	localVarQueryParams.Add("parentpool", a.Configuration.APIClient.ParameterToString(parentpool, ""))
	localVarQueryParams.Add("reservation", a.Configuration.APIClient.ParameterToString(reservation, ""))
	localVarQueryParams.Add("vcable", a.Configuration.APIClient.ParameterToString(vcable, ""))
	localVarQueryParams.Add("enabled", a.Configuration.APIClient.ParameterToString(enabled, ""))

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
	var successPayload = new(IpAssociationListResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListIPAssociation", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
