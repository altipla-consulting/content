package content

import (
	// "bytes"
	// "encoding/json"
	// "strings"
	"testing"

	"github.com/stretchr/testify/require"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var (
	sess          sqlbuilder.Database
	testingModels db.Collection
)

type testTranslatedModel struct {
	ID          int64      `db:"id,omitempty"`
	Name        Translated `db:"name"`
	Description Translated `db:"description"`
}

func initTranslatedDB(t *testing.T) {
	cnf := &mysql.ConnectionURL{
		User:     "dev-user",
		Password: "dev-password",
		Host:     "localhost",
		Database: "test",
		Options: map[string]string{
			"charset":   "utf8mb4",
			"collation": "utf8mb4_bin",
			"parseTime": "true",
		},
	}
	var err error
	sess, err = mysql.Open(cnf)
	require.Nil(t, err)

	_, err = sess.Exec(`DROP TABLE IF EXISTS translated_test`)
	require.Nil(t, err)

	_, err = sess.Exec(`
		CREATE TABLE translated_test (
	  	id INT(11) NOT NULL AUTO_INCREMENT,
			name JSON,
			description JSON,

			PRIMARY KEY(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
	`)
	require.Nil(t, err)

	testingModels = sess.Collection("translated_test")

	require.Nil(t, testingModels.Truncate())
}

func finishTranslatedDB() {
	sess.Close()
}

func TestLoadSave(t *testing.T) {
	initTranslatedDB(t)
	defer finishTranslatedDB()

	model := new(testTranslatedModel)
	require.Nil(t, testingModels.InsertReturning(model))

	require.EqualValues(t, model.ID, 1)

	other := new(testTranslatedModel)
	require.Nil(t, testingModels.Find(1).One(other))
}

func TestLoadSaveContent(t *testing.T) {
	initTranslatedDB(t)
	defer finishTranslatedDB()

	model := &testTranslatedModel{
		Name: map[string]string{"es": "foo", "en": "bar"},
	}
	require.Nil(t, testingModels.InsertReturning(model))

	other := new(testTranslatedModel)
	require.Nil(t, testingModels.Find(1).One(other))

	require.Equal(t, other.Name["es"], "foo")
	require.Equal(t, other.Name["en"], "bar")
}

func TestLangChain(t *testing.T) {
	tests := []struct {
		content Translated
		chain   string
	}{
		{
			map[string]string{"es": "foo", "en": "bar", "fr": "baz"},
			"baz",
		},
		{
			map[string]string{"es": "foo", "en": "bar", "de": "baz"},
			"bar",
		},
		{
			map[string]string{"es": "foo", "it": "bar", "de": "baz"},
			"foo",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.content.LangChain("fr"), test.chain)
	}
}
