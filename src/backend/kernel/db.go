package kernel

import (
    "context"
    "database/sql"
    "time"

    "github.com/gadelkareem/go-helpers"

    "github.com/astaxie/beego/logs"
    "github.com/gadelkareem/que"
    "github.com/go-pg/pg/v9"
    "github.com/jackc/pgx/v4/pgxpool"
    _ "github.com/lib/pq"
)

type PgDb struct {
    *pg.DB
}

var (
    dbURL string
)

func initDBConfig() {
    dbURL = App.ConfigOrEnvVar("dbAddress", "DATABASE_URL")
}

func NewDB() *PgDb {
    db, err := newDb(dbURL)
    if err != nil {
        logs.Error("Database connection error: %s", err)
    }

    if IsDev() {
        db.EnableLogging()
    }

    return db
}

func (db *PgDb) EnableLogging() {
    db.AddQueryHook(dbLogger{})
}

func newDb(u string) (db *PgDb, err error) {
    err = h.Retry(func() (e error) {

        db, e = NewDbWithOptions(u,
            15*time.Second,
            false,
            App.Config.DefaultInt("dbPoolSize", 5),
            0)
        if e != nil {
            time.Sleep(5 * time.Second)
        }
        return e
    }, MaxInt)
    return
}

func NewDbWithOptions(u string, timeout time.Duration, retryStatementTimeout bool, poolSize, maxRetries int) (*PgDb, error) {
    o, err := pg.ParseURL(u)
    h.PanicOnError(err)
    o.MaxRetries = maxRetries
    o.RetryStatementTimeout = retryStatementTimeout
    o.PoolSize = poolSize
    o.DialTimeout = timeout
    o.ReadTimeout = timeout
    o.WriteTimeout = timeout
    db := pg.Connect(o)

    _, err = db.Exec("SELECT 1")
    if err != nil {
        logs.Error("Error connecting to Postgres %v", err)
        return nil, err
    }
    return &PgDb{db}, err
}

func DB() (db *sql.DB) {
    err := h.Retry(func() (e error) {
        db, e = sql.Open(
            "postgres",
            dbURL,
        )
        if e != nil {
            time.Sleep(5 * time.Second)
        }
        return e
    }, MaxInt)
    if err != nil {
        logs.Error("Database connection error: %s", err)
    }
    return
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
    return c, nil
}
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
    query, err := q.FormattedQuery()
    h.PanicOnError(err)
    logs.Debug("%s", query)
    return err
}

func Que(maxConnections int32) (*que.Client, *pgxpool.Pool) {
    connPoolConfig, err := pgxpool.ParseConfig(dbURL)
    h.PanicOnError(err)
    connPoolConfig.MaxConns = maxConnections
    connPoolConfig.AfterConnect = que.PrepareStatements

    p, err := pgxpool.ConnectConfig(context.Background(), connPoolConfig)
    h.PanicOnError(err)
    return que.NewClient(p), p
}
