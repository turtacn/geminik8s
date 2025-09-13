# Troubleshooting Guide

This document provides solutions to common issues you may encounter when using `geminik8s`.

## Deployment Fails

If the `deploy` command fails, here are some steps to troubleshoot the issue:

1.  **Check the logs:**
    Run the `deploy` command with a higher log level to get more detailed information:
    ```bash
    gemin_k8s deploy --config-dir "./my-cluster-config" --log-level "debug"
    ```
    The logs will be printed to the console. Look for any error messages.

2.  **Check node connectivity:**
    Ensure that the two nodes can communicate with each other over the network. You can use `ping` to check connectivity:
    ```bash
    ping <other-node-ip>
    ```

3.  **Check SSH access:**
    `geminik8s` uses SSH to connect to the nodes. Ensure that you have passwordless SSH access from the machine where you are running `gemin_k8s` to both nodes.

## Cluster is Unhealthy

If the `gemin_k8s status` command shows the cluster as "Unhealthy", here are some things to check:

1.  **Check the status of the K3s service:**
    Log in to each node and check the status of the `k3s` service:
    ```bash
    systemctl status k3s
    ```

2.  **Check the status of the PostgreSQL service:**
    Log in to each node and check the status of the `postgresql` service:
    ```bash
    systemctl status postgresql
    ```

3.  **Check the database replication:**
    Log in to the leader node and check the status of the database replication.
    *(More detailed instructions to come)*

## Getting Help

If you are still unable to resolve the issue, you can get help from the community:

*   **[GitHub Discussions](https://github.com/turtacn/geminik8s/discussions)**: Ask questions and get help from the community.
*   **[GitHub Issues](https://github.com/turtacn/geminik8s/issues)**: Report bugs and request features.
