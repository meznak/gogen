package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("GOGEN_HOME", "..")
	// Test Singleton
	c := NewConfig()
	c2 := NewConfig()
	assert.Equal(t, c, c2)

	global := Global{Debug: false, Verbose: false, UseOutputQueue: true, GeneratorWorkers: 1, OutputWorkers: 1}
	assert.Equal(t, c.Global, global)
	defaultSample := Sample{Disabled: false, Generator: "sample", Outputter: "stdout", Rater: "config", Interval: 60, Delay: 0, Count: 0, Earliest: "now", Latest: "now", RandomizeCount: 0.20000000298023224, RandomizeEvents: true}
	assert.Equal(t, c.defaultSample, defaultSample)
}
