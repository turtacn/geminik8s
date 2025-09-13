# Configuration Reference

This document provides a complete reference for the `geminik8s` cluster configuration file.

The `geminik8s` cluster is configured using a YAML file. You can generate a default configuration file using the `gemin_k8s init` command.

## Generating the Configuration File

To generate a configuration file, use the `init` command:

```bash
gemin_k8s init \
  --name "my-cluster" \
  --node1-ip "10.10.10.1" \
  --node2-ip "10.10.10.2" \
  --vip "10.10.10.0" \
  --config-dir "./my-cluster-config"
```

This will create a `cluster.yaml` file in the `./my-cluster-config` directory.

## Configuration File Structure

Here is an example of a `cluster.yaml` file with all the fields explained:

```yaml
# API version for the configuration file format.
apiVersion: geminik8s.turtacn.com/v1alpha1
# The kind of configuration object.
kind: ClusterConfig
# Metadata about the cluster.
metadata:
  # The name of the cluster.
  name: my-cluster
# The specification for the cluster.
spec:
  # Network configuration.
  network:
    # The virtual IP (VIP) for the cluster. This IP will be used to access the
    # Kubernetes API server and will float between the two nodes.
    vip: 10.10.10.0
  # The list of nodes in the cluster.
  nodes:
    - # The IP address of the first node.
      ip: 10.10.10.1
      # The role of the node. Can be 'leader' or 'follower'.
      # The 'leader' is the initial master node.
      role: leader
    - # The IP address of the second node.
      ip: 10.10.10.2
      # The role of the node.
      role: follower
  # Configuration for the PostgreSQL database.
  database:
    # The port for the PostgreSQL database.
    port: 5432
    # The user for the PostgreSQL database.
    user: postgres
    # The password for the PostgreSQL database.
    password: "securepassword"
    # The name of the database.
    dbname: kubernetes
  # Advanced configuration options.
  advanced:
    # The version of K3s to install.
    k3sVersion: "v1.28.0+k3s1"
    # The directory where geminik8s will store its state and configuration on the nodes.
    dataDir: "/var/lib/geminik8s"
```

## Next Steps

After configuring your cluster, you can deploy it using the instructions in the [Operation Manual](./operations.md).
