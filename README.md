# detour

deTour is a small proxy server which redirects.

Source : `https://github.com/pulsiot/detour`

## Build

[![Build Containerimage](https://github.com/pulsiot/detour/actions/workflows/build-containerimage.yml/badge.svg)](https://github.com/pulsiot/detour/actions/workflows/build-containerimage.yml)

## Config File.

deTour uses two configuration files:

- `config.yaml` file to read the config parameters.
- `routes.yaml` file to read the routes mappings.

## Dependencies

Please note that the program assumes you have the necessary config.yaml, routes.yaml, server.crt, and server.key files.

## Create SSL certificates

```
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -sha256 -days 3650 -nodes -subj "/C=IN/ST=Kerala/L=Kochi/O=Detour/OU=DetourProxy/CN=localhost"
```

Or

```
openssl req -x509 -new -newkey rsa:4096 -sha256 -days 3650 -nodes -out server.crt -keyout server.key

```
