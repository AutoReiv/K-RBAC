# Kuberus: Simplify Your Kubernetes RBAC Management

Kuberus is the ultimate tool for managing your Kubernetes cluster's security with ease and precision. It offers an intuitive interface that simplifies the complex world of Kubernetes RBAC making it ideal for both seasoned Kubernetes admins and newcomers alike.

## Features

- **Simplified Authorization**: Streamline Kubernetes authorization with an intuitive interface.
- **Effortless Management**: Easily navigate and manage users, groups, and namespaces.
- **Comprehensive Control**: Create, update, and delete roles and bindings effortlessly.
- **Secure and Scalable**: Built to handle the demands of modern Kubernetes environments.

## Getting Started

### Prerequisites

- Docker installed on your machine
- A valid kubeconfig file for your Kubernetes cluster

### Installation and Running

1. Pull the Docker image:

   ```bash
   docker pull reiv/kuberus:latest
   ```

2. Run the application:

   Using Docker Hub image:
   ```bash
   docker run -d -v $HOME/.kube:/root/.kube -p 80:80 reiv/kuberus
   ```

   These commands assume your kubeconfig is in the default location (`$HOME/.kube/config`). If your kubeconfig is elsewhere, adjust the path accordingly.

3. Access the application:
   Open your web browser and navigate to `http://localhost:80`

Note: Kuberus requires a valid kubeconfig to connect to your Kubernetes cluster. If there is no valid kubeconfig available, the container will stop.

## Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
