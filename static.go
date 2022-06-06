package main

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListServices(namespace string) ([]corev1.Service, error) {
	if namespace == "" {
		namespace = "default"
	}
	ctx := context.TODO()
	services, err := c.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return services.Items, nil
}
