package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-simple-queue-service/mcp-server/config"
	"github.com/amazon-simple-queue-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_removepermissionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		AccountNumberVal, ok := args["AccountNumber"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: AccountNumber"), nil
		}
		AccountNumber, ok := AccountNumberVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: AccountNumber"), nil
		}
		QueueNameVal, ok := args["QueueName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: QueueName"), nil
		}
		QueueName, ok := QueueNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: QueueName"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["Label"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Label=%v", val))
		}
		if val, ok := args["Action"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Action=%v", val))
		}
		if val, ok := args["Version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Version=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/%s/%s/#Action=RemovePermission%s", cfg.BaseURL, AccountNumber, QueueName, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_removepermissionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_AccountNumber_QueueName_#Action=RemovePermission",
		mcp.WithDescription("<p>Revokes any permissions in the queue policy that matches the specified <code>Label</code> parameter.</p> <note> <ul> <li> <p>Only the owner of a queue can remove permissions from it.</p> </li> <li> <p>Cross-account permissions don't apply to this action. For more information, see <a href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name">Grant cross-account permissions to a role and a username</a> in the <i>Amazon SQS Developer Guide</i>.</p> </li> <li> <p>To remove the ability to change queue permissions, you must deny permission to the <code>AddPermission</code>, <code>RemovePermission</code>, and <code>SetQueueAttributes</code> actions in your IAM policy.</p> </li> </ul> </note>"),
		mcp.WithString("Label", mcp.Required(), mcp.Description("The identification of the permission to remove. This is the label added using the <code> <a>AddPermission</a> </code> action.")),
		mcp.WithNumber("AccountNumber", mcp.Required(), mcp.Description("The AWS account number")),
		mcp.WithString("QueueName", mcp.Required(), mcp.Description("The name of the queue")),
		mcp.WithString("Action", mcp.Required(), mcp.Description("")),
		mcp.WithString("Version", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_removepermissionHandler(cfg),
	}
}
