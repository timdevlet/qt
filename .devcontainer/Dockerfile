# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

ARG VARIANT="1.21"
ARG LINT="1.53.3"

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

# Update packages
RUN apt-get update

# [Package] Install golangci-lint 
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

# [Package] Install act
RUN curl -s https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
