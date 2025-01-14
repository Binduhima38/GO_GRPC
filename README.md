

```
protoc -Icalculator/proto --go_opt=module=github.com/emohankrishna/go-grpc --go_out=. --go-grpc_opt=module=github.com/emohankrishna/go-grpc --go-grpc_out=. calculator/proto/*.proto

```
Git has local and global settings
global :
    Setting or configs present in global config file has access to every repository

local : 
    Setting or configs present in local config has access to only that particular repository

All remember local config will overwrite the global config

I need to see remote url (Origin) and username and password.

```
git remote add origin https://github.com/emohankrishna/GO_GRPC.git
```

How to add local config
```
git config --local user.email "edemohankrishna1995@gmail.com"
```
when we do `git log` head is pointed to master HEAD -> master

This command will send the committed files to remote
```
git add .
git status
git commit -m "commit message"
git push origin master
```
Different approaches for pushing files to remote
### Approach 1
To establish connection to remote branch

```
git branch B1
git checkout B1
git push --set-upstream origin B1
```
This will work because in the before step i established the connection
```
git push
```

### Approach 2

```
git checkout B2
git push
```
### Approach 3
```
git checkout -b B3
git push origin B3
```
if we do `git push` it won't work because upstream branch is not linked

First check the local device for credential manager if credential are valid.
