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

		properties := strings.Split(cliCommand.GetString(constants.Properties), ",")
		basePropertyList := make([]types.BaseProperty, 0)
		for _, property := range properties {
			traitIDProperty := strings.Split(property, ":")
			if len(traitIDProperty) == 2 && traitIDProperty[0] != "" {
				basePropertyList = append(basePropertyList,
					types.BaseProperty{
						BaseID:   types.BaseID{IDString: traitIDProperty[0]},
						BaseFact: types.BaseFact{BaseString: traitIDProperty[1]},
					})
			}
		}

		message := Message{
			From:             cliCommand.FromAddress(),
			ChainID:          types.BaseID{IDString: cliCommand.GetString(constants.ChainID)},
			MaintainersID:    types.BaseID{IDString: cliCommand.GetString(constants.MaintainersID)},
			ClassificationID: types.BaseID{IDString: cliCommand.GetString(constants.ClassificationID)},
			Properties:       &types.BaseProperties{BasePropertyList: basePropertyList},
			Lock:             types.BaseHeight{Height: cliCommand.GetInt(constants.Lock)},
			Burn:             types.BaseHeight{Height: cliCommand.GetInt(constants.Burn)},
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
