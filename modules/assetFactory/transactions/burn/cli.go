package burn

import (
	"bufio"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TransactionCommand(codec *codec.Codec) *cobra.Command {

	command := &cobra.Command{
		Use:   constants.BurnTransaction,
		Short: "Create and sign a transaction to burn an asset",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(authClient.GetTxEncoder(codec))
			cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

			message := Message{
				from:    cliContext.GetFromAddress(),
				assetID: types.BaseID{IDString: viper.GetString(constants.AssetID)},
			}

			if Error := message.ValidateBasic(); Error != nil {
				return Error
			}

			return authClient.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{message})
		},
	}

	command.Flags().String(constants.AssetID, "", "assetID")
	return flags.PostCommands(command)[0]
}

func NewTransactionCommand(codecMarshaler codec.Marshaler, txGenerator tx.Generator, accountRetriever tx.AccountRetriever) *cobra.Command {

	command := &cobra.Command{
		Use:   constants.BurnTransaction,
		Short: "Create and sign a transaction to burn an asset",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			cliContext := context.NewCLIContextWithInputAndFrom(bufioReader, args[0]).WithMarshaler(codecMarshaler)
			txFactory := tx.NewFactoryFromCLI(bufioReader).WithTxGenerator(txGenerator).WithAccountRetriever(accountRetriever)

			message := Message{
				from:    cliContext.GetFromAddress(),
				assetID: types.BaseID{IDString: viper.GetString(constants.AssetID)},
			}

			if Error := message.ValidateBasic(); Error != nil {
				return Error
			}

			return tx.GenerateOrBroadcastTx(cliContext, txFactory, message)
		},
	}

	command.Flags().String(constants.AssetID, "", "assetID")
	return flags.PostCommands(command)[0]
}
