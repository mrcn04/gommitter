## gommitter

Daily automatic committer for Github.

E.g. use case; the company that you're working at isn't using Github
and you don't want your commit history to look empty.

You can deploy the project on Google Cloud run and create a Cloud Scheduler with the frequency of your choice.

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

#### Deployment on GCP

- Create an Artifact Registry
- Create a service on Cloud Run
- Create a Cloud Scheduler

Dockerize the project

```shell
docker build --tag gommitter-gcp . # add `--platform linux/amd64` if using Apple Silicon
docker run -p 8081:8081 gommitter-gcp # test if its running
docker tag <tag> <region>-docker.pkg.dev/<projectId>/<artifactRepository>/<imageName>
docker push <region>-docker.pkg.dev/<projectId>/<artifactRepository>/<imageName>
```

#### License

This project is under [GPL-3.0 license](https://github.com/mrcn04/gommitter/blob/master/LICENSE) license.
