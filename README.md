
# Centuria

## Secure CI/CD Pipeline Demo



This demo showcases a secure CI/CD pipeline.

Specifically, it looks at pull requests that might contain secrets, or sensitive data such as the following

- SSH keys
- API keys
- Passwords


Once a key is found, an alert is created to notify the user of the issue and prevents the pull request from merging. This is done by using [Gitleaks](https://github.com/gitleaks/gitleaks), an open-source tool created to detect passwords and API keys.

The project itself a deployed docker