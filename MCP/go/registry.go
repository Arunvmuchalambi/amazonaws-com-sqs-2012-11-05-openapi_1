package main

import (
	"github.com/amazon-simple-queue-service/mcp-server/config"
	"github.com/amazon-simple-queue-service/mcp-server/models"
	tools_accountnumber "github.com/amazon-simple-queue-service/mcp-server/tools/accountnumber"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_accountnumber.CreateGet_deletequeueTool(cfg),
		tools_accountnumber.CreateGet_tagqueueTool(cfg),
		tools_accountnumber.CreateGet_removepermissionTool(cfg),
		tools_accountnumber.CreateGet_untagqueueTool(cfg),
	}
}
