package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/QOM-One/QomApp/app"
	qomd "github.com/QOM-One/QomApp/cmd/qomd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := qomd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",     // Test the init cmd
		"qom-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "qom_7668378-1"),
	})

	err := svrcmd.Execute(rootCmd, "qomd", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := qomd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"dev0",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, "QOMD", app.DefaultNodeHome)
	require.Error(t, err)
}
