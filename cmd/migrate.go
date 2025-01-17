package cmd

import (
	"github.com/ankorstore/yokai/fxcore"
	"github.com/ankorstore/yokai/fxorm"
	"github.com/ankorstore/yokai/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/mohammadbnei/vos-template/internal/model"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run application ORM migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// bootstrap, apply migrations then shutdown
		fxcore.NewBootstrapper().
			WithOptions(fxorm.FxOrmModule).
			WithContext(cmd.Context()).
			RunApp(
				fx.Invoke(func(logger *log.Logger, db *gorm.DB, sd fx.Shutdowner) error {
					logger.Info().Msg("starting ORM auto migration")

					err := db.AutoMigrate(model.GetModels()...)
					if err != nil {
						logger.Error().Err(err).Msg("error during ORM auto migration")
					} else {
						logger.Info().Msg("ORM auto migration success")
					}

					// shutdown
					return sd.Shutdown()
				}),
			)
	},
}
