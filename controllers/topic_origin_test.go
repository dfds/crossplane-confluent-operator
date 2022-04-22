package controllers

import (
	"testing"

	topicv1alpha1 "github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/stretchr/testify/assert"
)

var (
	namespace = "default"
	label     = map[string]string{crossplanClaimNamespace: namespace}
	topic     = topicv1alpha1.Topic{}
)

func TestGetTopicOriginNamespace(t *testing.T) {
	assert := assert.New(t)

	// No label
	ns, err := getTopicOriginNamespace(topic)
	assert.EqualError(err, originNamespaceNotFound, "expected origin namespace not found")
	assert.Equal("", ns, "when namespace is not found it should be empty string")

	// With label
	topic.Labels = label
	ns, err = getTopicOriginNamespace(topic)
	assert.NoError(err)
	assert.Equal(namespace, ns, "expected to find an origin namespace")
}
