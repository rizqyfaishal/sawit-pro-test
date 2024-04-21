package modules

import (
	"github.com/golang-jwt/jwt/v5"
	"reflect"
	"testing"
	"time"
)

var validPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

var validPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEArSHQh7CvEA6hRqpjHzD9
BpvUpIMEjVbilFjquLV0DuuTc6+0U/mbyOtJSWULMKofi0rx/qP6Wi/0/IBdDE84
QuwXcVnXnuMzIXR0jM0IdF1GWpzxf6P7LuIny4zPdScnOlxXX9xx6AV43TPDf1qY
unHpSFBQCkfxS7r/8fgZ3F1rWyFJJtmHyRoNij5c0blNPugnZ8THVAFDYsmN8Sse
jheL7USp6TvrnMWQmHnS0/QpqaCHMOxgoklY2X6DYqpE9yr7729roESOqSU0lhFj
UXR8prl7Sz2aWjySrrBJtkVwnuWR3X8qSbnPX8Fkb+yfEdHgNvD7IRzExdPuSoqy
sD5NoQWZQpr88b7sjnpXz2SZNekpcGFuWO9biEP/QylcHgeLd/EdAExGiHArLuWK
dCxroDYfdygAIpv8yUwJfsIxoe+0h3KwZPw63D7DErOtooIBD3a0KAJLTKgxRTck
L8q7PfaWn3BUsb51vnOEoavXN2N84rmJgM6bv7TQW/emlEBVosh5cj4ZNfa5GNCV
MH5PKstrqOYVAGdjEuIMan4273RoaKQI3nEx7iAjwR10CerLbsf8aXEJ/dlfT7n+
XYi4oDPnrctUASOjWpY3m4rVDIZcBgVYBVVbVRpAkcVnWqxkUkndxjuGSvSUjiYw
2J86ZNU7o2amSgTO3HVPNlUCAwEAAQ==
-----END PUBLIC KEY-----
`)

var expectedExpiredTime = "2025-12-03 00:00:00"

func TestRS256Jwt_GenerateJwt(t *testing.T) {

	expiredTime, _ := time.Parse("2006-01-02 15:04:05", expectedExpiredTime)

	type fields struct {
		privateKey []byte
		publicKey  []byte
	}
	type args struct {
		claims CustomClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		{
			name: "When the private key is valid, then it will return valid jwt token",
			fields: fields{
				privateKey: validPrivateKey,
				publicKey:  validPublicKey,
			},
			args: args{
				struct {
					jwt.RegisteredClaims
					UserId int64
				}{RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "simple-app-service",
					ExpiresAt: jwt.NewNumericDate(expiredTime),
				}, UserId: 123},
			},
			wantErr: false,
		},
		{
			name: "When the private key is valid, then it will return valid jwt token",
			fields: fields{
				privateKey: []byte(`a_invalid_private_key`),
				publicKey:  []byte("a_invalid_public_key"),
			},
			args: args{
				struct {
					jwt.RegisteredClaims
					UserId int64
				}{RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "simple-app-service",
					ExpiresAt: jwt.NewNumericDate(expiredTime),
				}, UserId: 123},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RS256Jwt{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			_, err := r.GenerateJwt(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRS256Jwt_VerifyJwt(t *testing.T) {

	//expiredTime, _ := time.Parse("2006-01-02 15:04:05", expectedExpiredTime)

	type fields struct {
		privateKey []byte
		publicKey  []byte
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CustomClaims
		wantErr bool
	}{
		{
			name: "When it given invalid public key, it will return error caused by cannot parse the public key",
			fields: fields{
				privateKey: validPrivateKey,
				publicKey:  []byte("invalid_public_key"),
			},
			args: args{
				tokenString: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzaW1wbGUtYXBwLXNlcnZpY2UiLCJleHAiOjE3MTM2NzI1NjksIlVzZXJJZCI6MTIzfQ.czDMbp8JTv9N7tZPZlpIJ4LavPBk_kA43Kc4SDaiVjk-6K3a7KTu1f5A8RM3hIY7rRPsUtw5EflS7Of0xsm__V3FanRjXKl8PxgncWddG22SyeT6saI2DhgsgMbwvsQZQB6Ic5vHTt0gkmspzrrIQYnbFrsHt7SLVDuBiHyEjPryahuamSaDXKm2gIVqkJCOTVYwy16TxKJA52hLwY0aDJiwcHqo3cfym5md-CjAxKszzG88JCFF82m4nBkHXzBXyzRuPSWUOdkxPSklh_hFFNRAJ7CuBe6w86ww44t9R0c744YdUwj3mUiTWMvcPh_g81bB7-QEuV4wm-mFvxEeuZYFaKPghjQFWSRg2dsPhVH6762HpHr0Fq1DBBDsHf4aNzgrqsRY0PKfDHlncDJif0jljUi1v9VgnrYZHaHOdXE8VWyZtEOwReyT5DxM-AqyqM6EmSQ8aY_Dcpt2L5vkaU0yyqwP-afgmu3uUn4qpsyvRhnAA4lgYuKs-z0JF6swSoO29ATLrMl0iOBAZFbfMgdFd3Rd4v4mXcq82bsGNLZCYna1fJjoyzy_cOtpsEKpPP3tymvAvSjcAbrDbnTy4gBAbgPGgIl7fu3JFa2OZ6zqlTlmug_hlX-qk4aD90c7zQiAc63_ndCP_Nm17Ui2w4OkCo6tsfjhjdE3brJoh5s",
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "When it given valid public key, it will authorized with claims",
			fields: fields{
				privateKey: validPrivateKey,
				publicKey:  validPublicKey,
			},
			args: args{
				tokenString: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzaW1wbGUtYXBwLXNlcnZpY2UiLCJleHAiOjE3NjQ3MjAwMDAsIlVzZXJJZCI6MTIzfQ.Up1wGEPND_Y5T4ztSyVxekhsjZpDdwRIG0lhPfXyw3hB6cuH_C57fWOwPjpGN16sSofKsHDPCSSwgvUPDKZPQzYS76QoomxhOoyrUcxCebmwJrTbqxOVi3Awe68jN5AJLpiZdpPpQO4FIcaT78oC1k5P_KMScdreU2ysksmitlWCC64b5jE1nY5DnOkZc2-Hrm790UUwvOaU6Sbayx0cFt80o-7SQfIJSqc_ZlmwJp1W40c2gj2nov0WpIHiawdDIa07A16ON5fCSKihkS3SsbONiMTOnYg6JJSprV-TeDFY-rjDOTeujvp0PZAmVzrR8N3CuoSmZrQ0Q18Iuf4sN3pjaRUE6FA133Dmw12o4O_wEa3x0V6Jm85AwoOP1FnZ7rp_IkouTydPCTH0Ji7N3HxZy1bU_h6zR8UeweNrTCURPzCX5u0bDMo4m8yBsZ9SyrhAwl6rZbI50eVpEM0s0QaHWuc-aAbdtMM7OVh-pywNHXnfu8d3fM_v9epro8x-FvRVmQ46_r3xoq3h1iangiiSpmGCC3Y87hOVHZ7KTMlLDCwZ81bw5MkBE24jkPOn5hQuxg-kZPtDuL7dtJe5B3kjaTw5eVS7rtFWqvbKISnVJzNUf63RMRSW_LmO_c2zzZb57WCBsKGOsADbTTDDvHph8MLC3xOlC-I0eG5C108",
			},
			want: &CustomClaims{
				UserId: 123,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RS256Jwt{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			got, err := r.VerifyJwt(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyJwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyJwt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRS256Jwt(t *testing.T) {
	type args struct {
		privateKey []byte
		publicKey  []byte
	}
	tests := []struct {
		name string
		args args
		want JsonWebTokenUtilInterface
	}{
		{
			name: "When given valid public key, it will return the module",
			args: args{
				privateKey: validPrivateKey,
				publicKey:  validPublicKey,
			},
			want: RS256Jwt{
				privateKey: validPrivateKey,
				publicKey:  validPublicKey,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRS256Jwt(tt.args.privateKey, tt.args.publicKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRS256Jwt() = %v, want %v", got, tt.want)
			}
		})
	}
}
