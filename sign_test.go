package uniSign

import "testing"

func Test_getParamsString(t *testing.T) {
	type args struct {
		params map[string]any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "int",
			args: args{
				params: map[string]any{
					"intParam": 1,
				},
			},
			want: "intParam=1",
		},
		{
			name: "int64",
			args: args{
				params: map[string]any{
					"int64Param": 1,
				},
			},
			want: "int64Param=1",
		},
		{
			name: "int32",
			args: args{
				params: map[string]any{
					"int32Param": 1,
				},
			},
			want: "int32Param=1",
		},
		{
			name: "int16",
			args: args{
				params: map[string]any{
					"int16Param": 1,
				},
			},
			want: "int16Param=1",
		},
		{
			name: "int8",
			args: args{
				params: map[string]any{
					"int8Param": 1,
				},
			},
			want: "int8Param=1",
		},
		{
			name: "uint",
			args: args{
				params: map[string]any{
					"uintParam": 1,
				},
			},
			want: "uintParam=1",
		},
		{
			name: "uint64",
			args: args{
				params: map[string]any{
					"uint64Param": 1,
				},
			},
			want: "uint64Param=1",
		},
		{
			name: "uint32",
			args: args{
				params: map[string]any{
					"uint32Param": 1,
				},
			},
			want: "uint32Param=1",
		},
		{
			name: "uint16",
			args: args{
				params: map[string]any{
					"uint16Param": 1,
				},
			},
			want: "uint16Param=1",
		},
		{
			name: "uint8",
			args: args{
				params: map[string]any{
					"uint8Param": 1,
				},
			},
			want: "uint8Param=1",
		},
		{
			name: "float32",
			args: args{
				params: map[string]any{
					"float32Param": 1,
				},
			},
			want: "float32Param=1",
		},
		{
			name: "float64",
			args: args{
				params: map[string]any{
					"float64Param": 1,
				},
			},
			want: "float64Param=1",
		},
		{
			name: "bool",
			args: args{
				params: map[string]any{
					"boolParam": true,
				},
			},
			want: "boolParam=true",
		},
		{
			name: "string",
			args: args{
				params: map[string]any{
					"stringParam": "test",
				},
			},
			want: "stringParam=test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tool := Tool{AuthSecretKey: "a"}
			if got := tool.getParamsString(tt.args.params); got != tt.want {
				t.Errorf("getParamsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
