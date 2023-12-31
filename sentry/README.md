# sentry

Package provides fx.Module for sentry
## Configuration

Global

- `DEBUG`

Expected environments

- `SENTRY_DSN`
- `SENTRY_SAMPLE_RATE`
- `SENTRY_FLUSH`
- `SENTRY_RELEASE_VERSION`
- `SENTRY_RELEASE_ENVIRONMENT`

## Deps

- `build` - optional, use build.Context for Sentry release variables, such as a Version, Environment and Service
- `config`
