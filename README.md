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

## Usage Examples

Github Repo

`gitt github repo geerlingguy/ansible-for-devops`

Github User (all public repos)

`gitt github user geerlingguy`

Gitlab User (all public repos)

`gitt gitlab user inkscape`

Url (auto match)

`gitt url https://gitlab.com/inkscape/devel/chat-utils`

`gitt url https://github.com/geerlingguy/ansible-for-devops`

## Namespaces
You can user multiple alternative root download directories by:

Setting another custom environment variable

`export GITT_NAMESPACE_MYCUSTOM=/home/yourusername/Deskptop/MyCustomReferenceCode`

Passing the namespace flag

`gitt --namespace mycustom ...`

eg

`gitt --namespace mycustom github user rozifus`
