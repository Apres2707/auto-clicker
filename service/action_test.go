package service

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestAction_FillActions(t *testing.T) {
	tests := []struct {
		name string
		prepareCtxFn func () (context.Context, func())
		wantCount int
	}{
		{
			name: "success full fill",
			prepareCtxFn: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			wantCount: 10,
		},
		{
			name: "success fill two action",
			prepareCtxFn: func() (context.Context, func()) {
				ctx, cancel := context.WithTimeout(context.Background(), 11 * time.Second)

				return ctx, cancel
			},
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Action{}
			ctx, cancel := tt.prepareCtxFn()
			defer cancel()
			if got := c.FillActions(ctx); !reflect.DeepEqual(len(got), tt.wantCount) {
				t.Errorf("FillActions() got count = %v, want count %v", len(got), tt.wantCount)
			}
		})
	}
}
