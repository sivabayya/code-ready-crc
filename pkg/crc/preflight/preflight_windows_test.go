package preflight

import (
	"testing"

	"github.com/code-ready/crc/pkg/crc/config"
	"github.com/code-ready/crc/pkg/crc/network"
	"github.com/stretchr/testify/assert"
)

func TestCountConfigurationOptions(t *testing.T) {
	cfg := config.New(config.NewEmptyInMemoryStorage())
	RegisterSettings(cfg)
	assert.Len(t, cfg.AllConfigs(), 10)
}

func TestCountPreflights(t *testing.T) {
	assert.Len(t, getPreflightChecks(false, false, network.SystemNetworkingMode), 15)
	assert.Len(t, getPreflightChecks(true, true, network.SystemNetworkingMode), 15)

	assert.Len(t, getPreflightChecks(false, false, network.UserNetworkingMode), 16)
	assert.Len(t, getPreflightChecks(true, true, network.UserNetworkingMode), 16)
}
