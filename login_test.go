package main

import (
	"context"
	"testing"
)

func TestLYExamination_Login(t *testing.T) {
	type args struct {
		in0 context.Context
		in  *LoginReq
	}
	tests := []struct {
		name    string
		args    args
		want    *LoginRsp
		wantErr bool
	}{
		{
			args: args{
				in0: nil,
				in: &LoginReq{
					Phone: "Phone",
					Pwd:   "Pwd",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LYExamination{}
			got, err := l.Login(tt.args.in0, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Login() got = %v", got)
		})
	}
}
