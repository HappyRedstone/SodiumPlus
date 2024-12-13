package server

import (
	internalServer "github.com/HappyRedstone/SodiumPlus/builder/internal/server"
	"github.com/packwiz/packwiz/core"
	"github.com/spf13/cobra"
)

var CleanCommand = &cobra.Command{
	Use:     "clean",
	Short:   "Clean bundled server files.",
	Long:    `Clean bundled server files.`,
	Aliases: []string{"c"},

	RunE: func(cmd *cobra.Command, args []string) error {
		pack, err := core.LoadPack()

		if err != nil {
			return err
		}

		return internalServer.Clean("output", &pack)
	},
}
