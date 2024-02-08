package db

import (
	"encoding/json"
	"time"

	sq "github.com/huandu/go-sqlbuilder"
	"github.com/sorintlab/errors"

	"agola.io/agola/internal/services/notification/types"
	"agola.io/agola/internal/sqlg"
	"agola.io/agola/internal/sqlg/sql"
)

const runWebhookLimit = 25

var MigrationRunWebhookLimit = runWebhookLimit

func (d *DB) MigrateFuncs() map[uint]sqlg.MigrateFunc {
	return map[uint]sqlg.MigrateFunc{
		2: d.migrateV2,
		3: d.migrateV3,
		4: d.migrateV4,
	}
}

func (d *DB) migrateV2(tx *sql.Tx) error {
	var ddlPostgres = []string{
		"create table if not exists lastruneventsequence (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamptz NOT NULL, update_time timestamptz NOT NULL, value bigint NOT NULL, PRIMARY KEY (id))",
	}

	var ddlSqlite3 = []string{
		"create table if not exists lastruneventsequence (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, value bigint NOT NULL, PRIMARY KEY (id))",
	}

	var stmts []string
	switch d.sdb.Type() {
	case sql.Postgres:
		stmts = ddlPostgres
	case sql.Sqlite3:
		stmts = ddlSqlite3
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (d *DB) migrateV3(tx *sql.Tx) error {
	var ddlPostgres = []string{
		"create table if not exists commitstatus (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamptz NOT NULL, update_time timestamptz NOT NULL, project_id varchar NOT NULL, state varchar NOT NULL, commit_sha varchar NOT NULL, run_counter bigint NOT NULL, description varchar NOT NULL, context varchar NOT NULL, PRIMARY KEY (id))",
		"create table if not exists commitstatusdelivery (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamptz NOT NULL, update_time timestamptz NOT NULL, sequence bigint generated by default as identity NOT NULL UNIQUE, commit_status_id varchar NOT NULL, delivery_status varchar NOT NULL, delivered_at timestamptz, PRIMARY KEY (id), foreign key (commit_status_id) references commitstatus(id))",
		"create index if not exists commitstatusdelivery_sequence_idx on commitstatusdelivery(sequence)",
	}

	var ddlSqlite3 = []string{
		"create table if not exists commitstatus (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, project_id varchar NOT NULL, state varchar NOT NULL, commit_sha varchar NOT NULL, run_counter bigint NOT NULL, description varchar NOT NULL, context varchar NOT NULL, PRIMARY KEY (id))",
		"create table if not exists commitstatusdelivery (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, sequence integer NOT NULL UNIQUE, commit_status_id varchar NOT NULL, delivery_status varchar NOT NULL, delivered_at timestamp, PRIMARY KEY (id), foreign key (commit_status_id) references commitstatus(id))",
		"create index if not exists commitstatusdelivery_sequence_idx on commitstatusdelivery(sequence)",
	}

	var stmts []string
	switch d.sdb.Type() {
	case sql.Postgres:
		stmts = ddlPostgres
	case sql.Sqlite3:
		stmts = ddlSqlite3
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (d *DB) migrateV4(tx *sql.Tx) error {
	var ddlPostgres = []string{
		"alter table runwebhook add column project_id varchar",
	}

	var ddlSqlite3 = []string{
		"create table new_runwebhook (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, payload blob NOT NULL, project_id varchar NOT NULL, PRIMARY KEY (id))",
		"insert into new_runwebhook select *, '' from runwebhook",
		"DROP TABLE runwebhook",
		"ALTER TABLE new_runwebhook RENAME TO runwebhook",
	}

	var stmts []string
	switch d.sdb.Type() {
	case sql.Postgres:
		stmts = ddlPostgres
	case sql.Sqlite3:
		stmts = ddlSqlite3
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			return errors.WithStack(err)
		}
	}

	runwebhookSb := sq.NewSelectBuilder().Select("id", "revision", "creation_time", "update_time", "payload").From("runwebhook").OrderBy("id").Limit(MigrationRunWebhookLimit)

	type runWebhookType struct {
		ID           string
		CreationTime time.Time
		UpdateTime   time.Time
		Revision     uint64
		Payload      []byte
		ProjectID    string
	}

	runWebhooks := make([]runWebhookType, 0)

	var lastID string
	hasNext := true
	for hasNext {
		rows, err := d.query(tx, runwebhookSb)
		if err != nil {
			return errors.WithStack(err)
		}

		count := 0
		for rows.Next() {
			var r runWebhookType
			err = rows.Scan(&r.ID, &r.Revision, &r.CreationTime, &r.UpdateTime, &r.Payload)
			if err != nil {
				return errors.WithStack(err)
			}

			var runWebhook types.RunWebhook
			if err = json.Unmarshal(r.Payload, &runWebhook); err != nil {
				return errors.WithStack(err)
			}

			r.ProjectID = runWebhook.ProjectInfo.ProjectID

			runWebhooks = append(runWebhooks, r)
			lastID = r.ID
			count++
		}
		if err := rows.Err(); err != nil {
			return errors.WithStack(err)
		}

		for _, r := range runWebhooks {
			ub := sq.NewUpdateBuilder()
			ub.Update("runwebhook").Set(ub.Assign("project_id", r.ProjectID)).Where(ub.E("id", r.ID))
			if _, err := d.exec(tx, ub); err != nil {
				return errors.WithStack(err)
			}
		}

		hasNext = count == MigrationRunWebhookLimit
		runwebhookSb.Where(runwebhookSb.G("id", lastID))
	}

	ddlPostgres = []string{
		"alter table runwebhook alter column project_id set NOT NULL",
	}

	ddlSqlite3 = []string{}

	switch d.sdb.Type() {
	case sql.Postgres:
		stmts = ddlPostgres
	case sql.Sqlite3:
		stmts = ddlSqlite3
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
