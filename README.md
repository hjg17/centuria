# Centuria

## Secure CI/CD Pipeline Demo

This project demonstrates an automated CI/CD pipeline using GitHub Actions to inspect incoming pull requests and prevent security risks from reaching production.

### Security Scanning

The pipeline runs automated checks on every pull request for the following

- **Secrets Detection**
  -  The pipeline uses [Gitleaks](https://github.com/gitleaks/gitleaks) to detect the following secrets
     -  exposed API keys
     -  SSH keys
     -  passwords
- **Vulnerability Scanning**
  -  The pipeline uses [Trivy](https://github.com/aquasecurity/trivy) to identify vulnerabilities across dependencies and base container image layers

If secrets or vulnerabilities are detected, the security check fails and surfaces the findings directly in the pull request and repository security alerts.

### Application & Deployment

The application is a containerized Go API featuring

- A public `/healthz` endpoint for uptime monitoring and health checks
- A protected `/api/v1/data` endpoint secured via an API secret

When all security checks pass, the pipeline automatically builds and publishes the production container image to the [GitHub Container Registry (GHCR)](https://github.com/hjg17/centuria/pkgs/container/centuria).