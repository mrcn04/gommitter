## gommitter

Daily automatic committer for Github.

### Running the App

```
make run
```

#### Required Env Variables

```shell
GH_KEY=  # Github personal access token
```

#### Configuration

`./config.go`

```golang
User   = "mrcn04"
Repo   = "gommitter"
Branch = "master"
Type   = "commit"
Ref    = "refs/heads/master"
```

#### License

This project is under [GPL-3.0 license](https://github.com/mrcn04/gommitter/blob/master/LICENSE) license.
