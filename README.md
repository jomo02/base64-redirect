
# Base64Redirect Caddy Module

The `Base64Redirect` module for Caddy is an HTTP handler that encodes incoming request URLs using Base64 encoding and redirects clients to a specified target URL with the encoded URL appended. This can be useful for scenarios where you want to pass URLs as query parameters, protect the original URL structure, or create obfuscated redirects.


## Installation
To use the `Base64Redirect` module, you need to build Caddy with this module included. You can use xcaddy to build a custom Caddy binary with the `Base64Redirect` module.

### Using `xcaddy` to build Caddy with Base64Redirect
```
# Ensure you have xcaddy installed
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest

# Build Caddy with the Base64Redirect module
xcaddy build --with github.com/jomo02/base64-redirect
```

## Usage
To use the `Base64Redirect` module in your Caddy configuration, add a `base64_redirect` directive within your Caddyfile configuration block. The module takes one parameter, `target`, which specifies the base URL to which the encoded URL should be appended.

### Caddyfile Configuration

```
:80 {
    base64_redirect {
        target https://example.com/redirect?url=
    }
}
```

## Configuration Parameters
`target` (string): The base URL to which the Base64-encoded original URL will be appended.

## Example Configuration
### Simple Redirect Example
```
http://yourdomain.com {
    handle {
        base64_redirect {
            target https://example.com/redirect?url=
        }
    }
}
```
In this example, incoming requests to `http://yourdomain.com/somepath` will be redirected to `https://example.com/redirect?url=<Base64-encoded-url>`.

## License
This project is licensed under the MIT License.