package cmd

import (
	"reflect"
	"testing"
)

func Test_convertNum(t *testing.T) {
	tests := []struct {
		name      string
		strNum    string
		size      byte
		expectNum []byte
		wantErr   bool
	}{
		{
			"decimal",
			"10",
			8,
			[]byte{10},
			false,
		},
		{
			"binary",
			"0b11",
			8,
			[]byte{0b11},
			false,
		},
		{
			"hexidecimal",
			"0x12",
			8,
			[]byte{0x12},
			false,
		},
		{
			"invalid number",
			"XYZ",
			8,
			nil,
			true,
		},
		{
			"overflow",
			"9999999999",
			8,
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertNum(tt.strNum, tt.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertNum| expected error: %v but got: %v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(got, tt.expectNum) {
				t.Errorf("convertNum| expected number %v but got: %v", tt.expectNum, got)
			}
		})
	}
}
