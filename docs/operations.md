# Operation Manual

This document provides a guide for day-2 operations of a `geminik8s` cluster.

## Checking Cluster Status

To check the status of your cluster, use the `status` command:

```bash
gemin_k8s status --cluster "my-cluster"
```

The output will show the health of the cluster, the current leader, and the status of the nodes.

## Deploying the Cluster

To deploy the cluster, use the `deploy` command with the path to your configuration directory:

```bash
gemin_k8s deploy --config-dir "./my-cluster-config"
```

## Manual Failover

In the event of a planned maintenance or if you need to manually switch the leader node, you can use the `failover` command:

```bash
gemin_k8s failover --cluster "my-cluster" --promote "node2"
```

This will promote `node2` to be the new leader.

## Upgrading the Cluster

To upgrade the Kubernetes version of your cluster, use the `upgrade` command:

```bash
gemin_k8s upgrade --cluster "my-cluster" --version "v1.29.0+k3s1"
```

## Backing Up and Restoring the Cluster

To back up the cluster's database, use the `backup` command:

```bash
gemin_k8s backup --cluster "my-cluster" --output "/backups/backup-$(date +%F).sql"
```

To restore from a backup, use the `restore` command:

```bash
gemin_k8s restore --cluster "my-cluster" --backup-file "/backups/backup-2023-10-27.sql"
```

**Note:** The `backup` and `restore` commands are currently under development.

## Replacing a Node

If a node fails and needs to be replaced, you can use the `replace-node` command:

```bash
gemin_k8s replace-node --cluster "my-cluster" --old-node "node2" --new-node-ip "10.10.10.3"
```

**Note:** The `replace-node` command is currently under development.
