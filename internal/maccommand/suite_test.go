package maccommand

import (
	"github.com/andrewflash/benam-lora-network-server/internal/storage"
	"github.com/andrewflash/benam-lora-network-server/internal/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestBase struct {
	suite.Suite
}

func (ts *TestBase) SetupSuite() {
	assert := require.New(ts.T())
	conf := test.GetConfig()
	assert.NoError(storage.Setup(conf))
}

func (ts *TestBase) SetupTest() {
	storage.RedisClient().FlushAll()
}
