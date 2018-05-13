package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInspectNetwork(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(InspectNetwork("testing", "--subnet=172.30.0.0/16"))
	assert.Nil(RemoveNetwork("testing"))
}

func TestStartAndRemoveNode(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(RunNode("halo", "halo", "halo", "halo", "halo", []string{"10.14.12.22"}))
	assert.NotNil(RemoveNode([]string{"halo"}))
}

func TestShowNode(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_cluster_size';"))
	assert.Nil(MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_incoming_addresses';"))
}
