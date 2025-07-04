# ArgoCD Sync Actions

[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-Find%20and%20Replace-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/marketplace/actions/argocd-sync-app)
[![Actions Status](https://github.com/flxxyz/argocd-sync-app/actions/workflows/test.yml/badge.svg)](https://github.com/flxxyz/argocd-sync-app/actions/workflows/test.yml)
[![Actions Status](https://github.com/flxxyz/argocd-sync-app/actions/workflows/integration.yml/badge.svg)](https://github.com/flxxyz/argocd-sync-app/actions/workflows/integration.yml)

This action will sync ArgoCD application.

## Usage

### Example workflow

This example replaces syncs ArgoCD application.

```yaml
name: My Workflow
on: [ push, pull_request ]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Sync ArgoCD Application
        uses: flxxyz/argocd-sync-app@main
        with:
          address: "vault.example.com"
          token: ${{ secrets.ARGOCD_TOKEN }}
          insecure: false
          appName: "my-example-app"
          plaintext: false
```

### Inputs

| Input | Description|
| --- | --- |
| `address` | ArgoCD server address. |
| `token` | ArgoCD Token. |
| `insecure` | ArgoCD insecure. |
| `appName` | Application name to sync. |
| `plaintext` | use http instead of https |

## Examples

### Sync Application

You can sync ArgoCD application after building an image etc.

```yaml
name: My Workflow
on: [ push, pull_request ]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Sync ArgoCD Application
        uses: flxxyz/argocd-sync-app@main
        with:
          address: "vault.example.com"
          token: ${{ secrets.ARGOCD_TOKEN }}
          insecure: false
          appName: "my-example-app"
          plaintext: false
```

## Publishing

To publish a new version of this Action we need to update the Docker image tag in `action.yml` and also create a new
release on GitHub.

- Work out the next tag version number.
- Update the Docker image in `action.yml`.
- Create a new release on GitHub with the same tag.
