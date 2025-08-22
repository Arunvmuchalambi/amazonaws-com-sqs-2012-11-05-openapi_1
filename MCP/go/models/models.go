package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListDeadLetterSourceQueuesResult represents the ListDeadLetterSourceQueuesResult schema from the OpenAPI specification
type ListDeadLetterSourceQueuesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Queueurls interface{} `json:"queueUrls"`
}

// ChangeMessageVisibilityBatchRequest represents the ChangeMessageVisibilityBatchRequest schema from the OpenAPI specification
type ChangeMessageVisibilityBatchRequest struct {
	Entries interface{} `json:"Entries"`
	Queueurl interface{} `json:"QueueUrl"`
}

// SendMessageBatchResult represents the SendMessageBatchResult schema from the OpenAPI specification
type SendMessageBatchResult struct {
	Failed interface{} `json:"Failed"`
	Successful interface{} `json:"Successful"`
}

// MessageSystemAttributeMap represents the MessageSystemAttributeMap schema from the OpenAPI specification
type MessageSystemAttributeMap struct {
}

// PurgeQueueRequest represents the PurgeQueueRequest schema from the OpenAPI specification
type PurgeQueueRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
}

// GetQueueAttributesResult represents the GetQueueAttributesResult schema from the OpenAPI specification
type GetQueueAttributesResult struct {
	Attributes interface{} `json:"Attributes,omitempty"`
}

// ListDeadLetterSourceQueuesRequest represents the ListDeadLetterSourceQueuesRequest schema from the OpenAPI specification
type ListDeadLetterSourceQueuesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Queueurl interface{} `json:"QueueUrl"`
}

// AddPermissionRequest represents the AddPermissionRequest schema from the OpenAPI specification
type AddPermissionRequest struct {
	Awsaccountids interface{} `json:"AWSAccountIds"`
	Actions interface{} `json:"Actions"`
	Label interface{} `json:"Label"`
	Queueurl interface{} `json:"QueueUrl"`
}

