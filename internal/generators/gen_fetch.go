package main

import (
	"os"
	"text/template"

	idb "agola.io/agola/internal/db"
	csobjects "agola.io/agola/internal/services/configstore/db/objects"
	rsobjects "agola.io/agola/internal/services/runservice/db/objects"
)

func genFetch() {
	f, err := os.Create("fetch.go")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	switch componentName {
	case "runservice":
		err = fetchTemplate.Execute(f, struct {
			ObjectInfos []idb.ObjectInfo
			TypesPath   string
		}{
			ObjectInfos: rsobjects.ObjectsInfo,
			TypesPath:   "agola.io/agola/services/runservice/types",
		})
	case "configstore":
		err = fetchTemplate.Execute(f, struct {
			ObjectInfos []idb.ObjectInfo
			TypesPath   string
		}{
			ObjectInfos: csobjects.ObjectsInfo,
			TypesPath:   "agola.io/agola/services/configstore/types",
		})
	}
	if err != nil {
		panic(err)
	}
}

var fetchTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package db

import (
	stdsql "database/sql"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"

	"agola.io/agola/internal/errors"
	"agola.io/agola/internal/sql"
	"{{ .TypesPath }}"
)
{{- range $index, $oi := .ObjectInfos }}

func (d *DB) fetch{{ $oi.Name }}s(tx *sql.Tx, q sq.Sqlizer) ([]*types.{{ $oi.Name }}, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scan{{ $oi.Name }}s(rows, tx.ID())
}

func (d *DB) scan{{ $oi.Name }}(rows *stdsql.Rows, additionalFields []interface{}) (*types.{{ $oi.Name }}, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.{{ $oi.Name }}{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal {{ $oi.Name }}")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scan{{ $oi.Name }}s(rows *stdsql.Rows, txID string) ([]*types.{{ $oi.Name }}, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.{{ $oi.Name }}{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scan{{ $oi.Name }}(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		v.TxID = txID
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}
{{- end }}
`))
