package utils

import (
	"context"
	"fmt"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// GetKubeClient returns a k8s ClientSet using the in-cluster config
func GetKubeClient() (*kubernetes.Clientset, error) {
	k8sConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	api, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		return nil, err
	}
	return api, nil
}

// ListK8sServicesByLabel returns a k8s service list matching the passed labelSelector in the specified namespace
func ListK8sServicesByLabel(svcLabelSelector, namespace string) (*v1.ServiceList, error) {
	api, err := GetKubeClient()

	if err != nil {
		return nil, fmt.Errorf("could not initialize kubernetes client %w", err)
	}

	svcList, err := api.CoreV1().Services(namespace).List(
		context.TODO(),
		metav1.ListOptions{
			LabelSelector: svcLabelSelector,
		},
	)

	if err != nil {
		return nil, fmt.Errorf(
			"service matching LabelSelector %s in %s namespace is not available. %w", svcLabelSelector, namespace,
			err,
		)
	}
	return svcList, nil
}
