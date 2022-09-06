package deputize

import (
	"encoding/json"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_newTransactionRequest(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID, constants.MaintainedProperties, constants.AddMaintainer, constants.RemoveMaintainer, constants.MutateMaintainer})
	//cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	//fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	//require.Nil(t, err)

	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	//maintainedProperties, err := utilities.ReadProperties(maintainedProperty)
	//require.Equal(t, nil, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type args struct {
		baseReq              rest.BaseReq
		fromID               string
		toID                 string
		classificationID     string
		maintainedProperties string
		addMaintainer        bool
		removeMaintainer     bool
		mutateMaintainer     bool
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		// TODO: Add test cases.
		{"+ve", args{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false}, transactionRequest{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.maintainedProperties, tt.args.addMaintainer, tt.args.removeMaintainer, tt.args.mutateMaintainer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionRequest
	}{
		// TODO: Add test cases.
		{"+ve", transactionRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromCLI(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID, constants.MaintainedProperties, constants.AddMaintainer, constants.RemoveMaintainer, constants.MutateMaintainer})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	type args struct {
		cliCommand helpers.CLICommand
		cliContext context.CLIContext
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{BaseReq: rest.BaseReq{From: cliContext.GetFromAddress().String(), ChainID: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", ToID: "", ClassificationID: "", MaintainedProperties: "", AddMaintainer: false, RemoveMaintainer: false, MutateMaintainer: false}, args{cliCommand, cliContext}, transactionRequest{BaseReq: rest.BaseReq{From: cliContext.GetFromAddress().String(), ChainID: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", ToID: "", ClassificationID: "", MaintainedProperties: "", AddMaintainer: false, RemoveMaintainer: false, MutateMaintainer: false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.cliContext)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromJSON(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID, constants.MaintainedProperties, constants.AddMaintainer, constants.RemoveMaintainer, constants.MutateMaintainer})
	//cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	//fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	//require.Nil(t, err)

	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	//maintainedProperties, err := utilities.ReadProperties(maintainedProperty)
	//require.Equal(t, nil, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	jsonMessage, _ := json.Marshal(newTransactionRequest(testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false))
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	type args struct {
		rawMessage json.RawMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false}, args{jsonMessage}, transactionRequest{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			got, err := transactionRequest.FromJSON(tt.args.rawMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_GetBaseReq(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID, constants.MaintainedProperties, constants.AddMaintainer, constants.RemoveMaintainer, constants.MutateMaintainer})
	//cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	//fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	//require.Nil(t, err)

	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	//maintainedProperties, err := utilities.ReadProperties(maintainedProperty)
	//require.Equal(t, nil, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID, constants.MaintainedProperties, constants.AddMaintainer, constants.RemoveMaintainer, constants.MutateMaintainer})
	//cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	maintainedProperties, err := utilities.ReadProperties(maintainedProperty)
	require.Equal(t, nil, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false}, newMessage(fromAccAddress, baseIDs.NewID("fromID"), baseIDs.NewID("toID"), baseIDs.NewID("classificationID"), maintainedProperties, false, false, false), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			got, err := transactionRequest.MakeMsg()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_RegisterCodec(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID, constants.MaintainedProperties, constants.AddMaintainer, constants.RemoveMaintainer, constants.MutateMaintainer})
	//cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	//fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	//require.Nil(t, err)

	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	//maintainedProperties, err := utilities.ReadProperties(maintainedProperty)
	//require.Equal(t, nil, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			tr.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID, constants.MaintainedProperties, constants.AddMaintainer, constants.RemoveMaintainer, constants.MutateMaintainer})
	//cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	//fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	//require.Nil(t, err)

	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	//maintainedProperties, err := utilities.ReadProperties(maintainedProperty)
	//require.Equal(t, nil, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
