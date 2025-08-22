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

func Get_tagqueueHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["Tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Tags=%v", val))
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
		url := fmt.Sprintf("%s/%s/%s/#Action=TagQueue%s", cfg.BaseURL, AccountNumber, QueueName, queryString)
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

func CreateGet_tagqueueTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_AccountNumber_QueueName_#Action=TagQueue",
		mcp.WithDescription("<p>Add cost allocation tags to the specified Amazon SQS queue. For an overview, see <a href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html">Tagging Your Amazon SQS Queues</a> in the <i>Amazon SQS Developer Guide</i>.</p> <p>When you use queue tags, keep the following guidelines in mind:</p> <ul> <li> <p>Adding more than 50 tags to a queue isn't recommended.</p> </li> <li> <p>Tags don't have any semantic meaning. Amazon SQS interprets tags as character strings.</p> </li> <li> <p>Tags are case-sensitive.</p> </li> <li> <p>A new tag with a key identical to that of an existing tag overwrites the existing tag.</p> </li> </ul> <p>For a full list of tag restrictions, see <a href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html#limits-queues">Quotas related to queues</a> in the <i>Amazon SQS Developer Guide</i>.</p> <note> <p>Cross-account permissions don't apply to this action. For more information, see <a href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name">Grant cross-account permissions to a role and a username</a> in the <i>Amazon SQS Developer Guide</i>.</p> </note>"),
		mcp.WithObject("Tags", mcp.Required(), mcp.Description("The list of tags to be added to the specified queue.")),
		mcp.WithNumber("AccountNumber", mcp.Required(), mcp.Description("The AWS account number")),
		mcp.WithString("QueueName", mcp.Required(), mcp.Description("The name of the queue")),
		mcp.WithString("Action", mcp.Required(), mcp.Description("")),
		mcp.WithString("Version", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_tagqueueHandler(cfg),
	}
}
