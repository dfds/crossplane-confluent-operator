package controllers

import (
	"errors"

	topicv1alpha1 "github.com/dfds/provider-confluent/apis/topic/v1alpha1"
)

const (
	crossplanClaimNamespace = "crossplane.io/claim-namespace"
	originNamespaceNotFound = "origin namespace not found because no label set"
)

// GetTopicOrigin will look for labels tracing origin namespace of resource
func getTopicOriginNamespace(t topicv1alpha1.Topic) (string, error) {
	labels := t.GetLabels()
	if ns, exists := labels[crossplanClaimNamespace]; exists {
		return ns, nil
	}
	return "", errors.New(originNamespaceNotFound)
}
