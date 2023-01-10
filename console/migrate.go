package console

import (
	"github.com/notblessy/go-listing/db"
	"github.com/notblessy/go-listing/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Long:  `This subcommand used to migrate database`,
	Run:   migrate,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) {
	psql := db.InitDB()
	defer db.CloseDB(psql)

	if err := psql.AutoMigrate(&model.Product{}); err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("All migrations success!")
}
