## gommitter

Daily automatic committer for Github.

E.g. use case; the company that you're working at isn't using Github
and you don't want your commit history to look empty. You can adjust the daily commit count when creating
the GCP job.

### Running the App

```
make run
```

#### Required Env Variables

```shell
GH_KEY=  # Github personal access token
PORT=
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
