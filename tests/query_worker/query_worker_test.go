package query_worker

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	conf "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/marvin-hansen/goC8/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	verbose            = true
	silent             = true
	fabric             = "SouthEastAsia"
	collectionTeachers = "teachers"
	workerName         = "getAllTeachers"
)

func TestSetup(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	goC8.CreateCollection(c, fabric, collectionTeachers, types.DocumentCollectionType, false)
	goC8.ImportCollectionData(c, fabric, collectionTeachers, sample_data.GetTeacherData(), silent)
}

func TestCreateQueryWorker(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	query := "for t in teachers return t"
	bindVars := ""

	res, err := c.QueryWorker.CreateQueryWorker(fabric, workerName, query, bindVars)
	wait()

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestRunQueryWorker(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	bindVars := ""

	res, err := c.QueryWorker.RunQueryWorker(fabric, workerName, bindVars)
	wait()

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintJsonRes(res, verbose)
}

func TestReadAllQueryWorkers(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	res, err := c.QueryWorker.ReadAllQueryWorkers(fabric)
	wait()

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestUpdateQueryWorker(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	newQuery := "for t in testCollection return t"
	newBindVars := ""

	res, err := c.QueryWorker.UpdateQueryWorker(fabric, workerName, newQuery, newBindVars)
	wait()

	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestDeleteQueryWorker(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	err := c.QueryWorker.DeleteQueryWorker(fabric, workerName)
	wait()

	assert.NoError(t, err)
}

func TestTeardown(t *testing.T) {
	c := goC8.NewClient(conf.GetDefaultConfig())
	goC8.TeardownCollection(c, fabric, collectionTeachers)
}

func wait() {
	// tests randomly fail without the wait time.
	time.Sleep(time.Millisecond * 250)
}
