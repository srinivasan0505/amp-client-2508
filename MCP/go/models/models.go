package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DescribeRuleGroupsNamespaceResponseruleGroupsNamespace represents the DescribeRuleGroupsNamespaceResponseruleGroupsNamespace schema from the OpenAPI specification
type DescribeRuleGroupsNamespaceResponseruleGroupsNamespace struct {
	Modifiedat interface{} `json:"modifiedAt"`
	Name interface{} `json:"name"`
	Status CreateRuleGroupsNamespaceResponsestatus `json:"status"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
	Createdat interface{} `json:"createdAt"`
	Data interface{} `json:"data"`
}

// UpdateWorkspaceAliasRequest represents the UpdateWorkspaceAliasRequest schema from the OpenAPI specification
type UpdateWorkspaceAliasRequest struct {
	Alias interface{} `json:"alias,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// DescribeLoggingConfigurationResponseloggingConfiguration represents the DescribeLoggingConfigurationResponseloggingConfiguration schema from the OpenAPI specification
type DescribeLoggingConfigurationResponseloggingConfiguration struct {
	Modifiedat interface{} `json:"modifiedAt"`
	Status CreateLoggingConfigurationResponsestatus `json:"status"`
	Workspace interface{} `json:"workspace"`
	Createdat interface{} `json:"createdAt"`
	Loggrouparn interface{} `json:"logGroupArn"`
}

// CreateRuleGroupsNamespaceResponsestatus represents the CreateRuleGroupsNamespaceResponsestatus schema from the OpenAPI specification
type CreateRuleGroupsNamespaceResponsestatus struct {
	Statusreason interface{} `json:"statusReason,omitempty"`
	Statuscode interface{} `json:"statusCode"`
}

// ListWorkspacesRequest represents the ListWorkspacesRequest schema from the OpenAPI specification
type ListWorkspacesRequest struct {
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// AlertManagerDefinitionDescription represents the AlertManagerDefinitionDescription schema from the OpenAPI specification
type AlertManagerDefinitionDescription struct {
	Modifiedat interface{} `json:"modifiedAt"`
	Status CreateAlertManagerDefinitionResponsestatus `json:"status"`
	Createdat interface{} `json:"createdAt"`
	Data interface{} `json:"data"`
}

// LoggingConfigurationMetadata represents the LoggingConfigurationMetadata schema from the OpenAPI specification
type LoggingConfigurationMetadata struct {
	Loggrouparn interface{} `json:"logGroupArn"`
	Modifiedat interface{} `json:"modifiedAt"`
	Status CreateLoggingConfigurationResponsestatus `json:"status"`
	Workspace interface{} `json:"workspace"`
	Createdat interface{} `json:"createdAt"`
}

// AlertManagerDefinitionStatus represents the AlertManagerDefinitionStatus schema from the OpenAPI specification
type AlertManagerDefinitionStatus struct {
	Statuscode interface{} `json:"statusCode"`
	Statusreason interface{} `json:"statusReason,omitempty"`
}

// ListRuleGroupsNamespacesRequest represents the ListRuleGroupsNamespacesRequest schema from the OpenAPI specification
type ListRuleGroupsNamespacesRequest struct {
}

// WorkspaceDescriptionstatus represents the WorkspaceDescriptionstatus schema from the OpenAPI specification
type WorkspaceDescriptionstatus struct {
	Statuscode interface{} `json:"statusCode"`
}

// PutRuleGroupsNamespacerequest represents the PutRuleGroupsNamespacerequest schema from the OpenAPI specification
type PutRuleGroupsNamespacerequest struct {
	Clienttoken string `json:"clientToken,omitempty"` // An identifier used to ensure the idempotency of a write request.
	Data string `json:"data"` // The rule groups namespace data.
}

// DescribeLoggingConfigurationRequest represents the DescribeLoggingConfigurationRequest schema from the OpenAPI specification
type DescribeLoggingConfigurationRequest struct {
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// RuleGroupsNamespaceStatus represents the RuleGroupsNamespaceStatus schema from the OpenAPI specification
type RuleGroupsNamespaceStatus struct {
	Statusreason interface{} `json:"statusReason,omitempty"`
	Statuscode interface{} `json:"statusCode"`
}

// UpdateLoggingConfigurationRequest represents the UpdateLoggingConfigurationRequest schema from the OpenAPI specification
type UpdateLoggingConfigurationRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Loggrouparn interface{} `json:"logGroupArn"`
}

// LoggingConfigurationStatus represents the LoggingConfigurationStatus schema from the OpenAPI specification
type LoggingConfigurationStatus struct {
	Statuscode interface{} `json:"statusCode"`
	Statusreason interface{} `json:"statusReason,omitempty"`
}

// CreateWorkspaceResponsestatus represents the CreateWorkspaceResponsestatus schema from the OpenAPI specification
type CreateWorkspaceResponsestatus struct {
	Statuscode interface{} `json:"statusCode"`
}

// TagResourcerequest represents the TagResourcerequest schema from the OpenAPI specification
type TagResourcerequest struct {
	Tags map[string]interface{} `json:"tags"` // The list of tags assigned to the resource.
}

// UpdateWorkspaceAliasrequest represents the UpdateWorkspaceAliasrequest schema from the OpenAPI specification
type UpdateWorkspaceAliasrequest struct {
	Alias string `json:"alias,omitempty"` // A user-assigned workspace alias.
	Clienttoken string `json:"clientToken,omitempty"` // An identifier used to ensure the idempotency of a write request.
}

// DeleteAlertManagerDefinitionRequest represents the DeleteAlertManagerDefinitionRequest schema from the OpenAPI specification
type DeleteAlertManagerDefinitionRequest struct {
}

// CreateRuleGroupsNamespacerequest represents the CreateRuleGroupsNamespacerequest schema from the OpenAPI specification
type CreateRuleGroupsNamespacerequest struct {
	Name string `json:"name"` // The namespace name that the rule group belong to.
	Tags map[string]interface{} `json:"tags,omitempty"` // The list of tags assigned to the resource.
	Clienttoken string `json:"clientToken,omitempty"` // An identifier used to ensure the idempotency of a write request.
	Data string `json:"data"` // The rule groups namespace data.
}

// CreateAlertManagerDefinitionRequest represents the CreateAlertManagerDefinitionRequest schema from the OpenAPI specification
type CreateAlertManagerDefinitionRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Data interface{} `json:"data"`
}

// CreateAlertManagerDefinitionResponse represents the CreateAlertManagerDefinitionResponse schema from the OpenAPI specification
type CreateAlertManagerDefinitionResponse struct {
	Status CreateAlertManagerDefinitionResponsestatus `json:"status"`
}

// ListWorkspacesResponse represents the ListWorkspacesResponse schema from the OpenAPI specification
type ListWorkspacesResponse struct {
	Workspaces interface{} `json:"workspaces"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags map[string]interface{} `json:"tags,omitempty"` // The list of tags assigned to the resource.
}

// DeleteRuleGroupsNamespaceRequest represents the DeleteRuleGroupsNamespaceRequest schema from the OpenAPI specification
type DeleteRuleGroupsNamespaceRequest struct {
}

// WorkspaceStatus represents the WorkspaceStatus schema from the OpenAPI specification
type WorkspaceStatus struct {
	Statuscode interface{} `json:"statusCode"`
}

// CreateAlertManagerDefinitionResponsestatus represents the CreateAlertManagerDefinitionResponsestatus schema from the OpenAPI specification
type CreateAlertManagerDefinitionResponsestatus struct {
	Statuscode interface{} `json:"statusCode"`
	Statusreason interface{} `json:"statusReason,omitempty"`
}

// DescribeAlertManagerDefinitionRequest represents the DescribeAlertManagerDefinitionRequest schema from the OpenAPI specification
type DescribeAlertManagerDefinitionRequest struct {
}

// CreateWorkspaceRequest represents the CreateWorkspaceRequest schema from the OpenAPI specification
type CreateWorkspaceRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Alias interface{} `json:"alias,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// PutAlertManagerDefinitionrequest represents the PutAlertManagerDefinitionrequest schema from the OpenAPI specification
type PutAlertManagerDefinitionrequest struct {
	Clienttoken string `json:"clientToken,omitempty"` // An identifier used to ensure the idempotency of a write request.
	Data string `json:"data"` // The alert manager definition data.
}

// ListRuleGroupsNamespacesResponse represents the ListRuleGroupsNamespacesResponse schema from the OpenAPI specification
type ListRuleGroupsNamespacesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Rulegroupsnamespaces interface{} `json:"ruleGroupsNamespaces"`
}

// DescribeWorkspaceResponseworkspace represents the DescribeWorkspaceResponseworkspace schema from the OpenAPI specification
type DescribeWorkspaceResponseworkspace struct {
	Alias interface{} `json:"alias,omitempty"`
	Arn interface{} `json:"arn"`
	Createdat interface{} `json:"createdAt"`
	Prometheusendpoint interface{} `json:"prometheusEndpoint,omitempty"`
	Status WorkspaceDescriptionstatus `json:"status"`
	Tags interface{} `json:"tags,omitempty"`
	Workspaceid interface{} `json:"workspaceId"`
}

// WorkspaceDescription represents the WorkspaceDescription schema from the OpenAPI specification
type WorkspaceDescription struct {
	Arn interface{} `json:"arn"`
	Createdat interface{} `json:"createdAt"`
	Prometheusendpoint interface{} `json:"prometheusEndpoint,omitempty"`
	Status WorkspaceDescriptionstatus `json:"status"`
	Tags interface{} `json:"tags,omitempty"`
	Workspaceid interface{} `json:"workspaceId"`
	Alias interface{} `json:"alias,omitempty"`
}

// PutAlertManagerDefinitionResponse represents the PutAlertManagerDefinitionResponse schema from the OpenAPI specification
type PutAlertManagerDefinitionResponse struct {
	Status CreateAlertManagerDefinitionResponsestatus `json:"status"`
}

// DescribeAlertManagerDefinitionResponsealertManagerDefinition represents the DescribeAlertManagerDefinitionResponsealertManagerDefinition schema from the OpenAPI specification
type DescribeAlertManagerDefinitionResponsealertManagerDefinition struct {
	Modifiedat interface{} `json:"modifiedAt"`
	Status CreateAlertManagerDefinitionResponsestatus `json:"status"`
	Createdat interface{} `json:"createdAt"`
	Data interface{} `json:"data"`
}

// PutRuleGroupsNamespaceRequest represents the PutRuleGroupsNamespaceRequest schema from the OpenAPI specification
type PutRuleGroupsNamespaceRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Data interface{} `json:"data"`
}

