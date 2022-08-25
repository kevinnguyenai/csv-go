package pkg

import (
	"net/http"
	"testing"
)

func TestRequestor_create(t *testing.T) {
	type fields struct {
		option_fsyms string
		option_fsymf string
		option_app   string
		request      *http.Request
	}
	type args struct {
		url      string
		protocol string
		method   string
		token    string
		body     byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &Requestor{
				option_fsyms: tt.fields.option_fsyms,
				option_fsymf: tt.fields.option_fsymf,
				option_app:   tt.fields.option_app,
				request:      tt.fields.request,
			}
			if err := api.Create(tt.args.url, tt.args.protocol, tt.args.method, tt.args.token, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("Requestor.create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
