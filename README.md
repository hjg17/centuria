
# Centuria

## Secure CI/CD Pipeline Demo



This demo showcases a secure CI/CD pipeline.

Specifically, it looks at pull requests that might contain secrets, or sensitive data such as the following

- SSH keys
- API keys
- Passwords


Once a key is found, an alert is created to notify the developer of the issue and prevents the pull request from merging. This is done by using [Gitleaks](https://github.com/gitleaks/gitleaks), and [Trivy](https://github.com/aquasecurity/trivy) to detect container and dependency vulnerabilities.

The project is a containerized Go API deployed to the [GitHub Container Registry (GHCR)](https://github.com/hjg17/centuria/pkgs/container/centuria), featuring a public `/healthz` health check endpoint and a protected `/api/v1/data` endpoint secured via an API secret. If all security scans pass on the main branch, the pipeline automatically builds and publishes the container image for production use.