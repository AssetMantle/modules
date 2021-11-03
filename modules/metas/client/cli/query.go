package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/persistenceOne/persistenceSDK/modules/metas/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the staking module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetCmdQueryMeta(),
		GetCmdQueryParams(),
	)

	return queryCmd
}

// GetCmdQueryValidator implements the validator query command.
func GetCmdQueryMeta() *cobra.Command {
	bech32PrefixValAddr := sdk.GetConfig().GetBech32ValidatorAddrPrefix()

	cmd := &cobra.Command{
		Use:   "query [metaID]",
		Short: "Query meta",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about an individual validator.

Example:
$ %s query staking validator %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj
`,
				version.AppName, bech32PrefixValAddr,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryMetaRequest{MetaID: args[0]}
			res, err := queryClient.GetMeta(cmd.Context(), params)
			if err != nil {
				return err
			}

			if res.GetSuccess() {
				return clientCtx.PrintProto(&res.Value)
			} else {
				return fmt.Errorf(res.Error)
			}

		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parameters",
		Args:  cobra.NoArgs,
		Short: "Query the current meta parameters information",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.GetParameters(context.Background(), &types.QueryParametersRequest{})
			if err != nil {
				return err
			}

			if res.GetSuccess() {
				return clientCtx.PrintProto(&res.Parameters)
			} else {
				return fmt.Errorf(res.Error)
			}

		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
