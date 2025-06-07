package crypto

import (
	"os"
	"testing"
)

func TestSign(t *testing.T) {
	key, _ := os.ReadFile("paperless_private_key.pem")
	msg := []byte("hello")

	type args struct {
		message []byte
		privKey []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{name: "test", args: args{message: msg, privKey: key}, want: "LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=", want1: "MEYCIQDOnqdo8TsP6gAlzOO4VVA3MPwDziYPa9t3PAEM5rYOrwIhAMWRLYPjVxRHDDNmkh4//siSVrBM4qrF/+x3uN6Tje9o", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Sign(tt.args.message, tt.args.privKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sign() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Sign() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
