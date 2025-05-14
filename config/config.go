package config

import (
    "os"
)

var (
    KubeConfig     = os.Getenv("KUBECONFIG")
    CloudProvider  = os.Getenv("CLOUD_PROVIDER")
    ClusterName    = os.Getenv("CLUSTER_NAME")
    NodeName       = os.Getenv("NODE_NAME")
    ServiceName    = os.Getenv("SERVICE_NAME")
)

