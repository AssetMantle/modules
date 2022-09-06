package issue

import (
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/helpers"
	"reflect"
	"testing"
)

func Test_newTransactionResponse(t *testing.T) {
	type args struct {
		error error
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionResponse
	}{
		// TODO: Add test cases.
		{"+ve", args{nil}, transactionResponse{Success: true, Error: nil}},
		{"-ve", args{errors.IncorrectMessage}, transactionResponse{Success: false, Error: errors.IncorrectMessage}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionResponse(tt.args.error); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransactionResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionResponse_GetError(t *testing.T) {
	type fields struct {
		Success bool
		Error   error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"-ve", fields{Success: false, Error: errors.IncorrectMessage}, true},
		{"+ve", fields{Success: true, Error: nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionResponse := transactionResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
			}
			if err := transactionResponse.GetError(); (err != nil) != tt.wantErr {
				t.Errorf("GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_transactionResponse_IsSuccessful(t *testing.T) {
	type fields struct {
		Success bool
		Error   error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"-ve", fields{Success: false, Error: errors.IncorrectMessage}, false},
		{"+ve", fields{Success: true, Error: nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionResponse := transactionResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
			}
			if got := transactionResponse.IsSuccessful(); got != tt.want {
				t.Errorf("IsSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}
