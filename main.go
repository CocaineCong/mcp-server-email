package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Email Sender",
		"1.0.0",
	)
	// Add tool
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("email",
			mcp.Required(),
			mcp.Description("The email address to send"),
		),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("The email content to send"),
		),
	)
	// Add tool handler
	s.AddTool(tool, emailHandler)
	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func emailHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	email, ok := request.Params.Arguments["email"].(string)
	if !ok {
		return nil, errors.New("email must be a string")
	}
	content, ok := request.Params.Arguments["content"].(string)
	if !ok {
		return nil, errors.New("content must be a string")
	}
	emailSender := NewEmailSender()
	err := emailSender.Send(email, content, content)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(fmt.Sprintf("Send %s with content about %s Successfully", email, content)), nil
}