// DescribeAlertManagerDefinitionResponse represents the DescribeAlertManagerDefinitionResponse schema from the OpenAPI specification
type DescribeAlertManagerDefinitionResponse struct {
	Alertmanagerdefinition DescribeAlertManagerDefinitionResponsealertManagerDefinition `json:"alertManagerDefinition"`
}

// DescribeRuleGroupsNamespaceRequest represents the DescribeRuleGroupsNamespaceRequest schema from the OpenAPI specification
type DescribeRuleGroupsNamespaceRequest struct {
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags map[string]interface{} `json:"tags"` // The list of tags assigned to the resource.
}

// WorkspaceSummary represents the WorkspaceSummary schema from the OpenAPI specification
type WorkspaceSummary struct {
	Status WorkspaceDescriptionstatus `json:"status"`
	Tags interface{} `json:"tags,omitempty"`
	Workspaceid interface{} `json:"workspaceId"`
	Alias interface{} `json:"alias,omitempty"`
	Arn interface{} `json:"arn"`
	Createdat interface{} `json:"createdAt"`
}

// DeleteLoggingConfigurationRequest represents the DeleteLoggingConfigurationRequest schema from the OpenAPI specification
type DeleteLoggingConfigurationRequest struct {
}

// CreateRuleGroupsNamespaceResponse represents the CreateRuleGroupsNamespaceResponse schema from the OpenAPI specification
type CreateRuleGroupsNamespaceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
	Name interface{} `json:"name"`
	Status CreateRuleGroupsNamespaceResponsestatus `json:"status"`
}

