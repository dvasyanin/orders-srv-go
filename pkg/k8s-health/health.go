package k8s_health

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/heptiolabs/healthcheck"
	"time"
)

func PostgresPingCheck(db *pg.DB, timeout time.Duration) healthcheck.Check {
	return func() error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if db == nil {
			return fmt.Errorf("pg is nil")
		}

		db = db.WithContext(ctx)
		_, err := db.Exec("SELECT 1")
		if err != nil {
			return err
		}
		return nil
	}
}
