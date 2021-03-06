package dinosql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const inputGoose = `
-- +goose Up
ALTER TABLE archived_jobs ADD COLUMN expires_at TIMESTAMP WITH TIME ZONE;

-- +goose Down
ALTER TABLE archived_jobs DROP COLUMN expires_at;
`

const outputGoose = `
-- +goose Up
ALTER TABLE archived_jobs ADD COLUMN expires_at TIMESTAMP WITH TIME ZONE;
`

const inputMigrate = `
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE people (id int);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE people;
`

const outputMigrate = `
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE people (id int);
`

const inputTern = `
-- Write your migrate up statements here
ALTER TABLE todo RENAME COLUMN done TO is_done;
---- create above / drop below ----
ALTER TABLE todo RENAME COLUMN is_done TO done;
`

const outputTern = `
-- Write your migrate up statements here
ALTER TABLE todo RENAME COLUMN done TO is_done;
`

func TestRemoveRollback(t *testing.T) {
	if diff := cmp.Diff(outputGoose, RemoveRollbackStatements(inputGoose)); diff != "" {
		t.Errorf("goose migration mismatch:\n%s", diff)
	}
	if diff := cmp.Diff(outputMigrate, RemoveRollbackStatements(inputMigrate)); diff != "" {
		t.Errorf("sql-migrate migration mismatch:\n%s", diff)
	}
	if diff := cmp.Diff(outputTern, RemoveRollbackStatements(inputTern)); diff != "" {
		t.Errorf("tern migration mismatch:\n%s", diff)
	}
}
