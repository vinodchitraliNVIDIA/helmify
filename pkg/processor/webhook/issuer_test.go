package webhook

import (
	"testing"

	"github.com/vinodchitraliNVIDIA/helmify/pkg/metadata"

	"github.com/vinodchitraliNVIDIA/helmify/internal"
	"github.com/stretchr/testify/assert"
)

const issuerYaml = `apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: my-operator-selfsigned-issuer
  namespace: my-operator-system
spec:
  selfSigned: {}`

func Test_issuer_Process(t *testing.T) {
	var testInstance issuer

	t.Run("processed", func(t *testing.T) {
		obj := internal.GenerateObj(issuerYaml)
		processed, _, err := testInstance.Process(&metadata.Service{}, obj)
		assert.NoError(t, err)
		assert.Equal(t, true, processed)
	})
	t.Run("skipped", func(t *testing.T) {
		obj := internal.TestNs
		processed, _, err := testInstance.Process(&metadata.Service{}, obj)
		assert.NoError(t, err)
		assert.Equal(t, false, processed)
	})
}
