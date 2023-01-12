# gain

Performance load tool for http(s) servers.
Inspired by [k6](https://github.com/grafana/k6) and [yandex-tank](https://github.com/yandex/yandex-tank).

## ** IN DEVELOPMENT **

<img src="https://github.com/andr-ii/gain/blob/master/assets/gain.png?raw=true"/>

## Installation

Make sure your machine has [go1.19](https://go.dev/doc/install) (or higher) installed before running installation script.

### Linux or MacOS

Simply run:

```bash
curl -s -L 'https://raw.githubusercontent.com/andr-ii/gain/master/scripts/install.sh' | bash
```

Then add following line to your `.bashrc` | `.zshrc` | etc.

```bash
export PATH=$PATH:/home/$USER/gain/bin
```

Terminal **reload is required** after this. Then check current installed version:

```bash
# SHORT
gain -v

# LONG
gain --version
```

## Usage

Create a `plan.json` file (any name could be chosen).

<!-- TODO add docs -->

To learn more abut available plans `see here`(**in development**).

Then start performance testing by running:

```bash
gain ./plan.json # or path to your 'plan'
```
