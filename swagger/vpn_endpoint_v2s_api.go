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

type VPNEndpointV2sApi struct {
	Configuration *Configuration
}

func NewVPNEndpointV2sApi() *VPNEndpointV2sApi {
	configuration := NewConfiguration()
	return &VPNEndpointV2sApi{
		Configuration: configuration,
	}
}

func NewVPNEndpointV2sApiWithBasePath(basePath string) *VPNEndpointV2sApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &VPNEndpointV2sApi{
		Configuration: configuration,
	}
}

/**
 * Create a VPN Endpoint V2
 * This endpoint is not available on Oracle Cloud Machine.&lt;p&gt;Creates an IPSec VPN connection from Oracle Cloud to your data centers using VPN as a Service (VPNaaS).&lt;p&gt;&lt;b&gt;Note:&lt;/b&gt; You can use VPNaaS to set up a tunnel to instances that are on IP networks. However, VPNaaS doesn&#39;t support VPN connections to instances that don&#39;t have any interface on IP networks. To establish a VPN tunnel to instances that are on the shared network, follow the steps for creating a single-homed Corente Services Gateway instance in &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;STCSG-GUID-07FAD4A9-DDE4-4391-BB9C-EA4784552EDA\&quot;&gt;Setting Up VPN&lt;/a&gt; in &lt;i&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/i&gt;.&lt;p&gt;&lt;b&gt;Prerequisites&lt;/b&gt;&lt;p&gt;Ensure that you complete the following tasks and noted the required information before creating a VPN connection.&lt;p&gt;* Create an IP network or use an existing IP network. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;STCSG-GUID-9A826000-2728-4837-905A-7835FA775F9B\&quot;&gt;Creating an IP Network&lt;/a&gt; in &lt;i&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/i&gt;. Make a note of the name of this IP network.&lt;p&gt;* Configure a supported third-party VPN device at your data center and make a note of the public IP address of this gateway. The third-party VPN device must be ready for the VPN connection to be established. For information about certified third-party VPN device configurations, see &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;STCSG-GUID-D5186CEE-08A9-4673-BD80-91CC7646988A\&quot;&gt;About Setting Up VPN&lt;/a&gt; in &lt;i&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/i&gt;.&lt;p&gt;* Ensure that you have the pre-shared key (PSK) that you want to use for this VPN connection.&lt;p&gt;* Create a vNICset. When you create instances, specify this vNICset for each vNIC that is added to an IP network that will be reachable over the VPN connection. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;cloud&amp;id&#x3D;STCSG-GUID-565766C2-21AF-4EF2-B7C4-728B7604543A\&quot;&gt;Creating a vNICset&lt;/a&gt; in &lt;i&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/i&gt;.&lt;p&gt;While your VPN connection is being configured, its status is &lt;code&gt;PENDING&lt;/code&gt;. It can take around 20 to 30 minutes for your VPN gateway to be created. When the cloud VPN gateway is created, the &lt;code&gt;localGatewayAddress&lt;/code&gt; parameter provides its public IP address. To monitor the status of your VPN connection and retrieve the public IP address of the cloud VPN gateway, send the &lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-vpnendpoint-v2-{name}-get.html\&quot;&gt;GET /vpnendpoint/v2/{name}&lt;/a&gt; request. You&#39;ll have to update the third-party VPN device in your data center with the public IP address of your cloud VPN gateway. If the third-party device in your data center is configured and ready, the VPN connection is established. The value of the &lt;code&gt;tunnelStatus&lt;/code&gt; parameter changes to &lt;code&gt;UP&lt;/code&gt; when the connection is established.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param body 
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *VpnConnectionResponse
 */
func (a VPNEndpointV2sApi) AddVPNEndpointV2(body VpnConnectionPostRequest, cookie string) (*VpnConnectionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/vpnendpoint/v2/"

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
	var successPayload = new(VpnConnectionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "AddVPNEndpointV2", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Delete a VPN Endpoint V2
 * This endpoint is not available on Oracle Cloud Machine.&lt;p&gt;You may want to add new vNIC sets or update other parameters of an existing VPN connection. Not all parameters of a VPN connection can be updated using the PUT /vpnendpoint/v2/{name} request. In such cases, you can delete the VPN connection, and then recreate the it with the updated parameters by using the POST /vpnendpoint/v2/ request.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return void
 */
func (a VPNEndpointV2sApi) DeleteVPNEndpointV2(name string, cookie string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/vpnendpoint/v2/{name}"
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
	var localVarAPIResponse = &APIResponse{Operation: "DeleteVPNEndpointV2", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Retrieve Details of a VPN Endpoint V2
 * This endpoint is not available on Oracle Cloud Machine.&lt;p&gt;Retrieve Details of the specified VPN connection. You can retrieve details of a VPN connection to track the status of the tunnel. The tunnel can be in one of the following states:&lt;p&gt;* PENDING: indicates that your VPN connection is being set up.&lt;p&gt;* UP: indicates that your VPN connection is established.&lt;p&gt;* DOWN: indicates that your VPN connection is down.&lt;p&gt;* ERROR: indicates that your VPN connection is in the error state.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *VpnConnectionResponse
 */
func (a VPNEndpointV2sApi) GetVPNEndpointV2(name string, cookie string) (*VpnConnectionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/vpnendpoint/v2/{name}"
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
	var successPayload = new(VpnConnectionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetVPNEndpointV2", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Retrieve Details of all VPN Endpoint V2s in a Container
 * This endpoint is not available on Oracle Cloud Machine.&lt;p&gt;Retrieves details of the VPN connections in the specified container and match the specified query criteria. If you don&#39;t specify any query criteria, then details of all the VPN connections in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:&lt;p&gt;&lt;code&gt;?parameter1&#x3D;value1&amp;ampparameter2&#x3D;value2&amp;ampparameterN&#x3D;valueN&lt;/code&gt;&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param container &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt;
 * @param name Specify the three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the VPN connection.
 * @param customerVpnGateway The public IP address of the VPN gateway in your data center through which you want to connect to the Oracle Cloud VPN gateway.
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *VpnConnectionListResponse
 */
func (a VPNEndpointV2sApi) ListVPNEndpointV2(container string, name string, customerVpnGateway string, cookie string) (*VpnConnectionListResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/vpnendpoint/v2/{container}/"
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
	localVarQueryParams.Add("customer_vpn_gateway", a.Configuration.APIClient.ParameterToString(customerVpnGateway, ""))

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
	var successPayload = new(VpnConnectionListResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListVPNEndpointV2", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Update a VPN Endpoint V2
 * This endpoint is not available on Oracle Cloud Machine.&lt;p&gt;Updates the values for &lt;code&gt;psk&lt;/code&gt; and &lt;code&gt;reachable_routes&lt;/code&gt; for the specified VPN connection. If you want to update values for any other parameter, you&#39;ll have to delete the VPN connection and then re-create with the new parameters using the POST /vpnendpoint/v2/ request. Although you can only update the values for &lt;code&gt;psk&lt;/code&gt; and &lt;code&gt;reachable_routes&lt;/code&gt; using this request, you must specify the current values for all the existing parameters in the request body.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param body 
 * @param cookie The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call.
 * @return *VpnConnectionResponse
 */
func (a VPNEndpointV2sApi) UpdateVPNEndpointV2(name string, body VpnConnectionPutRequest, cookie string) (*VpnConnectionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/vpnendpoint/v2/{name}"
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
	var successPayload = new(VpnConnectionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateVPNEndpointV2", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

