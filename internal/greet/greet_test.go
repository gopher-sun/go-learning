package greet

import (
	"errors"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{
			name:    "正常输入",
			input:   "gopher-sun",
			want:    "Hello, gopher-sun! 欢迎来到 Go 世界。",
			wantErr: nil,
		},
		{
			name:    "空字符串应报错",
			input:   "",
			want:    "",
			wantErr: ErrEmptyName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hello(tt.input)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("期望错误 %v，实际 %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("不应报错，实际 %v", err)
			}
			if got != tt.want {
				t.Errorf("期望 %q，实际 %q", tt.want, got)
			}
		})
	}
}
