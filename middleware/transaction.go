package middleware

import (
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// TransactionHandler トランザクションハンドラ
func TransactionHandler(dbMap *gorp.DbMap) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			tx, _ := dbMap.Begin()

			c.Set("Tx", tx)

			err := next(c)
			if err != nil {
				tx.Rollback()
				return err
			}
			tx.Commit()

			return nil
		})
	}
}
