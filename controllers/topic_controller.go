package controllers

import (
	"context"
	"fmt"

	topicv1alpha1 "github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TopicPrereqsReconciler reconciles prerequisites for accessing Confluent Kafka topics
type TopicReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kafka.confluent.crossplane.io,resources=topic,verbs=get;list;watch
// +kubebuilder:rbac:groups=v1,resources=secret,verbs=get;list;watch

func (t *TopicReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// Get topics
	var topic topicv1alpha1.Topic
	if err := t.Get(ctx, req.NamespacedName, &topic); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Fetch origin namespace from label
	ns, err := getTopicOriginNamespace(topic)
	if err != nil {
		return ctrl.Result{}, err
	}
	fmt.Println("Topic name:", topic.ObjectMeta.Name)
	fmt.Println("Topic origin namespace:", ns)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (t *TopicReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&topicv1alpha1.Topic{}).
		Complete(t)
}
