package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"reflect"
	"testing"
)

func TestNewID(t *testing.T) {
	type args struct {
		idString string
	}
	tests := []struct {
		name string
		args args
		want ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", args{"ID"}, NewID("ID")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewID(tt.args.idString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    id
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", args{NewID("ID")}, id{IDString: "ID"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := idFromInterface(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("idFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_id_Bytes(t *testing.T) {
	type fields struct {
		IDString string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{"ID"}, []byte("ID")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := id{
				IDString: tt.fields.IDString,
			}
			if got := id.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_id_Compare(t *testing.T) {
	type fields struct {
		IDString string
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"+ve", fields{"ID"}, args{NewID("ID")}, 0},
		// TODO: It Should fail
		{"-ve", fields{"ID"}, args{NewID("ID2")}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := id{
				IDString: tt.fields.IDString,
			}
			if got := id.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_id_String(t *testing.T) {
	type fields struct {
		IDString string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{"ID"}, "ID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := id{
				IDString: tt.fields.IDString,
			}
			if got := id.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
