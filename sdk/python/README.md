# Harbor Resource Provider

The Harbor Resource Provider lets you manage [Harbor](https://goharbor.io) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/harbor
```

or `yarn`:

```bash
yarn add @pulumiverse/harbor
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumiverse-harbor
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-harbor/sdk/v3
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumiverse.Harbor
```

## Configuration

The following configuration points are available for the `harbor` provider:

- `harbor:url` - (Required) The url of the harbor.
- `harbor:username` - (Required) The username to be used to access harbor.
- `harbor:password` - (Required) The password to be used to access harbor.
- `harbor:insecure` - (Optional) Choose to ignore certificate errors
- `harbor:apiVersion` - (Optional) Choose which version of the api you would like to use `1` or `2`. Default is `2`.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/harbor/api-docs/).
