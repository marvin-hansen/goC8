package query_worker

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	conf "github.com/marvin-hansen/goC8/tests/conf"
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
	c := goC8.NewClient(conf.GetDefaultConfig())
	goC8.CreateCollection(c, fabric, collectionTeachers, types.DocumentCollectionType, false)
	goC8.ImportCollectionData(c, fabric, collectionTeachers, sample_data.GetTeacherData(), silent)
}

func TestCreateQueryWorker(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	//query := "for t in teachers return t"
	res, err := c.QueryWorker.CreateQueryWorker(fabric)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestRunQueryWorker(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	//query := "for t in teachers return t"
	workerName := "name"
	res, err := c.QueryWorker.RunQueryWorker(fabric, workerName)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestReadAllQueryWorkers(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	//query := "for t in teachers return t"
	res, err := c.QueryWorker.ReadAllQueryWorkers(fabric)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestUpdateQueryWorker(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	//query := "for t in teachers return t"
	workerName := "name"
	res, err := c.QueryWorker.UpdateQueryWorker(fabric, workerName)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestDeleteQueryWorker(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	//query := "for t in teachers return t"
	workerName := "name"
	res, err := c.QueryWorker.DeleteQueryWorker(fabric, workerName)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestTeardown(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	goC8.TeardownCollection(c, fabric, collectionTeachers)
}