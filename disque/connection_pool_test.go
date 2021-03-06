package disque

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
)

type DisquePoolSuite struct {
	suite.Suite
}

func TestDisquePoolSuite(t *testing.T) {
	suite.Run(t, new(DisquePoolSuite))
}

func (s *DisquePoolSuite) SetupTest() {
}

func (s *DisquePoolSuite) SetupSuite() {
}

func (s *DisquePoolSuite) TestWithPoolOfOne() {
	hosts := []string{"127.0.0.1:7711"}
	d := NewDisquePool(hosts, 1000, 1, 1, time.Hour)

	c, err := d.Get(context.Background())
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), c)

	// return the connection and close the pool
	d.Put(c)
	d.Close()

	// try getting a connection from the closed pool
	c, err = d.Get(context.Background())
	assert.NotNil(s.T(), err)
}
