// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package coralogixexporter

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/service/servicetest"
)

func TestLoadConfig(t *testing.T) {
	factories, _ := componenttest.NopFactories()
	factory := NewFactory()
	factories.Exporters[typestr] = factory
	// t.Log("new exporter " + typestr)
	cfg, err := servicetest.LoadConfigAndValidate(path.Join(".", "example", "config.yaml"), factories)
	require.NoError(t, err)
	apiConfig := cfg.Exporters[config.NewComponentID(typestr)].(*Config)
	err = apiConfig.Validate()
	require.NoError(t, err)
	assert.Equal(t, apiConfig, &Config{
		ExporterSettings: config.NewExporterSettings(config.NewComponentID("coralogix")),
		QueueSettings:    exporterhelper.DefaultQueueSettings(),
		RetrySettings:    exporterhelper.DefaultRetrySettings(),
		Endpoint:         "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		PrivateKey:       "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		AppName:          "APP_NAME",
		SubSystem:        "SUBSYSTEM_NAME",
	})
}
