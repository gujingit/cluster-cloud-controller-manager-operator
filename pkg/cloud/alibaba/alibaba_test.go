package alibaba

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetResources(t *testing.T) {
	resources := GetResources()
	assert.Len(t, resources, 1)

	var names, kinds []string
	for _, r := range resources {
		names = append(names, r.GetName())
		kinds = append(kinds, r.GetObjectKind().GroupVersionKind().Kind)
	}

	assert.Contains(t, names, "alibaba-cloud-controller-manager")
	assert.Contains(t, kinds, "Deployment")

}
