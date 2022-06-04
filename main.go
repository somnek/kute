package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubepath := Bite("kube_path")
	kubeconfig := flag.String("kuberconfig", kubepath, "kubeconfig file path")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	// get all available namespace
	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, namespace := range namespaces.Items {
		fmt.Println(namespace.Name)
	}

	// delete all namespace
	for _, namespace := range namespaces.Items {
		err := clientset.CoreV1().Namespaces().Delete(ctx, namespace.Name, metav1.DeleteOptions{})
		if err != nil {
			log.Fatal(err)
		}
	}

}
