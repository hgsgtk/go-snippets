package main

import (
	"testing"
)

func TestValidateLength(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid minimum length",
			length:  8,
			wantErr: false,
		},
		{
			name:    "valid maximum length",
			length:  100,
			wantErr: false,
		},
		{
			name:    "valid middle length",
			length:  50,
			wantErr: false,
		},
		{
			name:    "too short - below minimum",
			length:  7,
			wantErr: true,
			errMsg:  "length must be at least 8 characters",
		},
		{
			name:    "too short - zero",
			length:  0,
			wantErr: true,
			errMsg:  "length must be at least 8 characters",
		},
		{
			name:    "too short - negative",
			length:  -1,
			wantErr: true,
			errMsg:  "length must be at least 8 characters",
		},
		{
			name:    "too long - above maximum",
			length:  101,
			wantErr: true,
			errMsg:  "length must be at most 100 characters",
		},
		{
			name:    "too long - much above maximum",
			length:  1000,
			wantErr: true,
			errMsg:  "length must be at most 100 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateLength(tt.length)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("ValidateLength() expected error but got none")
					return
				}
				if err.Error() != tt.errMsg {
					t.Errorf("ValidateLength() error = %v, want %v", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateLength() unexpected error = %v", err)
				}
			}
		})
	}
}
