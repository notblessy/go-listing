package console

import (
	"github.com/notblessy/go-listing/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runHTTPServer = &cobra.Command{
	Use:   "httpsrv",
	Short: "run http server",
	Long:  `This subcommand is for starting the http server`,
	Run:   runHTTP,
}

func init() {
	rootCmd.AddCommand(runHTTPServer)
}

func runHTTP(cmd *cobra.Command, args []string) {
	psql := db.InitDB()
	defer db.CloseDB(psql)

	logrus.Info("DB Connected on HTTP")
}
