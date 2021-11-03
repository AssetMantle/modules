package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	metaTypes "github.com/persistenceOne/persistenceSDK/modules/metas/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/spf13/cobra"
	"strconv"
)

const (
	FlagNodeID = "node-id"
	FlagIP     = "ip"
)

// NewTxCmd returns a root CLI command handler for all x/staking transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        metaTypes.ModuleName,
		Short:                      "Meta module txs",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewRevealCmd(),
	)

	return txCmd
}

func NewRevealCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reveal [dataType] [dataValue]",
		Short: "reveals data",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFrom(clientCtx.From)
			data, err := getData(args[0], args[1])
			msg := metaTypes.NewMsgReveal(clientCtx.GetFromAddress(), data)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagIP, "", fmt.Sprintf("The node's public IP. It takes effect only when used in combination with --%s", flags.FlagGenerateOnly))
	cmd.Flags().String(FlagNodeID, "", "The node's ID")
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func getData(dataType string, dataValue string) (types.Data, error) {
	switch dataType {
	case "S":
		return base.NewStringData(dataValue), nil
	case "H":
		h, err := strconv.ParseInt(dataValue, 10, 64)
		if err != nil {
			return nil, err
		}
		return base.NewHeightData(base.NewHeight(h)), nil
	case "I":
		id := base.NewID(dataValue)
		return base.NewIDData(&id), nil
	case "A":
		return base.ReadAccAddressData(dataValue)
	case "D":
		dec, err := sdkTypes.NewDecFromStr(dataValue)
		if err != nil {
			return nil, err
		}
		return base.NewDecData(dec), nil
	default:
		return nil, errors.IncorrectMessage

	}
}