// CreateWorkspacerequest represents the CreateWorkspacerequest schema from the OpenAPI specification
type CreateWorkspacerequest struct {
	Alias string `json:"alias,omitempty"` // A user-assigned workspace alias.
	Clienttoken string `json:"clientToken,omitempty"` // An identifier used to ensure the idempotency of a write request.
	Tags map[string]interface{} `json:"tags,omitempty"` // The list of tags assigned to the resource.
}

// CreateLoggingConfigurationResponse represents the CreateLoggingConfigurationResponse schema from the OpenAPI specification
type CreateLoggingConfigurationResponse struct {
	Status CreateLoggingConfigurationResponsestatus `json:"status"`
}

// CreateLoggingConfigurationResponsestatus represents the CreateLoggingConfigurationResponsestatus schema from the OpenAPI specification
type CreateLoggingConfigurationResponsestatus struct {
	Statuscode interface{} `json:"statusCode"`
	Statusreason interface{} `json:"statusReason,omitempty"`
}

// DeleteWorkspaceRequest represents the DeleteWorkspaceRequest schema from the OpenAPI specification
type DeleteWorkspaceRequest struct {
}

// PutAlertManagerDefinitionRequest represents the PutAlertManagerDefinitionRequest schema from the OpenAPI specification
type PutAlertManagerDefinitionRequest struct {
	Data interface{} `json:"data"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// RuleGroupsNamespaceDescription represents the RuleGroupsNamespaceDescription schema from the OpenAPI specification
type RuleGroupsNamespaceDescription struct {
	Createdat interface{} `json:"createdAt"`
	Data interface{} `json:"data"`
	Modifiedat interface{} `json:"modifiedAt"`
	Name interface{} `json:"name"`
	Status CreateRuleGroupsNamespaceResponsestatus `json:"status"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// DescribeWorkspaceRequest represents the DescribeWorkspaceRequest schema from the OpenAPI specification
type DescribeWorkspaceRequest struct {
}

// CreateRuleGroupsNamespaceRequest represents the CreateRuleGroupsNamespaceRequest schema from the OpenAPI specification
type CreateRuleGroupsNamespaceRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Data interface{} `json:"data"`
	Name interface{} `json:"name"`
}

