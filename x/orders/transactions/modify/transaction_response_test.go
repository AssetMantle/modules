// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
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
		{"+ve", args{nil}, transactionResponse{true, nil}},
		{"-ve", args{errorConstants.IncorrectFormat}, transactionResponse{false, errorConstants.IncorrectFormat}}}
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
		{"+ve", fields{Success: true, Error: nil}, false},
		{"-ve", fields{Error: errorConstants.IncorrectFormat}, true},
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
		{"+ve", fields{Success: true, Error: nil}, true},
		{"-ve", fields{Error: errorConstants.IncorrectFormat}, false},
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
