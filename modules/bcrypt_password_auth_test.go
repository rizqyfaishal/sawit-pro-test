package modules

import "testing"

func TestBcryptPasswordAuth_CompareHashedPassword(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "When given correct hashed password and the real passowrd, it will return true",
			args: args{
				hashedPassword: "$2a$10$3CnCHHqLsCb9R7.WEiG7yOwIyQdFtComVcNsOoM9Ns5mb/L03if0i",
				password:       "Asdasd123#",
			},
			want:    true,
			wantErr: false,
		},

		{
			name: "When given incorrect hashed password and the real passowrd, it will return false",
			args: args{
				hashedPassword: "$2a$10$3CnCHHqLsCb9R7.WEiG7yOwIyQdFtComVcNsOoM9Ns5mb/L03if0i",
				password:       "Asdasd125#",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BcryptPasswordAuth{}
			got, err := b.CompareHashedPassword(tt.args.hashedPassword, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareHashedPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompareHashedPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBcryptPasswordAuth_GenerateHashedPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "When given correct string input, it will return hashed string",
			args: args{
				password: "asdasd123",
			},
			wantErr: false,
		},

		{
			name: "When given incorrect string input (because its too long), it will return error",
			args: args{
				password: "asdasd123asdasd123asdasd123asdasd123asdasd123asdasd123asdasd123asdasd123asdasd123asdasd123asdasd123asdasd123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BcryptPasswordAuth{}
			_, err := b.GenerateHashedPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHashedPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
