# gh-labeler
`gh-labeler` is a Golang package which lets you programmatically define more
useful labels for your Github issues, on a per-repository basis.

## Installation
Installing gh-labeler is as simple as running `go get` like so:

```bash
$ go get github.com/moogar0880/gh-labeler/cmd/ghlabels
```

## Docker Image
You can also pull down and run the publicly available docker image:

```bash
$ docker run --rm moogar0880/ghlabeler -token ${MY_ACCESS_TOKEN}
```

## Configuration
Before you can start using this tool, you need a
[personal Github Access Token](https://github.com/settings/tokens) with proper
repository permissions.

## Usage
To run ghlabels once it's installed (assuming `${GOPATH}/bin` is on your `$PATH`)
you can run

```bash
$ ghlabels -h
Define more useful labels for your Github issues.

Usage:
  ghlabels [OPTIONS] COMMAND [arg...] [flags]

Flags:
  -f, --file string    Specify Config File to Load (default "labels.json")
  -r, --remove         Remove labels that are not present in the config file
  -t, --token string   The Github Access Token to use
  -v, --version        Print version information and quit
```

## Extended Usage
As of 0.2.0, you may now specify a list of repositories to apply the label
configuration to. You may list out multiple repositories like so:
```json
{
  "owner": "moogar0880",
  "repos": [
    "gh-labeler",
    "PyTrakt"
  ],
  "labels": [
    {
      "name": "bug",
      "color": "ee0701"
    }
  ]
}
```

## Default Labels
To help get you started, the default Github Issues have been defined in the
provided [default.json](default.json).
