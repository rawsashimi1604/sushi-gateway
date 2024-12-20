# Contributor Guide

Thank you for your interest in contributing to Sushi Gateway! 🎉 

We value all contributions, no matter how big or small. Whether it’s bug reports, fixes, documentation improvements, or new examples—your efforts are appreciated! Before diving in, please take a moment to read through this guide to ensure smooth collaboration.

## Table of Contents

- [Reporting an Issue](#reporting-an-issue)
- [Before You Contribute](#before-you-contribute)
  - [Code Reviews](#code-reviews)
  - [Coding Guidelines](#coding-guidelines)
  - [Continuous Integration](#continuous-integration)
  - [Tests and Documentation](#tests-and-documentation)
- [The Fine Print](#the-fine-print)

---

## Reporting an Issue

We use GitHub Issues to track bugs, feature requests, and other tasks. When reporting an issue:

1. Clearly describe the problem.
2. Provide steps to reproduce it.
3. Share what you observed and what you expected to happen.
4. Attach logs or screenshots, if applicable.

Help us help you by making your report as detailed as possible!

---

## Before You Contribute

Contributions are made via GitHub Pull Requests (PRs) from your **own fork** of the repository. Ensure your Git author details are correctly configured:

```bash
git config --global user.name "Your Full Name"
git config --global user.email your.email@example.com
```

This ensures you receive proper credit for your contributions. If you use multiple machines, make sure your Git configuration is consistent across them.

---

### Code Reviews

All submissions, including those by project maintainers, require at least one approval from a Sushi Gateway committer before merging. 

We follow the [GitHub Pull Request Review Process](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/reviewing-changes-in-pull-requests/about-pull-request-reviews). Please address all feedback promptly and ensure your PR is clear and concise.

---

### Coding Guidelines

Adhere to the following coding best practices:

- Write clean, modular, and readable code.
- Follow established patterns and conventions used in the repository.
- Include comments where necessary, especially for complex logic.

---

### Continuous Integration

To ensure stability, all changes must pass our CI workflows before merging. 

Sushi Gateway uses GitHub Actions for CI. When you submit a PR, automated tests and checks will run. Monitor the status of these workflows and resolve any failures promptly.

---

### Tests and Documentation

Tests and documentation are **not optional**:

- **Tests**: Include unit tests or integration tests relevant to your changes.
- **Documentation**: Update the Vitepress documentation located in the `/docs` directory to reflect your changes.

Remember, well-documented code and robust tests make everyone's life easier.

---


## The Fine Print

Sushi Gateway is an open-source project, and we expect all contributors to adhere to these principles:

1. **Be Respectful**: Engage with others politely and constructively.
2. **Act Responsibly**: Submit well-tested and high-quality changes.
3. **Enjoy**: Open source is a collaborative and rewarding experience—have fun!

We look forward to your contributions and thank you for making Sushi Gateway even better! 🎉
