package output

import (
	"testing"
	"time"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/stretchr/testify/assert"
)

func TestRawOutput_Format(t *testing.T) {
	t.Parallel()

	timestamp, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	someLabels := labels.New(labels.Label{Name: "type", Value: "test"})

	tests := map[string]struct {
		options      *LogOutputOptions
		timestamp    time.Time
		lbls         *labels.Labels
		maxLabelsLen int
		line         string
		expected     string
	}{
		"empty line": {
			&LogOutputOptions{Timezone: time.UTC, NoLabels: false},
			timestamp,
			&someLabels,
			0,
			"",
			"",
		},
		"non empty line": {
			&LogOutputOptions{Timezone: time.UTC, NoLabels: false},
			timestamp,
			&someLabels,
			0,
			"Hello world",
			"Hello world",
		},
	}

	for testName, testData := range tests {
		testData := testData

		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out := &RawOutput{testData.options}
			actual := out.Format(testData.timestamp, testData.lbls, testData.maxLabelsLen, testData.line)

			assert.Equal(t, testData.expected, actual)
		})
	}
}
