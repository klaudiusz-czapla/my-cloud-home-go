package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitJwtCommand(v *viper.Viper) *cobra.Command {
	var jwtCmd = &cobra.Command{
		Use:   "jwt",
		Short: "Aggregates operations which can be done on JWT token",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Print("executing 'jwt' command..")
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	jwtCmd.Flags().Bool("decode-id-token", false, "Decode id token.")
	jwtCmd.Flags().Bool("decode-access-token", false, "Decode access token.")

	return jwtCmd
}
