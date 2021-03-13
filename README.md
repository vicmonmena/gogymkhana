# About this project (Go Gymkhana)
Repository for personal gymkhana. This project is developed in Go

## About this project files

- [go.mod](go.mod): dependencies project definitions
- [authd.go](gymkhana.go): Endpoints provider

## Run project

On command line:
```
go run .
```

### Note
`go run` build and run this project

## Build project

On command line:
```
go build
```

## Debug project
On command line:

```
go run gymkhana.go
```

### Debug on VS Code

Add the following configuration to launch.json file:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/gymkhana.go"
    }
  ]
}
```

## About heroku

Find heroku apps [here](https://dashboard.heroku.com/apps/)

### Upload changes and publish
git push heroku-gymkhana main

### Renaming apps
heroku apps:rename azgymkhana 

### Renaming remote
 git remote rename gogymkhana heroku-gymkhana