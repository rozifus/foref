# fref
Quickly collect public git repositories into a reference folder from the command line

## Installation (ubuntu)
Build fref

`go build -o fref ./cmd/fref/*`

Copy fref to system

`sudo cp ./fref /usr/local/bin`

## Setup (ubuntu)
Choose a root directory for saved repos

`export FREF_NAMESPACE_DEFAULT=/home/yourusername/Desktop/ReferenceCode`

## Usage

```
Usage: fref <identifier> ...

Arguments:
  <identifier> ...

Flags:
  -h, --help                   Show context-sensitive help.
  -n, --namespace="DEFAULT"    Which folder namespace to use.
  -s, --source="github"
```

## Usage Examples

Github Repo

`fref -s github repo geerlingguy/ansible-for-devops`

Github User (all public repos)

`fref -s github user geerlingguy`

Gitlab User (all public repos)

`fref -s gitlab user inkscape`

Url (auto match)

`fref https://gitlab.com/inkscape/devel/chat-utils`

`fref https://github.com/geerlingguy/ansible-for-devops`

## Namespaces

You can have multiple root download directories

Setting a custom environment variable / namespace

`export FREF_NAMESPACE_MYCUSTOM=/home/yourusername/Deskptop/MyCustomReferenceCode`

Passing the namespace flag

`fref --namespace mycustom ...`

eg

`fref --namespace mycustom --source github user rozifus`

