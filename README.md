# Silo Comic Pages

Extracts comic archives on the Silo server and serves their pages to the [Silo Aidoku source](https://github.com/crowquillx/aidoku-silo-sources).

Part of [Crowquillx Silo Plugins](https://github.com/crowquillx/crowquillx-silo-plugins#install). Supports Linux amd64 and arm64.

## Install

1. Add the [Crowquillx catalog](https://github.com/crowquillx/crowquillx-silo-plugins#install) in **Administration → Plugins → Catalog** and install **Comic Pages**.
2. In **Installed**, open the plugin's **Configure** button or settings gear, then **Global Configuration**.
3. Set **Silo base URL** to an address reachable from the plugin, including Silo's HTTP port. Set **Private cache directory** to a dedicated writable directory outside your media library, then select **Save config**.
4. Install or update the [Silo Aidoku source](https://github.com/crowquillx/aidoku-silo-sources) to v11 or newer and sign in to Silo. It detects Comic Pages automatically for CBR reading.

The cache defaults to 2 GiB and is cleared when the plugin starts or reconfigures. Give each installation its own cache directory. The first chapter open may take longer while its pages are extracted.

## Build

Requires Go 1.26 or newer.

```sh
make test
make build
```

The binary is `bin/plugin`.

## Acknowledgments

- [rardecode](https://github.com/nwaples/rardecode) for RAR extraction.
- [Go image libraries](https://pkg.go.dev/golang.org/x/image) for WebP decoding.
- [Silo plugin SDK](https://github.com/Silo-Server/silo-plugin-sdk) for Silo integration.

[MIT license](LICENSE). Dependency licenses are in [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).
