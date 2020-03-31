package auth

import (
	"testing"
	"time"
)

func TestNewToken(t *testing.T) {
	type args struct {
		credentials    *Credentials
		expirationTime time.Time
		secretKey      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewToken(tt.args.credentials, tt.args.expirationTime, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeToken(t *testing.T) {
	type args struct {
		strToken  string
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeToken(tt.args.strToken, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
