# Go Template

A personal template with the following dependencies:
- [`zerolog`](https://github.com/rs/zerolog)
- [`viper`](https://github.com/spf13/viper)
- [`echo`](https://github.com/labstack/echo)

## Setup
1. Clone this repo
2. Replace instances of `go-template` in code with that of your actual project.
3. Update `SERVICE` const in `internal/config/configKeys` to name of your project.
4. Add config fields via prefixing env vars with your service name set in previous step (eg. `SERVICE_KEY=VALUE`)
5. Add routes under `internal/server` if desired.
6. Add utils under `internal/utils` if desired.
7. Add go code, dependencies as needed.

## To Add:
- [ ] Dockerfile declaration
- [ ] docker compose declaration
- [ ] `.env.template` declaration
- [ ] `air` or other hot reload support