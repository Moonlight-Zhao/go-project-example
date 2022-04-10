package service

import (
	"testing"
)

func TestQueryPageInfo(t *testing.T) {
	type args struct {
		topicId int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "查询页面",
			args: args{
				topicId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := QueryPageInfo(tt.args.topicId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryPageInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
