# Foo Resource Provider

The Foo Resource Provider lets you manage [AquaSec Cloud](https://cloud.aquasec.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/aquasec
```

or `yarn`:

```bash
yarn add @pulumi/aquasec
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_aquasec
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-aquasec/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Aquasec
```

## Configuration

The following configuration points are available for the `aquasec` provider:

- `aquasec:username` (environment: `AQUA_USERNAME`) - the username used to login
- `aquasec:password` (environment: `AUQA_PASSWORD`) - the password used to login
- `aquasec:aqua_url` (environment: `AQUA_URL`) - Aquasec cloud url `https://cloud.aquasec.com/`

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/foo/api-docs/).
