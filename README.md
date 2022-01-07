# foref
Quickly collect public git repositories into a reference folder from the command line

## Installation (ubuntu)
Build foref

`go build -o foref ./cmd/foref/*`

Copy foref to system

`sudo cp ./foref /usr/local/bin`

## Setup (ubuntu)
Choose a root directory for saved repos

`export FOREF_NAMESPACE_DEFAULT=/home/yourusername/Desktop/ReferenceCode`

## Usage

```
Usage: foref <identifier> ...

Arguments:
  <identifier> ...

Flags:
  -h, --help                   Show context-sensitive help.
  -n, --namespace="DEFAULT"    Which folder namespace to use.
  -s, --source="github"
```

## Usage Examples

Github Repo

`foref -s github repo geerlingguy/ansible-for-devops`

Github User (all public repos)

`foref -s github user geerlingguy`

Gitlab User (all public repos)

`foref -s gitlab user inkscape`

Url (auto match)

`foref https://gitlab.com/inkscape/devel/chat-utils`

`foref https://github.com/geerlingguy/ansible-for-devops`

## Namespaces

You can have multiple root download directories

Setting a custom environment variable / namespace

`export FOREF_NAMESPACE_MYCUSTOM=/home/yourusername/Deskptop/MyCustomReferenceCode`

Passing the namespace flag

`foref --namespace mycustom ...`

eg

`foref --namespace mycustom --source github user rozifus`

