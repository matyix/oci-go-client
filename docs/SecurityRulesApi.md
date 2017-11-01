# \SecurityRulesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecurityRule**](SecurityRulesApi.md#AddSecurityRule) | **Post** /network/v1/secrule/ | Create a Security Rule for IP Networks
[**DeleteSecurityRule**](SecurityRulesApi.md#DeleteSecurityRule) | **Delete** /network/v1/secrule/{name} | Delete a Security Rule Used in IP Networks
[**GetSecurityRule**](SecurityRulesApi.md#GetSecurityRule) | **Get** /network/v1/secrule/{name} | Retrieve Details of a Security Rule Used in IP Networks
[**ListSecurityRule**](SecurityRulesApi.md#ListSecurityRule) | **Get** /network/v1/secrule/{container}/ | Retrieve Details of all Security Rules in a Container
[**UpdateSecurityRule**](SecurityRulesApi.md#UpdateSecurityRule) | **Put** /network/v1/secrule/{name} | Update a Security Rule Used in IP Networks


# **AddSecurityRule**
> SecurityRuleResponse AddSecurityRule($body, $cookie)

Create a Security Rule for IP Networks

 Adds a security rule. A security rule permits traffic from a specified source or to a specified destination. You must specify the direction of a security rule - either ingress or egress. In addition, you can specify the source or destination of permitted traffic, and the security protocol and port used to send or receive packets. Each of the parameters that you specify in a security rule provides a criterion that the type of traffic permitted by that rule must match. Only packets that match all of the specified criteria are permitted. If you don't specify match criteria in the security rule, all traffic in the specified direction is permitted.<p>When you create a security rule with a specified direction, say ingress, you should also create a corresponding security rule for the opposite direction - in this case, egress. This is generally required to ensure that when traffic is permitted in one direction, responses or acknowledgement packets in the opposite direction are also permitted.<p>When you create a security rule, you specify the ACL that it belongs to. ACLs apply to vNICsets. You can apply multiple ACLs to a vNICset and you can apply each ACL to multiple vNICsets. When an ACL is applied to a vNICset, every security rule that belongs to the ACL applies to every vNIC that is specified in the vNICset.<p>A security rule allows you to specify the following parameters:<p>* The flow direction - ingress or egress<p>* (Optional) A source vNICset or a list of source IP address prefix sets, or both<p>* (Optional) A destination vNICset or a list of destination IP address prefix sets, or both<p>* (Optional) A list of security protocols<p>* (Optional) The name of the ACL that contains this rule<p>* (Optional) An option to disable the security rule<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SecurityRulePostRequest**](SecurityRulePostRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SecurityRuleResponse**](SecurityRule-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecurityRule**
> DeleteSecurityRule($name, $cookie)

Delete a Security Rule Used in IP Networks

Deletes the specified security rule. Before deleting a security rule, ensure that it is not being used.<p>If you want to disable a security rule, use the <code>PUT /network/v1/secrule/</code> method to change the  value of <code>enabledFlag</code> to <code>false</code>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecurityRule**
> SecurityRuleResponse GetSecurityRule($name, $cookie)

Retrieve Details of a Security Rule Used in IP Networks

Retrieves details of the specified security rule.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SecurityRuleResponse**](SecurityRule-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecurityRule**
> SecurityRuleListResponse ListSecurityRule($container, $cookie)

Retrieve Details of all Security Rules in a Container

Retrieves details of all the security rules in the specified container. This request is for security rules used in IP networks.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SecurityRuleListResponse**](SecurityRule-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityRule**
> SecurityRuleResponse UpdateSecurityRule($name, $body, $cookie)

Update a Security Rule Used in IP Networks

You can update values of all the parameters of a security rule except the name. You can also disable a security rule, by setting the value of the <code>enabledFlag</code> parameter as <code>false</code>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**SecurityRulePutRequest**](SecurityRulePutRequest.md)|  | 
 **cookie** | **string**| The Cookie: header must be included with every request to the service. It must be set to the value of the set-cookie header in the response received to the POST /authenticate/ call. | [optional] 

### Return type

[**SecurityRuleResponse**](SecurityRule-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

