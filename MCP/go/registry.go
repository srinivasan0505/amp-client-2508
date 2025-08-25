package main

import (
	"github.com/amazon-prometheus-service/mcp-server/config"
	"github.com/amazon-prometheus-service/mcp-server/models"
	tools_workspaces "github.com/amazon-prometheus-service/mcp-server/tools/workspaces"
	tools_tags "github.com/amazon-prometheus-service/mcp-server/tools/tags"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_workspaces.CreateDeleteworkspaceTool(cfg),
		tools_workspaces.CreateDescribeworkspaceTool(cfg),
		tools_workspaces.CreateUpdateworkspacealiasTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_workspaces.CreateCreatealertmanagerdefinitionTool(cfg),
		tools_workspaces.CreatePutalertmanagerdefinitionTool(cfg),
		tools_workspaces.CreateDeletealertmanagerdefinitionTool(cfg),
		tools_workspaces.CreateDescribealertmanagerdefinitionTool(cfg),
		tools_workspaces.CreateListrulegroupsnamespacesTool(cfg),
		tools_workspaces.CreateCreaterulegroupsnamespaceTool(cfg),
		tools_workspaces.CreateDeleterulegroupsnamespaceTool(cfg),
		tools_workspaces.CreateDescriberulegroupsnamespaceTool(cfg),
		tools_workspaces.CreatePutrulegroupsnamespaceTool(cfg),
		tools_workspaces.CreateListworkspacesTool(cfg),
		tools_workspaces.CreateCreateworkspaceTool(cfg),
		tools_workspaces.CreateDeleteloggingconfigurationTool(cfg),
		tools_workspaces.CreateDescribeloggingconfigurationTool(cfg),
		tools_workspaces.CreateCreateloggingconfigurationTool(cfg),
		tools_workspaces.CreateUpdateloggingconfigurationTool(cfg),
	}
}