// ListMessageMoveTasksRequest represents the ListMessageMoveTasksRequest schema from the OpenAPI specification
type ListMessageMoveTasksRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sourcearn interface{} `json:"SourceArn"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// ChangeMessageVisibilityRequest represents the ChangeMessageVisibilityRequest schema from the OpenAPI specification
type ChangeMessageVisibilityRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
	Receipthandle interface{} `json:"ReceiptHandle"`
	Visibilitytimeout interface{} `json:"VisibilityTimeout"`
}

// CancelMessageMoveTaskRequest represents the CancelMessageMoveTaskRequest schema from the OpenAPI specification
type CancelMessageMoveTaskRequest struct {
	Taskhandle interface{} `json:"TaskHandle"`
}

// CreateQueueResult represents the CreateQueueResult schema from the OpenAPI specification
type CreateQueueResult struct {
	Queueurl interface{} `json:"QueueUrl,omitempty"`
}

// CancelMessageMoveTaskResult represents the CancelMessageMoveTaskResult schema from the OpenAPI specification
type CancelMessageMoveTaskResult struct {
	Approximatenumberofmessagesmoved interface{} `json:"ApproximateNumberOfMessagesMoved,omitempty"`
}

// ChangeMessageVisibilityBatchResult represents the ChangeMessageVisibilityBatchResult schema from the OpenAPI specification
type ChangeMessageVisibilityBatchResult struct {
	Successful interface{} `json:"Successful"`
	Failed interface{} `json:"Failed"`
}

// DeleteMessageRequest represents the DeleteMessageRequest schema from the OpenAPI specification
type DeleteMessageRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
	Receipthandle interface{} `json:"ReceiptHandle"`
}

// MessageSystemAttributeValue represents the MessageSystemAttributeValue schema from the OpenAPI specification
type MessageSystemAttributeValue struct {
	Binarylistvalues interface{} `json:"BinaryListValues,omitempty"`
	Binaryvalue interface{} `json:"BinaryValue,omitempty"`
	Datatype interface{} `json:"DataType"`
	Stringlistvalues interface{} `json:"StringListValues,omitempty"`
	Stringvalue interface{} `json:"StringValue,omitempty"`
}

// DeleteMessageBatchRequest represents the DeleteMessageBatchRequest schema from the OpenAPI specification
type DeleteMessageBatchRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
	Entries interface{} `json:"Entries"`
}

// DeleteQueueRequest represents the DeleteQueueRequest schema from the OpenAPI specification
type DeleteQueueRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
}

// SendMessageBatchRequestEntry represents the SendMessageBatchRequestEntry schema from the OpenAPI specification
type SendMessageBatchRequestEntry struct {
	Messagegroupid interface{} `json:"MessageGroupId,omitempty"`
	Messagesystemattributes interface{} `json:"MessageSystemAttributes,omitempty"`
	Delayseconds interface{} `json:"DelaySeconds,omitempty"`
	Id interface{} `json:"Id"`
	Messageattributes interface{} `json:"MessageAttributes,omitempty"`
	Messagebody interface{} `json:"MessageBody"`
	Messagededuplicationid interface{} `json:"MessageDeduplicationId,omitempty"`
}

// MessageBodyAttributeMap represents the MessageBodyAttributeMap schema from the OpenAPI specification
type MessageBodyAttributeMap struct {
}

// DeleteMessageBatchResultEntry represents the DeleteMessageBatchResultEntry schema from the OpenAPI specification
type DeleteMessageBatchResultEntry struct {
	Id interface{} `json:"Id"`
}

// ListQueueTagsResult represents the ListQueueTagsResult schema from the OpenAPI specification
type ListQueueTagsResult struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// SendMessageResult represents the SendMessageResult schema from the OpenAPI specification
type SendMessageResult struct {
	Md5ofmessagebody interface{} `json:"MD5OfMessageBody,omitempty"`
	Md5ofmessagesystemattributes interface{} `json:"MD5OfMessageSystemAttributes,omitempty"`
	Messageid interface{} `json:"MessageId,omitempty"`
	Sequencenumber interface{} `json:"SequenceNumber,omitempty"`
	Md5ofmessageattributes interface{} `json:"MD5OfMessageAttributes,omitempty"`
}

// StartMessageMoveTaskResult represents the StartMessageMoveTaskResult schema from the OpenAPI specification
type StartMessageMoveTaskResult struct {
	Taskhandle interface{} `json:"TaskHandle,omitempty"`
}

// GetQueueUrlRequest represents the GetQueueUrlRequest schema from the OpenAPI specification
type GetQueueUrlRequest struct {
	Queuename interface{} `json:"QueueName"`
	Queueownerawsaccountid interface{} `json:"QueueOwnerAWSAccountId,omitempty"`
}

// ChangeMessageVisibilityBatchRequestEntry represents the ChangeMessageVisibilityBatchRequestEntry schema from the OpenAPI specification
type ChangeMessageVisibilityBatchRequestEntry struct {
	Id interface{} `json:"Id"`
	Receipthandle interface{} `json:"ReceiptHandle"`
	Visibilitytimeout interface{} `json:"VisibilityTimeout,omitempty"`
}

// ListMessageMoveTasksResult represents the ListMessageMoveTasksResult schema from the OpenAPI specification
type ListMessageMoveTasksResult struct {
	Results interface{} `json:"Results,omitempty"`
}

// QueueAttributeMap represents the QueueAttributeMap schema from the OpenAPI specification
type QueueAttributeMap struct {
}

// RemovePermissionRequest represents the RemovePermissionRequest schema from the OpenAPI specification
type RemovePermissionRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
	Label interface{} `json:"Label"`
}

// SendMessageBatchRequest represents the SendMessageBatchRequest schema from the OpenAPI specification
type SendMessageBatchRequest struct {
	Entries interface{} `json:"Entries"`
	Queueurl interface{} `json:"QueueUrl"`
}

// StartMessageMoveTaskRequest represents the StartMessageMoveTaskRequest schema from the OpenAPI specification
type StartMessageMoveTaskRequest struct {
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Maxnumberofmessagespersecond interface{} `json:"MaxNumberOfMessagesPerSecond,omitempty"`
	Sourcearn interface{} `json:"SourceArn"`
}

// ReceiveMessageRequest represents the ReceiveMessageRequest schema from the OpenAPI specification
type ReceiveMessageRequest struct {
	Waittimeseconds interface{} `json:"WaitTimeSeconds,omitempty"`
	Attributenames interface{} `json:"AttributeNames,omitempty"`
	Maxnumberofmessages interface{} `json:"MaxNumberOfMessages,omitempty"`
	Messageattributenames interface{} `json:"MessageAttributeNames,omitempty"`
	Queueurl interface{} `json:"QueueUrl"`
	Receiverequestattemptid interface{} `json:"ReceiveRequestAttemptId,omitempty"`
	Visibilitytimeout interface{} `json:"VisibilityTimeout,omitempty"`
}

// SetQueueAttributesRequest represents the SetQueueAttributesRequest schema from the OpenAPI specification
type SetQueueAttributesRequest struct {
	Attributes interface{} `json:"Attributes"`
	Queueurl interface{} `json:"QueueUrl"`
}

// TagQueueRequest represents the TagQueueRequest schema from the OpenAPI specification
type TagQueueRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
	Tags interface{} `json:"Tags"`
}

// CreateQueueRequest represents the CreateQueueRequest schema from the OpenAPI specification
type CreateQueueRequest struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Queuename interface{} `json:"QueueName"`
	Tags interface{} `json:"tags,omitempty"`
}

// GetQueueAttributesRequest represents the GetQueueAttributesRequest schema from the OpenAPI specification
type GetQueueAttributesRequest struct {
	Attributenames interface{} `json:"AttributeNames,omitempty"`
	Queueurl interface{} `json:"QueueUrl"`
}

// ReceiveMessageResult represents the ReceiveMessageResult schema from the OpenAPI specification
type ReceiveMessageResult struct {
	Messages interface{} `json:"Messages,omitempty"`
}

// Message represents the Message schema from the OpenAPI specification
type Message struct {
	Receipthandle interface{} `json:"ReceiptHandle,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Body interface{} `json:"Body,omitempty"`
	Md5ofbody interface{} `json:"MD5OfBody,omitempty"`
	Md5ofmessageattributes interface{} `json:"MD5OfMessageAttributes,omitempty"`
	Messageattributes interface{} `json:"MessageAttributes,omitempty"`
	Messageid interface{} `json:"MessageId,omitempty"`
}

