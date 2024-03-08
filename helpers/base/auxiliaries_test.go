package base

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	randomUniqueIdentifiers = random.GenerateUniqueIdentifierList("", 2)
)

type mockAuxiliary struct {
	name string
}

func (m mockAuxiliary) GetName() string {
	return m.name
}

func (m mockAuxiliary) GetKeeper() helpers.AuxiliaryKeeper {
	// Stubbed return as it is not used in GetAuxiliary
	return nil
}

func (m mockAuxiliary) Initialize(helpers.Mapper, helpers.ParameterManager, ...interface{}) helpers.Auxiliary {
	// Stubbed return as it is not used in GetAuxiliary
	return nil
}

func TestGetAuxiliary(t *testing.T) {
	tests := []struct {
		name        string
		auxiliaries auxiliaries
		arg         string
		want        helpers.Auxiliary
	}{
		{
			name:        "Empty AuxiliaryList",
			auxiliaries: auxiliaries{},
			arg:         "auxiliary1",
			want:        nil,
		},
		{
			name:        "Nil Return Non-Existing Name",
			auxiliaries: auxiliaries{auxiliaryList: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, mockAuxiliary{randomUniqueIdentifiers[1]}}},
			arg:         "nonexistent",
			want:        nil,
		},
		{
			name:        "Flyweight1",
			auxiliaries: auxiliaries{auxiliaryList: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, mockAuxiliary{randomUniqueIdentifiers[1]}}},
			arg:         randomUniqueIdentifiers[0],
			want:        mockAuxiliary{randomUniqueIdentifiers[0]},
		},
		{
			name:        "Flyweight2",
			auxiliaries: auxiliaries{auxiliaryList: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, mockAuxiliary{randomUniqueIdentifiers[1]}}},
			arg:         randomUniqueIdentifiers[1],
			want:        mockAuxiliary{randomUniqueIdentifiers[1]},
		},
		{
			name:        "Unordered Flyweight1",
			auxiliaries: auxiliaries{auxiliaryList: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[1]}, mockAuxiliary{randomUniqueIdentifiers[0]}}},
			arg:         randomUniqueIdentifiers[0],
			want:        mockAuxiliary{randomUniqueIdentifiers[0]},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.auxiliaries.GetAuxiliary(tt.arg); got != tt.want {
				t.Errorf("GetAuxiliary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaries_Get(t *testing.T) {
	type fields struct {
		auxiliaryList []helpers.Auxiliary
	}
	tests := []struct {
		name   string
		fields fields
		want   []helpers.Auxiliary
	}{
		{
			name: "Empty AuxiliaryList",
			fields: fields{
				auxiliaryList: []helpers.Auxiliary{},
			},
			want: []helpers.Auxiliary{},
		},
		{
			name: "Non-empty AuxiliaryList",
			fields: fields{
				auxiliaryList: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, mockAuxiliary{randomUniqueIdentifiers[1]}},
			},
			want: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, mockAuxiliary{randomUniqueIdentifiers[1]}},
		},
		{
			name: "Nil Auxiliary",
			fields: fields{
				auxiliaryList: []helpers.Auxiliary{nil},
			},
			want: []helpers.Auxiliary{nil},
		},
		{
			name: "Nil Auxiliary in list",
			fields: fields{
				auxiliaryList: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, nil, mockAuxiliary{randomUniqueIdentifiers[1]}},
			},
			want: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, nil, mockAuxiliary{randomUniqueIdentifiers[1]}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaries := auxiliaries{
				auxiliaryList: tt.fields.auxiliaryList,
			}
			assert.Equalf(t, tt.want, auxiliaries.Get(), "Get()")
		})
	}
}

func TestNewAuxiliaries(t *testing.T) {
	type args struct {
		auxiliaryList []helpers.Auxiliary
	}
	tests := []struct {
		name string
		args args
		want helpers.Auxiliaries
	}{
		{
			name: "Empty AuxiliaryList",
			args: args{
				auxiliaryList: []helpers.Auxiliary{},
			},
			want: auxiliaries{
				auxiliaryList: []helpers.Auxiliary{},
			},
		},
		{
			name: "Non-empty AuxiliaryList",
			args: args{
				auxiliaryList: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, mockAuxiliary{randomUniqueIdentifiers[1]}},
			},
			want: auxiliaries{
				auxiliaryList: []helpers.Auxiliary{mockAuxiliary{randomUniqueIdentifiers[0]}, mockAuxiliary{randomUniqueIdentifiers[1]}},
			},
		},
		{
			name: "Nil Auxiliary",
			args: args{
				auxiliaryList: []helpers.Auxiliary{nil},
			},
			want: auxiliaries{
				auxiliaryList: []helpers.Auxiliary{nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxListInterface := make([]interface{}, len(tt.args.auxiliaryList))
			for i, v := range tt.args.auxiliaryList {
				auxListInterface[i] = v
			}
			assert.Equalf(t, tt.want, NewAuxiliaries(tt.args.auxiliaryList...), "NewAuxiliaries(%v)", auxListInterface...)
		})
	}
}
