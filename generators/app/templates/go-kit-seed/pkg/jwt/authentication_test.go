package jwt

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc/metadata"
)

func createContext() context.Context {
	md := metadata.Pairs(
		"authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI3ODBlYmNlYS1jNDk4LTRhOWQtODU1ZC05Y2E3NGY2MGFhNmYiLCJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSJ9.wWknYyX1DoA0GReFcmlVgLAhMkpYrqVvXUvuu6lCSKg",
	)

	return metadata.NewIncomingContext(context.Background(), md)
}

func Test_extractPayLoadFromJWT(t *testing.T) {
	type args struct {
		jwt string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			"Successfully extracted payload",
			args{
				jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSIsImZpcnN0TmFtZSI6IkFuZHJpdXMiLCJsYXN0TmFtZSI6Ik1venVyYWl0aXMiLCJhZG1pbiI6dHJ1ZSwidXVpZCI6ImRhODhjYzViLWU1ZmYtNGQxYS1iZDc2LWIwNDQwZjNhYTJiNiJ9.c89jokXVMcdpU_gwXc0eZbk_U1WHckYyIRF1DIAHDYs",
			},
			map[string]interface{}{
				"email":     "andriusmozuraitis@gmail.com",
				"firstName": "Andrius",
				"lastName":  "Mozuraitis",
				"admin":     true,
				"uuid":      "da88cc5b-e5ff-4d1a-bd76-b0440f3aa2b6",
			},
			false,
		},
		{
			"Expect error: Malformed jwt",
			args{
				jwt: "eyJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSIsImZpcnN0TmFtZSI6IkFuZHJpdXMiLCJsYXN0TmFtZSI6Ik1venVyYWl0aXMiLCJhZG1pbiI6dHJ1ZSwidXVpZCI6ImRhODhjYzViLWU1ZmYtNGQxYS1iZDc2LWIwNDQwZjNhYTJiNiJ9.c89jokXVMcdpU_gwXc0eZbk_U1WHckYyIRF1DIAHDYs",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractPayLoadFromJWT(tt.args.jwt)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractPayLoadFromJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractPayLoadFromJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Auth(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    Authentication
		wantErr bool
	}{
		{
			"Successfuly get user payload Correct metadata",
			args{
				createContext(),
			},
			Authentication{
				UserID: "780ebcea-c498-4a9d-855d-9ca74f60aa6f",
				Email:  "andriusmozuraitis@gmail.com",
				JWT:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI3ODBlYmNlYS1jNDk4LTRhOWQtODU1ZC05Y2E3NGY2MGFhNmYiLCJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSJ9.wWknYyX1DoA0GReFcmlVgLAhMkpYrqVvXUvuu6lCSKg",
				Payload: map[string]interface{}{
					"userId": "780ebcea-c498-4a9d-855d-9ca74f60aa6f",
					"email":  "andriusmozuraitis@gmail.com",
				},
			},
			false,
		},
		{
			"Fail no userId property",
			args{
				func() context.Context {
					md := metadata.Pairs(
						"Authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSJ9.nkxGsIYjGNBsYdb63Mv2wYAM9D0BU82ziJ_eOZCVgoI",
					)

					return metadata.NewIncomingContext(context.Background(), md)
				}(),
			},
			Authentication{},
			true,
		},
		{
			"Fail no email property",
			args{
				func() context.Context {
					md := metadata.Pairs(
						"Authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI3ODBlYmNlYS1jNDk4LTRhOWQtODU1ZC05Y2E3NGY2MGFhNmYifQ.2_08VC4hB6mXLb_ZVmufQIbMTzWk25n2moEo3zQnB8A",
					)

					return metadata.NewIncomingContext(context.Background(), md)
				}(),
			},
			Authentication{},
			true,
		},
		{
			"Fail incorect userId",
			args{
				func() context.Context {
					md := metadata.Pairs(
						"Authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJhc2QtYzQ5OC1hc2QtODU1ZC1hc2QiLCJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSJ9.qwK_Edembznx_pplSGYbFwiDyGDnDdFpKrcmcBQ8eWs",
					)

					return metadata.NewIncomingContext(context.Background(), md)
				}(),
			},
			Authentication{},
			true,
		},
		{
			"Fail incorect email",
			args{
				func() context.Context {
					md := metadata.Pairs(
						"Authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI3ODBlYmNlYS1jNDk4LTRhOWQtODU1ZC05Y2E3NGY2MGFhNmYiLCJlbWFpbCI6InNhZGFzZHNhZGFzZGFzZGFzZCJ9.U6IOY34rlC_6m_orN2iDVkAY0IaUpsO_aVTaKub-noo",
					)

					return metadata.NewIncomingContext(context.Background(), md)
				}(),
			},
			Authentication{},
			true,
		},
		{
			"No Authentication meta key",
			args{
				func() context.Context {
					return context.Background()
				}(),
			},
			Authentication{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Auth(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}

var bContext = createContext()

func BenchmarkAuth(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		_, err := Auth(bContext)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}
