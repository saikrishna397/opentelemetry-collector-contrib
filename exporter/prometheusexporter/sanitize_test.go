// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheusexporter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSanitize(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	require.Equal(t, "", sanitize("", cfg.sanitizeLabel), "")
	require.Equal(t, "key_test", sanitize("_test", cfg.sanitizeLabel))
	require.Equal(t, "key_0test", sanitize("0test", cfg.sanitizeLabel))
	require.Equal(t, "test", sanitize("test", cfg.sanitizeLabel))
	require.Equal(t, "test__", sanitize("test_/", cfg.sanitizeLabel))
	require.Equal(t, "key__test", sanitize("__test", cfg.sanitizeLabel))
	//enable sanitizeLabel
	cfg.sanitizeLabel = true
	require.Equal(t, "_test", sanitize("_test", cfg.sanitizeLabel))
	require.Equal(t, "", sanitize("", cfg.sanitizeLabel), "")
	require.Equal(t, "key_0test", sanitize("0test", cfg.sanitizeLabel))
	require.Equal(t, "test", sanitize("test", cfg.sanitizeLabel))
	require.Equal(t, "key__test", sanitize("__test", cfg.sanitizeLabel))

}
