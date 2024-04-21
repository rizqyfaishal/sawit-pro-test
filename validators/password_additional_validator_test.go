package validators

import (
	"testing"
)

func Test_validateAtLeastXCapitalChar(t *testing.T) {
	type args struct {
		str           string
		minimumNumber int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "When the words did not have minimum capital characters, then return false",
			args: args{
				str:           "rizqy faishal tanjung",
				minimumNumber: 1,
			},
			want: false,
		},
		{
			name: "When the words did not have minimum capital characters, then return false",
			args: args{
				str:           "Rizqy faishal tanjung",
				minimumNumber: 1,
			},
			want: true,
		},
		{
			name: "When the words did not have minimum capital characters, then return false",
			args: args{
				str:           "rizqy faishal tanjung",
				minimumNumber: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateAtLeastXCapitalChar(tt.args.str, tt.args.minimumNumber); got != tt.want {
				t.Errorf("validateAtLeastXCapitalChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateAtLeastXSpecialChar(t *testing.T) {
	type args struct {
		str           string
		minimumNumber int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "When the words did not have minimum capital characters, then return false",
			args: args{
				str:           "rizqy faishal tanjung",
				minimumNumber: 1,
			},
			want: false,
		},
		{
			name: "When the words did not have minimum capital characters, then return false",
			args: args{
				str:           "Rizqy faishal tanjun#g",
				minimumNumber: 1,
			},
			want: true,
		},
		{
			name: "When the words did not have minimum capital characters, then return false",
			args: args{
				str:           "Rizqy faishal tanjung#",
				minimumNumber: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateAtLeastXSpecialChar(tt.args.str, tt.args.minimumNumber); got != tt.want {
				t.Errorf("validateAtLeastXSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
