/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

// TokenUsage is the token usage for the model.
type TokenUsage struct {
	// PromptTokens is the number of prompt tokens.
	PromptTokens int
	// CompletionTokens is the number of completion tokens.
	CompletionTokens int
	// TotalTokens is the total number of tokens.
	TotalTokens int
}

// Config is the config for the model.
type Config struct {
	// Model is the model name.
	Model string
	// MaxTokens is the max number of tokens, if reached the max tokens, the model will stop generating, and mostly return an finish reason of "length".
	MaxTokens int
	// Temperature is the temperature, which controls the randomness of the model.
	Temperature float32
	// TopP is the top p, which controls the diversity of the model.
	TopP float32
	// Stop is the stop words, which controls the stopping condition of the model.
	Stop []string
}

// CallbackInput is the input for the model callback.
type CallbackInput struct {
	// Messages is the messages to be sent to the model.
	Messages []*schema.Message
	// Tools is the tools to be used in the model.
	Tools []*schema.ToolInfo
	// Config is the config for the model.
	Config *Config
	// Extra is the extra information for the callback.
	Extra map[string]any
}

// CallbackOutput is the output for the model callback.
type CallbackOutput struct {
	// Message is the message generated by the model.
	Message *schema.Message
	// Config is the config for the model.
	Config *Config
	// TokenUsage is the token usage of this request.
	TokenUsage *TokenUsage
	// Extra is the extra information for the callback.
	Extra map[string]any
}

// ConvCallbackInput converts the callback input to the model callback input.
func ConvCallbackInput(src callbacks.CallbackInput) *CallbackInput {
	switch t := src.(type) {
	case *CallbackInput: // when callback is triggered within component implementation, the input is usually already a typed *model.CallbackInput
		return t
	case []*schema.Message: // when callback is injected by graph node, not the component implementation itself, the input is the input of Chat Model interface, which is []*schema.Message
		return &CallbackInput{
			Messages: t,
		}
	default:
		return nil
	}
}

// ConvCallbackOutput converts the callback output to the model callback output.
func ConvCallbackOutput(src callbacks.CallbackOutput) *CallbackOutput {
	switch t := src.(type) {
	case *CallbackOutput: // when callback is triggered within component implementation, the output is usually already a typed *model.CallbackOutput
		return t
	case *schema.Message: // when callback is injected by graph node, not the component implementation itself, the output is the output of Chat Model interface, which is *schema.Message
		return &CallbackOutput{
			Message: t,
		}
	default:
		return nil
	}
}
