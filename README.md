# geminik8s

<img src="logo.png" alt="geminik8s Logo" width="200" height="200">

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/github/go-mod/go-version/turtacn/geminik8s)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/turtacn/geminik8s)](https://github.com/turtacn/geminik8s/releases)

> **A cost-effective dual-node high availability Kubernetes solution for resource-optimized deployments**

[中文版 README](./README-zh.md) | [Architecture Documentation](./docs/architecture.md)

## 🚀 Mission Statement

geminik8s revolutionizes Kubernetes deployments by providing a cost-effective dual-node high availability solution that reduces infrastructure costs by up to 33% compared to traditional three-node clusters, while maintaining enterprise-grade reliability and performance.

## 🎯 Why geminik8s?

**Traditional Kubernetes HA Pain Points:**
- 🔸 **High Cost Barrier**: Traditional HA requires minimum 3 nodes due to etcd quorum requirements
- 🔸 **Resource Waste**: Over-provisioning for basic high availability needs
- 🔸 **Complexity Overhead**: Managing etcd clusters adds operational burden
- 🔸 **Vendor Lock-in**: Limited options for alternative storage backends

**geminik8s Advantages:**
- ✅ **Cost Optimization**: 33% reduction in hardware, power, and operational costs
- ✅ **Simplified Architecture**: PostgreSQL + Kine replaces complex etcd management
- ✅ **Production Ready**: Battle-tested solution for cost-sensitive deployments
- ✅ **Kubernetes Native**: Full compatibility with standard K8s/K3s distributions
- ✅ **Automatic Failover**: Intelligent leader election and seamless node promotion
- ✅ **Zero-Downtime Upgrades**: A/B partition strategy with automatic rollback

## ✨ Key Features

### 🏗️ Architecture & Infrastructure
- **Dual-Node HA**: Cost-effective alternative to traditional 3-node clusters
- **PostgreSQL Backend**: Leveraging robust SQL databases instead of etcd
- **Kine Integration**: Seamless etcd-to-SQL translation layer
- **A/B Partitioning**: Immutable OS images with automatic rollback capability

### 🔄 Lifecycle Management  
- **Intelligent Orchestration**: Automated leader election and failover
- **Health Monitoring**: Comprehensive liveness checks and recovery mechanisms
- **Upgrade Automation**: Zero-downtime cluster upgrades with rollback support
- **Node Replacement**: Seamless replacement of failed nodes

### 🛠️ Developer Experience
- **CLI-First Design**: Intuitive command-line interface powered by Cobra
- **Plugin Architecture**: Extensible system for custom functionality
- **Configuration as Code**: Declarative cluster configuration management
- **Multi-Platform Support**: Cross-platform binary distribution

### 🔒 Enterprise Features
- **Production Grade**: Designed for enterprise reliability requirements
- **Observability**: Built-in monitoring and logging capabilities
- **Security**: Standard Kubernetes RBAC and security policies
- **Backup & Recovery**: Automated backup strategies for data protection

## 🚀 Getting Started

### Installation

Install geminik8s using Go:

```bash
go install github.com/turtacn/geminik8s/cmd/gemin_k8s@latest
````

Or download pre-built binaries from [releases](https://github.com/turtacn/geminik8s/releases).

### Quick Start

#### 1. Initialize a new dual-node cluster configuration

```bash
# Create cluster configuration
gemin_k8s init --name "my-cluster" \
  --node1-ip "10.10.10.1" \
  --node2-ip "10.10.10.2" \
  --vip "10.10.10.0" \
  --config-dir "./cluster-config"
```

#### 2. Deploy the cluster

```bash
# Deploy to both nodes
gemin_k8s deploy --config-dir "./cluster-config" \
  --bootstrap-leader "node1"
```

#### 3. Verify cluster status

```bash
# Check cluster health
gemin_k8s status --cluster "my-cluster"

# Sample output:
# Cluster: my-cluster
# Status: Healthy
# Leader: node1 (10.10.10.1)
# Follower: node2 (10.10.10.2)
# VIP: 10.10.10.0 (Active on node1)
# Database: PostgreSQL + Kine (Replication: Active)
```

#### 4. Advanced operations

```bash
# Trigger manual failover
gemin_k8s failover --cluster "my-cluster" --promote "node2"

# Upgrade cluster
gemin_k8s upgrade --cluster "my-cluster" \
  --image "my-registry/k8s-image:v1.28.0"

# Replace failed node
gemin_k8s replace-node --cluster "my-cluster" \
  --old-node "node2" --new-node-ip "10.10.10.3"
```

## 📖 Documentation

* [Architecture Overview](./docs/architecture.md) - Detailed system architecture and design decisions
* [Installation Guide](./docs/installation.md) - Comprehensive installation instructions
* [Configuration Reference](./docs/configuration.md) - Complete configuration options
* [Operation Manual](./docs/operations.md) - Day-2 operations guide
* [Troubleshooting](./docs/troubleshooting.md) - Common issues and solutions

## 🤝 Contributing

We welcome contributions from the community! geminik8s is built with ❤️ by engineers who understand the real-world challenges of Kubernetes operations.

### Ways to Contribute

* 🐛 **Report Issues**: Found a bug? [Create an issue](https://github.com/turtacn/geminik8s/issues)
* 💡 **Feature Requests**: Have an idea? We'd love to hear it!
* 🔧 **Code Contributions**: Submit pull requests for bug fixes or new features
* 📝 **Documentation**: Help improve our docs and examples
* 🧪 **Testing**: Help us test on different environments

### Development Setup

```bash
# Clone the repository
git clone https://github.com/turtacn/geminik8s.git
cd geminik8s

# Install dependencies
go mod tidy

# Run tests
make test

# Build binary
make build
```

See [CONTRIBUTING.md](./CONTRIBUTING.md) for detailed contribution guidelines.

## 📜 License

geminik8s is licensed under the [Apache License 2.0](LICENSE). This means you can:

* ✅ Use it commercially
* ✅ Modify and distribute
* ✅ Use it privately
* ✅ Include it in patents

## 🌟 Community & Support

* 💬 **Discussions**: [GitHub Discussions](https://github.com/turtacn/geminik8s/discussions)
* 🐛 **Issues**: [Bug Reports & Feature Requests](https://github.com/turtacn/geminik8s/issues)
* 📧 **Contact**: [maintainers@geminik8s.io](mailto:maintainers@geminik8s.io)
* 🔗 **Website**: [https://geminik8s.io](https://geminik8s.io)

---

**Built with ❤️ for the Kubernetes community**

*Reducing costs shouldn't mean compromising on reliability. geminik8s proves that smart architecture can deliver both.*
