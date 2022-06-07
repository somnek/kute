package main

import (
	"context"
	"fmt"

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

func (c *Client) DeleteNamespace(namespace, name string) error {
	ctx := context.TODO()
	if namespace == "default" || namespace == "kube-system" {
		return fmt.Errorf("you can't delete default / kube-system namespaces you fucking moron")
	}
	gracePeriodSeconds := int64(0)
	return c.CoreV1().Namespaces().Delete(ctx, namespace, metav1.DeleteOptions{GracePeriodSeconds: &gracePeriodSeconds})
}

func (c *Client) GetService(namespace, name string) (*corev1.Service, error) {
	ctx := context.TODO()
	services, err := c.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (c *Client) CreateNamespace(namespace string) error {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
			Labels: map[string]string{
				"created_by": "Kute! 💮",
			},
		},
	}
	_, err := c.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	return err
}
