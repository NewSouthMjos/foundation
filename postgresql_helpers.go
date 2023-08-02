package foundation

import (
	"database/sql"
	"errors"

	"github.com/getsentry/sentry-go"
	fpg "github.com/ri-nat/foundation/postgresql"
)

func (app *Application) GetPostgreSQL() *sql.DB {
	component := app.GetComponent(fpg.ComponentName)
	if component == nil {
		err := errors.New("PostgreSQL component is not registered")
		sentry.CaptureException(err)
		app.Logger.Fatal(err)
	}

	pg, ok := component.(*fpg.PostgreSQLComponent)
	if !ok {
		err := errors.New("PostgreSQL component is not of type *fpg.PostgreSQLComponent")
		sentry.CaptureException(err)
		app.Logger.Fatal(err)
	}

	return pg.Connection
}