// CreateWorkspaceResponse represents the CreateWorkspaceResponse schema from the OpenAPI specification
type CreateWorkspaceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
	Workspaceid interface{} `json:"workspaceId"`
	Arn interface{} `json:"arn"`
	Status CreateWorkspaceResponsestatus `json:"status"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// RuleGroupsNamespaceSummary represents the RuleGroupsNamespaceSummary schema from the OpenAPI specification
type RuleGroupsNamespaceSummary struct {
	Createdat interface{} `json:"createdAt"`
	Modifiedat interface{} `json:"modifiedAt"`
	Name interface{} `json:"name"`
	Status CreateRuleGroupsNamespaceResponsestatus `json:"status"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
}

// CreateLoggingConfigurationRequest represents the CreateLoggingConfigurationRequest schema from the OpenAPI specification
type CreateLoggingConfigurationRequest struct {
	Loggrouparn interface{} `json:"logGroupArn"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// DescribeRuleGroupsNamespaceResponse represents the DescribeRuleGroupsNamespaceResponse schema from the OpenAPI specification
type DescribeRuleGroupsNamespaceResponse struct {
	Rulegroupsnamespace DescribeRuleGroupsNamespaceResponseruleGroupsNamespace `json:"ruleGroupsNamespace"`
}

// DescribeLoggingConfigurationResponse represents the DescribeLoggingConfigurationResponse schema from the OpenAPI specification
type DescribeLoggingConfigurationResponse struct {
	Loggingconfiguration DescribeLoggingConfigurationResponseloggingConfiguration `json:"loggingConfiguration"`
}

// UpdateLoggingConfigurationrequest represents the UpdateLoggingConfigurationrequest schema from the OpenAPI specification
type UpdateLoggingConfigurationrequest struct {
	Loggrouparn string `json:"logGroupArn"` // The ARN of the CW log group to which the vended log data will be published.
	Clienttoken string `json:"clientToken,omitempty"` // An identifier used to ensure the idempotency of a write request.
}

// PutRuleGroupsNamespaceResponse represents the PutRuleGroupsNamespaceResponse schema from the OpenAPI specification
type PutRuleGroupsNamespaceResponse struct {
	Status CreateRuleGroupsNamespaceResponsestatus `json:"status"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
	Name interface{} `json:"name"`
}

// UpdateLoggingConfigurationResponse represents the UpdateLoggingConfigurationResponse schema from the OpenAPI specification
type UpdateLoggingConfigurationResponse struct {
	Status CreateLoggingConfigurationResponsestatus `json:"status"`
}

// DescribeWorkspaceResponse represents the DescribeWorkspaceResponse schema from the OpenAPI specification
type DescribeWorkspaceResponse struct {
	Workspace DescribeWorkspaceResponseworkspace `json:"workspace"`
}
