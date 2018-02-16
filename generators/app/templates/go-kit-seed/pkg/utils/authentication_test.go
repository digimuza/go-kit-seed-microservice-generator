package utils

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestGetJWTPayload(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    JWTPayload
		wantErr bool
	}{
		{
			"OK",
			args{
				func() context.Context {
					md := metadata.Pairs(
						"authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJlMTM1MzBkZC1kY2U3LTQ2M2ItOTk2MC00MjMzMTUwYmNjNTIiLCJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSJ9.0Kc4rDe8rYX4cffvef7PRh9Gba7yGSINjkWj4RUbs3w",
					)

					return metadata.NewIncomingContext(context.Background(), md)
				}(),
			},
			JWTPayload{
				UserID: "e13530dd-dce7-463b-9960-4233150bcc52",
				Email:  "andriusmozuraitis@gmail.com",
			},
			false,
		},
		{
			"No authentication",
			args{
				context.Background(),
			},
			JWTPayload{},
			true,
		},
		{
			"No authentication",
			args{
				context.Background(),
			},
			JWTPayload{},
			true,
		},
		{
			"Context as nill",
			args{
				nil,
			},
			JWTPayload{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJWTPayload(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJWTPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJWTPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJWTToken(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"OK",
			args{
				func() context.Context {
					md := metadata.Pairs(
						"authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJlMTM1MzBkZC1kY2U3LTQ2M2ItOTk2MC00MjMzMTUwYmNjNTIiLCJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSJ9.0Kc4rDe8rYX4cffvef7PRh9Gba7yGSINjkWj4RUbs3w",
					)

					return metadata.NewIncomingContext(context.Background(), md)
				}(),
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJlMTM1MzBkZC1kY2U3LTQ2M2ItOTk2MC00MjMzMTUwYmNjNTIiLCJlbWFpbCI6ImFuZHJpdXNtb3p1cmFpdGlzQGdtYWlsLmNvbSJ9.0Kc4rDe8rYX4cffvef7PRh9Gba7yGSINjkWj4RUbs3w",
			false,
		},
		{
			"No authentication",
			args{
				context.Background(),
			},
			"",
			true,
		},
		{
			"Context as nill",
			args{
				nil,
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJWTToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetJWTToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
