package main

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListNamespaces() ([]corev1.Namespace, error) {
	ctx := context.TODO()
	namespaces, err := c.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces.Items, nil
}

func (c *Client) ListServices(namespace string) ([]corev1.Service, error) {
	// namespace "" means all namespaces
	ctx := context.TODO()
	services, err := c.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return services.Items, nil
}

func (c *Client) GetService(namespace, name string) (*corev1.Service, error) {
	ctx := context.TODO()
	services, err := c.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return services, nil
}
