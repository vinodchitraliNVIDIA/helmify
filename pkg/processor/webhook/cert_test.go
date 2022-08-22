package webhook

import (
	"testing"

	"github.com/vinodchitraliNVIDIA/helmify/pkg/metadata"

	"github.com/vinodchitraliNVIDIA/helmify/internal"
	"github.com/stretchr/testify/assert"
)

const certYaml = `apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: my-operator-serving-cert
  namespace: my-operator-system
spec:
  dnsNames:
  - my-operator-webhook-service.my-operator-system.svc
  - my-operator-webhook-service.my-operator-system.svc.cluster.local
  issuerRef:
    kind: Issuer
    name: my-operator-selfsigned-issuer
  secretName: webhook-server-cert`

func Test_cert_Process(t *testing.T) {
	var testInstance cert

	t.Run("processed", func(t *testing.T) {
		obj := internal.GenerateObj(certYaml)
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
