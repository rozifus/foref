# gitt
Quickly collect public git repositories into a reference folder from the command line

## Installation (ubuntu)
Build gitt

`go build -o gitt ./cmd/gitt/*`

Copy gitt to system

`sudo cp ./gitt /usr/local/bin`

## Setup (ubuntu)
Choose a root directory for saved repos

`export GITT_NAMESPACE_DEFAULT=/home/yourusername/Desktop/ReferenceCode`

## Usage

```
Usage: gitt <identifier> ...

Arguments:
  <identifier> ...

Flags:
  -h, --help                   Show context-sensitive help.
  -n, --namespace="DEFAULT"    Which folder namespace to use.
  -s, --source="github"
```

## Usage Examples

Github Repo

`gitt -s github repo geerlingguy/ansible-for-devops`

Github User (all public repos)

`gitt -s github user geerlingguy`

Gitlab User (all public repos)

`gitt -s gitlab user inkscape`

Url (auto match)

`gitt https://gitlab.com/inkscape/devel/chat-utils`

`gitt https://github.com/geerlingguy/ansible-for-devops`

## Namespaces
You can have multiple root download directories

Setting a custom environment variable / namespace

`export GITT_NAMESPACE_MYCUSTOM=/home/yourusername/Deskptop/MyCustomReferenceCode`

Passing the namespace flag

`gitt --namespace mycustom ...`

eg

`gitt --namespace mycustom --source github user rozifus`

