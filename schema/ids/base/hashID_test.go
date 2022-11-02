package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/ids"
)

func TestReadHashID(t *testing.T) {
	type args struct {
		hashIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.HashID
		wantErr bool
	}{
		{name: "empty string", args: args{hashIDString: ""}, want: PrototypeHashID(), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadHashID(tt.args.hashIDString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadHashID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
