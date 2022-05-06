package query

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/marvin-hansen/goC8/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose            = true
	silent             = true
	fabric             = "SouthEastAsia"
	collectionTeachers = "teachers"
)

func TestSetup(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	goC8.CreateCollection(c, fabric, collectionTeachers, types.DocumentCollectionType, false)
	goC8.ImportCollectionData(c, fabric, collectionTeachers, sample_data.GetTeacherData(), silent)
}

func TestValidateQuery(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	query := "for t in teachers return t"

	res, err := c.Query.ValidateQuery(fabric, query)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res.Parsed, "Should be true / parsed")
	assert.Equal(t, false, res.Error, "Should be false / no error")
	goC8.PrintRes(res, verbose)
}

func TestQuery(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	query := "for t in teachers return t"
	res, err := c.Query.Query(fabric, query, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintQuery(res, verbose)
}

func TestTeardown(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	goC8.TeardownCollection(c, fabric, collectionTeachers)
}
