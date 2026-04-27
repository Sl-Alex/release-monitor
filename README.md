# Release Monitor

Lightweight version checker written in TypeScript + Deno.

Checks installed software versions against remote sources such as:

- HTML pages
- GitHub releases

Portable-friendly:

- local Deno import mirror support
- bundled Deno + gh

## Prerequisites:

1. `libs` folder must be added to the environment as DENO_LIB_PATH
2. `tools` folder must be added to the PATH variable, so that both `deno` and `gh` tools are available

Run `install.bat` once, it will download Deno executable and unzip it to the `tools` folder. The executable is too big to include it to this repository.

## Structure

```
/libs/
  deno_dom@v0.1.56/

/tools/
  deno.exe
  gh.exe

/src/
  common.ts
  main.ts

apps.json - a list of installed apps/versions to check
deno.json - deno configuration allowing to fetch imports before moving them to the libs folder
```

### apps.json

```json
[
  {
    "name": "AMAP",
    "installed": "0.30",
    "type": "html",
    "url": "https://www.sikorskiy.net/info/prj/amap/index.html",
    "selector": "#download + ul > li:nth-child(2) ul li a",
    "transform": "split:v"
  },
  {
    "name": "ConEmu",
    "installed": "v23.07.24",
    "type": "github",
    "repo": "ConEmu/ConEmu"
  },
  {
    "name": "CurrPorts",
    "installed": "v2.77",
    "type": "html",
    "url": "https://www.nirsoft.net/utils/cports.html",
    "selector": ".utilcaption tbody tr :nth-child(2)",
    "transform": "regex:v\\d+\\.\\d+"
  }
]
```

## Run 

The following command will take the `apps.json` and check if there are any updates available:

```bash
deno run --allow-all src/main.ts
```

Sample output:

```bash
$ deno run --allow-all src/main.ts
> AMAP: update available (0.30 → 0.34)
  ConEmu: up to date (v23.07.24)
  CurrPorts: up to date (v2.77)
```

## Local Deno Imports

Imports can be mirrored locally inside `libs` folder.
By default, `deno_dom` is already included. If you need other imports and would like to keep them locally, then take a look at `src/common.ts`, there is an example.

## License

| Tool            | URL                                        | License   |
|-----------------|--------------------------------------------|-----------|
| Deno            | https://github.com/denoland/deno           | MIT       |
| GitHub CLI      | https://github.com/cli/cli                 | MIT       |
| Release Monitor | https://github.com/Sl-Alex/release-monitor | Unlicense | 
