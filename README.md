# todocli
# Run Go in Docker
```shell
docker run -it -v ${PWD}:/work golang sh
go version
```
# Using CLI
```
# get all todos
todocli get --all

# get done todos
todocli get --done

# get due todos
todocli get --due

# get todo by ID
todocli get --id <todo-id>

# add todo to the list
todocli add -title

# mark todo as done by ID
todocli done -id
```

# Using prod image

```
# build prod image
docker build . -t todocli

# use prod image
docker run -it todocli get -all 
```