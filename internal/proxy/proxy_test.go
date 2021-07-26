package proxy

import "testing"

// curl --proxy "socks5://127.0.0.1:1080" http://pipe.oook.fun
func TestProxyServer_Serve(t *testing.T) {
	type fields struct {
		port uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "test 1",
			fields:  fields{port: 1080},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProxyServer(tt.fields.port)
			if err := p.Serve(); (err != nil) != tt.wantErr {
				t.Errorf("ProxyServer.Serve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
