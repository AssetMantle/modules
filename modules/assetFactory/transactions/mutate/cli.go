package mutate

import (
	"bufio"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

func TransactionCommand(codec *codec.Codec) *cobra.Command {

	command := &cobra.Command{
		Use:   constants.MutateTransaction,
		Short: "Create and sign a transaction to mutate an asset",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(auth.DefaultTxEncoder(codec))
			cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

			var propertyList []types.Property
			for i := 0; i <= constants.MaxTraitCount; i++ {
				if viper.GetString(viper.GetString(constants.TraitID+strconv.Itoa(i))) != "" {
					var basePropertyList []types.BaseProperty
					basePropertyList = append(basePropertyList,
						types.BaseProperty{
							BaseID:   types.BaseID{BaseString: viper.GetString(constants.TraitID + strconv.Itoa(i))},
							BaseFact: types.BaseFact{BaseBytes: []byte(viper.GetString(constants.Property + strconv.Itoa(i)))},
						})
				}
			}
			message := Message{
				from:             cliContext.GetFromAddress(),
				chainID:          types.BaseID{BaseString: viper.GetString(constants.ChainID)},
				maintainersID:    types.BaseID{BaseString: viper.GetString(constants.MaintainersID)},
				classificationID: types.BaseID{BaseString: viper.GetString(constants.ClassificationID)},
				propertyList:     propertyList,
				lock:             types.BaseHeight{Height: viper.GetInt(constants.Lock)},
				burn:             types.BaseHeight{Height: viper.GetInt(constants.Burn)},
			}

			if Error := message.ValidateBasic(); Error != nil {
				return Error
			}

			return client.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{message})
		},
	}
	command.Flags().String(constants.ChainID, "", "chainID")
	command.Flags().String(constants.MaintainersID, "", "maintainersID")
	command.Flags().String(constants.ClassificationID, "", "classificationID")
	for i := 0; i <= constants.MaxTraitCount; i++ {
		command.Flags().String(constants.TraitID+strconv.Itoa(i), "", "traitID")
		command.Flags().String(constants.Property+strconv.Itoa(i), "", "property")
	}
	command.Flags().Int(constants.Lock, -1, "lock")
	command.Flags().Int(constants.Burn, -1, "burn")
	return command
}