// ListQueuesResult represents the ListQueuesResult schema from the OpenAPI specification
type ListQueuesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Queueurls interface{} `json:"QueueUrls,omitempty"`
}

// ChangeMessageVisibilityBatchResultEntry represents the ChangeMessageVisibilityBatchResultEntry schema from the OpenAPI specification
type ChangeMessageVisibilityBatchResultEntry struct {
	Id interface{} `json:"Id"`
}

// ListQueueTagsRequest represents the ListQueueTagsRequest schema from the OpenAPI specification
type ListQueueTagsRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
}

// GetQueueUrlResult represents the GetQueueUrlResult schema from the OpenAPI specification
type GetQueueUrlResult struct {
	Queueurl interface{} `json:"QueueUrl,omitempty"`
}

// DeleteMessageBatchRequestEntry represents the DeleteMessageBatchRequestEntry schema from the OpenAPI specification
type DeleteMessageBatchRequestEntry struct {
	Id interface{} `json:"Id"`
	Receipthandle interface{} `json:"ReceiptHandle"`
}

// UntagQueueRequest represents the UntagQueueRequest schema from the OpenAPI specification
type UntagQueueRequest struct {
	Queueurl interface{} `json:"QueueUrl"`
	Tagkeys interface{} `json:"TagKeys"`
}

// ListMessageMoveTasksResultEntry represents the ListMessageMoveTasksResultEntry schema from the OpenAPI specification
type ListMessageMoveTasksResultEntry struct {
	Approximatenumberofmessagesmoved interface{} `json:"ApproximateNumberOfMessagesMoved,omitempty"`
	Approximatenumberofmessagestomove interface{} `json:"ApproximateNumberOfMessagesToMove,omitempty"`
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Sourcearn interface{} `json:"SourceArn,omitempty"`
	Startedtimestamp interface{} `json:"StartedTimestamp,omitempty"`
	Taskhandle interface{} `json:"TaskHandle,omitempty"`
	Maxnumberofmessagespersecond interface{} `json:"MaxNumberOfMessagesPerSecond,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// SendMessageRequest represents the SendMessageRequest schema from the OpenAPI specification
type SendMessageRequest struct {
	Messagededuplicationid interface{} `json:"MessageDeduplicationId,omitempty"`
	Messagegroupid interface{} `json:"MessageGroupId,omitempty"`
	Messagesystemattributes interface{} `json:"MessageSystemAttributes,omitempty"`
	Queueurl interface{} `json:"QueueUrl"`
	Delayseconds interface{} `json:"DelaySeconds,omitempty"`
	Messageattributes interface{} `json:"MessageAttributes,omitempty"`
	Messagebody interface{} `json:"MessageBody"`
}

// BatchResultErrorEntry represents the BatchResultErrorEntry schema from the OpenAPI specification
type BatchResultErrorEntry struct {
	Senderfault interface{} `json:"SenderFault"`
	Code interface{} `json:"Code"`
	Id interface{} `json:"Id"`
	Message interface{} `json:"Message,omitempty"`
}

// MessageBodySystemAttributeMap represents the MessageBodySystemAttributeMap schema from the OpenAPI specification
type MessageBodySystemAttributeMap struct {
}

// ListQueuesRequest represents the ListQueuesRequest schema from the OpenAPI specification
type ListQueuesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Queuenameprefix interface{} `json:"QueueNamePrefix,omitempty"`
}

// MessageAttributeValue represents the MessageAttributeValue schema from the OpenAPI specification
type MessageAttributeValue struct {
	Binarylistvalues interface{} `json:"BinaryListValues,omitempty"`
	Binaryvalue interface{} `json:"BinaryValue,omitempty"`
	Datatype interface{} `json:"DataType"`
	Stringlistvalues interface{} `json:"StringListValues,omitempty"`
	Stringvalue interface{} `json:"StringValue,omitempty"`
}

// DeleteMessageBatchResult represents the DeleteMessageBatchResult schema from the OpenAPI specification
type DeleteMessageBatchResult struct {
	Failed interface{} `json:"Failed"`
	Successful interface{} `json:"Successful"`
}

// SendMessageBatchResultEntry represents the SendMessageBatchResultEntry schema from the OpenAPI specification
type SendMessageBatchResultEntry struct {
	Md5ofmessageattributes interface{} `json:"MD5OfMessageAttributes,omitempty"`
	Md5ofmessagebody interface{} `json:"MD5OfMessageBody"`
	Md5ofmessagesystemattributes interface{} `json:"MD5OfMessageSystemAttributes,omitempty"`
	Messageid interface{} `json:"MessageId"`
	Sequencenumber interface{} `json:"SequenceNumber,omitempty"`
	Id interface{} `json:"Id"`
}
