//go:build unit

package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestNormalizeDeepSeekV4OpenAIReasoningEffort(t *testing.T) {
	tests := []struct {
		name, model, body, want string
		changed                 bool
	}{
		{"medium to high", "DeepSeek-V4.1-Flash", `{"reasoning_effort":"medium"}`, "high", true},
		{"minimal to low", "deepseek-v4-pro", `{"reasoning_effort":"minimal"}`, "low", true},
		{"prefixed flash alias", "deepseek/deepseek-flash", `{"reasoning_effort":"Medium"}`, "high", true},
		{"accepted value kept", "DeepSeek-V4.1-Flash", `{"reasoning_effort":"xhigh"}`, "xhigh", false},
		{"none kept", "DeepSeek-V4.1-Flash", `{"reasoning_effort":"none"}`, "none", false},
		{"absent stays absent", "DeepSeek-V4.1-Flash", `{}`, "", false},
		{"other model untouched", "gpt-5.5", `{"reasoning_effort":"medium"}`, "medium", false},
		{"older deepseek untouched", "deepseek-chat", `{"reasoning_effort":"medium"}`, "medium", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, changed := NormalizeDeepSeekV4OpenAIReasoningEffort([]byte(tt.body), tt.model)
			require.Equal(t, tt.changed, changed)
			require.Equal(t, tt.want, gjson.GetBytes(got, "reasoning_effort").String())
		})
	}
}
