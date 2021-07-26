package util

import "testing"

func TestGetOutBoundIP(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "测试获取本机IP",
			want:    "192.168.50.119",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOutBoundIP()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOutBoundIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetOutBoundIP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
