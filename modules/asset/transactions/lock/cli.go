package lock

import (
	"bufio"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/persistenceOne/persistenceSDK/modules/asset/constants"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func TransactionCommand(codec *codec.Codec) *cobra.Command {

	command := &cobra.Command{
		Use:   constants.LockTransaction,
		Short: "Create and sign transaction to lock an asset",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(auth.DefaultTxEncoder(codec))
			cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

			message := Message{
				From:    cliContext.GetFromAddress(),
				Address: viper.GetString(constants.AddressFlag),
				Lock:    viper.GetBool(constants.LockFlag),
			}

			if err := message.ValidateBasic(); err != nil {
				return err
			}

			return client.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{message})
		},
	}

	command.Flags().String(constants.AddressFlag, "", "address")
	command.Flags().Bool(constants.LockFlag, false, "lock")
	return command
}
