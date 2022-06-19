# cloudflare

[![Go Reference](https://pkg.go.dev/badge/go.felesatra.moe/cloudflare.svg)](https://pkg.go.dev/go.felesatra.moe/cloudflare)

Package cloudflare provides a client for the Cloudflare API v4.

This is not an official package.

This package provides a simple client and functions that implement
simple functionality like updating a single A record.

This package handles authentication and basic marshaling of
requests/unmarshaling of responses.  You will need to reference the
API docs and do your own method calls for anything beyond the
provided basics.
