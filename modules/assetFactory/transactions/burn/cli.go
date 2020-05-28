package burn

import (
	"bufio"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
)

func TransactionCommand(codec *codec.Codec) *cobra.Command {

	command := &cobra.Command{
		Use:   constants.BurnTransaction,
		Short: "Create and sign transaction to burn an assetFactory",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(auth.DefaultTxEncoder(codec))
			cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

			message := Message{
				from:    cliContext.GetFromAddress(),
				assetID: types.BaseID{BaseString: viper.GetString(constants.AssetID)},
			}

			if err := message.ValidateBasic(); err != nil {
				return err
			}

			return client.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{message})
		},
	}

	command.Flags().String(constants.AssetID, "", "assetID")
	return command
}
