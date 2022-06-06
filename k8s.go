package main

import (
	"flag"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	*kubernetes.Clientset
}

func NewClient() (*Client, error) {
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

	return &Client{
		Clientset: clientset,
	}, nil
}
