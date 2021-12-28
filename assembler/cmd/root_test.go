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

func Test_assemble(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		wantExe []byte
		wantErr bool
	}{
		{
			"bad argument",
			"<.  FAIL",
			nil,
			true,
		},
		{
			"invalid cmd",
			"CMD 0x00",
			nil,
			true,
		},
		{
			"undefined labeled",
			"< [Test_Label]",
			nil,
			true,
		},
		{
			"valid code",
			`<.    0x5
<.   0x15
+.
<  0x20
>`,
			//     <.    x5   <.    x15   +.   <
			[]byte{0x10, 0x5, 0x10, 0x15, 0x0, 0x13,
				//x20 (1)   (2)   (3)   (4)   (5)   (6)   (7)
				0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				//>
				0x17},
			false,
		},
		{
			"named constants",
			`\testA = 0x10
\testB = 0x3c

<  \testA
<  #\testA
+
<  0x15
+
<  \testB
>`,
			//     <     x10   (1)   (2)   (3)   (4)   (5)   (6)   (7)
			[]byte{0x13, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				//<#  x10   (1)   (2)   (3)   (4)   (5)   (6)   (7)
				0x48, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				//+   <     x15   (1)   (2)   (3)   (4)   (5)   (6)   (7)
				0x03, 0x13, 0x15, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				//+   <     x3c   (1)   (2)   (3)   (4)   (5)   (6)   (7)
				0x03, 0x13, 0x3c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				//>
				0x17,
			},
			false,
		},
		{
			"valid code with labels",
			`[Loop]
<. 0x05
<. 0x20
+.
<. 0xff
[Inner_Loop]
	--.
< [Inner_Loop]
?>.
< [Loop]
|>`,
			//     <.    x10   <.    x20   +.    <.    xff
			[]byte{0x10, 0x05, 0x10, 0x20, 0x00, 0x10, 0xff,
				//--. <     [INNER_LOOP](2)   (3)   (4)   (5)   (6)   (7)   ?>.
				0x3d, 0x13, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2c,
				//<   [LOOP](1)   (2)   (3)   (4)   (5)   (6)   (7)   |>
				0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x34,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := assemble(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("assemble| wanted error %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(got, tt.wantExe) {
				t.Errorf("assemble| compiled code was wrong, wanted:\n%v\n but got:\n%v\n", tt.wantExe, got)
			}
		})
	}
}
