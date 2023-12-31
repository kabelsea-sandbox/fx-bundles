# resty

Package provides fx.Module for http client macking API request via `go-resty`.

## Configuration

Global

- `DEBUG`

Expected

- `RESTY_TRACE`
- `RESTY_TIMEOUT`

- `RESTY_URL`           - base url
- `RESTY_USER_AGENT`    - default gatefi/1.0.0
- `RESTY_HEADERS`       - optional

Retry policies

- `RESTY_RETRY_COUNT`   - default is zero

Basic auth, optional

- `RESTY_AUTH_BASIC_USERNAME`
- `RESTY_AUTH_BASIC_PASSWORD`

Token auth, optional

- `RESTY_AUTH_TOKEN_SCHEME` - default, bearer
- `RESTY_AUTH_TOKEN_VALUE`

## Deps

- `logger`
- `config`
- `monitoring`
