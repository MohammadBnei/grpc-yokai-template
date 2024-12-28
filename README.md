# Grpc Yokai Template

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go version](https://img.shields.io/badge/Go-1.23-blue)](https://go.dev/)
[![Documentation](https://img.shields.io/badge/Doc-online-cyan)](https://ankorstore.github.io/yokai/)

A customizable gRPC application template based on the [Yokai](https://github.com/ankorstore/yokai) Go framework, now including Kubernetes deployment, GitHub Actions workflows, and a Dockerized Postgres database.

## Table of Contents

- [Grpc Yokai Template](#grpc-yokai-template)
  - [Table of Contents](#table-of-contents)
  - [Documentation](#documentation)
  - [Getting started](#getting-started)
    - [Installation](#installation)
      - [With GitHub](#with-github)
      - [With gonew](#with-gonew)
    - [Usage](#usage)
  - [Features](#features)

## Documentation

For more information about the [Yokai](https://github.com/ankorstore/yokai) framework, you can check its [documentation](https://ankorstore.github.io/yokai).

## Getting started

### Installation

#### With GitHub

You can create your repository using this template: <https://github.com/new?template_name=grpc-yokai-template&template_owner=MohammadBnei>

It will automatically rename your project resources and push them. This operation may take a few minutes.

Once ready, after cloning and going into your repository, simply run:

```shell
make fresh
```

#### With gonew

You can install [gonew](https://go.dev/blog/gonew), and simply run:

```shell
gonew github.com/MohammadBnei/grpc-yokai-template github.com/foo/bar
cd bar
make fresh
```

### Usage

This template provides a basic structure for building a gRPC application with Yokai. You can customize it to fit your needs.

## Features

*   **Kubernetes Deployment**: Includes a `kustomize` folder for easy Kubernetes deployment.
*   **GitHub Actions Workflows**:
    *   Release automation
    *   Changelog generation
    *   Docker build and push to public registry (e.g., Docker Hub)
*   **GORM Integration**: Uses the GORM module from Yokai for database operations
*   **Dockerized Postgres Database**: Includes a Compose file with a Postgres database for development and testing purposes
