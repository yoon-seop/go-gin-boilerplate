package database

import (
	"fmt"
	"go-gin-boilerplate/config"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func WithTransaction(fn func(tx *gorm.DB) error, identifier string) (err error) {
	tx := config.DB.Begin()

	if tx.Error != nil {
		log.Error().Err(tx.Error).Str("identifier", identifier).Msg("Transaction begin failed")
		return fmt.Errorf("transaction begin failed: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			if rerr := tx.Rollback().Error; rerr != nil {
				log.Error().Interface("recovered_value", r).Err(rerr).Str("identifier", identifier).Msg("Panic occurred, and Rollback failed")
			} else {
				log.Error().Interface("recovered_value", r).Str("identifier", identifier).Msg("Panic occurred, successfully rollbacked")
			}
			err = fmt.Errorf("panic occurred. (identifier: %s, recovered: %v)", identifier, r)
		}
	}()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback().Error; rerr != nil {
			log.Error().Err(rerr).Str("identifier", identifier).Msg("Function failed, and Rollback failed")
			return fmt.Errorf("transaction failed (%w) and rollback failed (%w)", err, rerr)
		}
		log.Info().Err(err).Str("identifier", identifier).Msg("Transaction function failed, successfully rollbacked")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Error().Err(err).Str("identifier", identifier).Msg("Failed to commit transaction")
		return fmt.Errorf("failed to commit transaction. (identifier: %s, %w)", identifier, err)
	}

	log.Info().Str("identifier", identifier).Msg("Successfully committed transaction")

	return nil
}
