// Package connector provides the core interface and types for cc-connect,
// a unified interface for connecting to various AI model providers.
package connector

import (
	"context"
	"io"
)

// Role represents the role of a message in a conversation.
type Role string

const (
	// RoleUser represents a message from the user.
	RoleUser Role = "user"
	// RoleAssistant represents a message from the AI assistant.
	RoleAssistant Role = "assistant"
	// RoleSystem represents a system-level instruction message.
	RoleSystem Role = "system"
)

// Message represents a single message in a conversation.
type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// Request holds the parameters for a completion request.
type Request struct {
	// Model is the identifier of the model to use.
	Model string `json:"model"`
	// Messages is the conversation history.
	Messages []Message `json:"messages"`
	// MaxTokens limits the number of tokens in the response.
	MaxTokens int `json:"max_tokens,omitempty"`
	// Temperature controls the randomness of the output (0.0 - 2.0).
	Temperature float64 `json:"temperature,omitempty"`
	// Stream indicates whether to stream the response.
	Stream bool `json:"stream,omitempty"`
	// ExtraParams allows provider-specific parameters.
	ExtraParams map[string]any `json:"extra_params,omitempty"`
}

// Response holds the result of a completion request.
type Response struct {
	// ID is the unique identifier for this response.
	ID string `json:"id"`
	// Model is the model that generated the response.
	Model string `json:"model"`
	// Choices contains the generated completion choices.
	Choices []Choice `json:"choices"`
	// Usage contains token usage statistics.
	Usage Usage `json:"usage"`
}

// Choice represents a single completion choice.
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Usage contains token usage statistics for a request.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// StreamChunk represents a single chunk in a streaming response.
type StreamChunk struct {
	ID      string        `json:"id"`
	Model   string        `json:"model"`
	Choices []StreamChoice `json:"choices"`
}

// StreamChoice represents a single choice in a streaming chunk.
type StreamChoice struct {
	Index int    `json:"index"`
	Delta Delta  `json:"delta"`
	FinishReason string `json:"finish_reason"`
}

// Delta holds the incremental content in a streaming response.
type Delta struct {
	Role    Role   `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

// Connector defines the interface that all AI provider connectors must implement.
type Connector interface {
	// Complete sends a completion request and returns the full response.
	Complete(ctx context.Context, req *Request) (*Response, error)

	// Stream sends a completion request and streams the response chunks
	// to the provided writer.
	Stream(ctx context.Context, req *Request, w io.Writer) error

	// Name returns the identifier of the provider (e.g., "openai", "anthropic").
	Name() string

	// ListModels returns the list of available models for this provider.
	ListModels(ctx context.Context) ([]string, error)
}
