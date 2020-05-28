package mint

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
		Use:   constants.MintTransaction,
		Short: "Create and sign transaction to mint an assetFactory",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(auth.DefaultTxEncoder(codec))
			cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

			var properties [2][]string
			for i := 0; i <= constants.MaxTraitCount; i++ {
				if viper.GetString(viper.GetString(constants.TraitID+strconv.Itoa(i))) != "" {
					traitID := viper.GetString(constants.TraitID + strconv.Itoa(i))
					property := viper.GetString(constants.Property + strconv.Itoa(i))
					properties[0] = append(properties[0], traitID)
					properties[1] = append(properties[1], property)
				}
			}
			message := message{
				from:             cliContext.GetFromAddress(),
				chainID:          types.BaseID{Binary: []byte(viper.GetString(constants.ChainID))},
				maintainersID:    types.BaseID{Binary: []byte(viper.GetString(constants.MaintainersID))},
				classificationID: types.BaseID{Binary: []byte(viper.GetString(constants.ClassificationID))},
				properties:       properties,
				lock:             viper.GetInt(constants.Lock),
				burn:             viper.GetInt(constants.Burn),
			}

			if err := message.ValidateBasic(); err != nil {
				return err
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
