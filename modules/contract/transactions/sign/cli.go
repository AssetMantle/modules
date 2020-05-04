package sign

import (
	"bufio"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/persistenceOne/persistenceSDK/modules/contract/constants"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func TransactionCommand(codec *codec.Codec) *cobra.Command {

	command := &cobra.Command{
		Use:   "sign",
		Short: "Create and sign transaction to sign at contract",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(auth.DefaultTxEncoder(codec))
			cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

			message := Message{
				From: cliContext.GetFromAddress(),
			}

			if Error := message.ValidateBasic(); Error != nil {
				return Error
			}

			return client.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{message})
		},
	}

	command.Flags().String(constants.ContractFlag, "", "Contract")
	return command
}
