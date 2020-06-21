package mint

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"github.com/spf13/cobra"
	"strings"
)

func TransactionCommand(codec *codec.Codec) *cobra.Command {
	makeMessage := func(cliCommand types.CLICommand) sdkTypes.Msg {

		properties := strings.Split(cliCommand.ReadString(constants.Properties), constants.PropertiesSeparator)
		basePropertyList := make([]types.Property, 0)
		for _, property := range properties {
			traitIDAndProperty := strings.Split(property, constants.TraitIDAndPropertySeparator)
			if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
				basePropertyList = append(basePropertyList, types.NewProperty(types.NewID(traitIDAndProperty[0]), types.NewFact(traitIDAndProperty[1], types.NewSignatures(nil))))
			}
		}

		message := Message{
			From:             cliCommand.GetFromAddress(),
			ChainID:          types.NewID(cliCommand.ReadString(constants.ChainID)),
			MaintainersID:    types.NewID(cliCommand.ReadString(constants.MaintainersID)),
			ClassificationID: types.NewID(cliCommand.ReadString(constants.ClassificationID)),
			Properties:       types.NewProperties(basePropertyList),
			Lock:             types.NewHeight(cliCommand.ReadInt(constants.Lock)),
			Burn:             types.NewHeight(cliCommand.ReadInt(constants.Burn)),
		}
		return message
	}
	return types.NewCLICommand(
		constants.MintTransaction,
		constants.MintTransactionShort,
		constants.MintTransactionLong,
		[]types.CLIFlag{
			constants.Properties,
			constants.ChainID,
			constants.MaintainersID,
			constants.ClassificationID,
			constants.Lock,
			constants.Burn},
	).CreateTransactionCommand(codec, makeMessage)
}
