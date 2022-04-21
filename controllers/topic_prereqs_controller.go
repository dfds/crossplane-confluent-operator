package controllers

import (
	"context"
	"fmt"

	topicv1alpha1 "github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TopicPrereqsReconciler reconciles prerequisites for accessing Confluent Kafka topics
type TopicPrereqsReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kafka.confluent.crossplane.io,resources=topic,verbs=get;list;watch

func (t *TopicPrereqsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// var ns corev1.Namespace
	var topic topicv1alpha1.Topic

	if err := t.Get(ctx, req.NamespacedName, &topic); err != nil {
		return ctrl.Result{}, err
	}

	fmt.Println("Topic name:", topic.ObjectMeta.Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (t *TopicPrereqsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&topicv1alpha1.Topic{}).
		Complete(t)
}
