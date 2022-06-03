package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/apis/metav1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kuberconfig", "xxx", "kubeconfig file path")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background
	pods, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods {
		fmt.Println(pod.Name)
	}

}
