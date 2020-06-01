package interNFT

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func QueryCommand(codec *codec.Codec) *cobra.Command {
	command := &cobra.Command{
		Use:   constants.InterNFTQuery,
		Short: "Query an assetFactory.",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			cliContext := context.NewCLIContext().WithCodec(codec)

			bytes := packageCodec.MustMarshalJSON(query{
				Address: viper.GetString(constants.InterNFTID),
			})

			response, _, queryWithDataError := cliContext.QueryWithData(strings.Join([]string{"", "custom", constants.QuerierRoute, constants.InterNFTQuery}, "/"), bytes)
			if queryWithDataError != nil {
				return queryWithDataError
			}

			var interNFT types.InterNFT
			unmarshalJSONError := codec.UnmarshalJSON(response, &interNFT)
			if unmarshalJSONError != nil {
				return unmarshalJSONError
			}
			return cliContext.PrintOutput(interNFT)
		},
	}

	command.Flags().String(constants.InterNFTID, "", "interNFTID")
	return command
}
