# Concourse Quickstart

## Prequisites

### Install fly

easiest way is to use brew

```sh
brew install fly
```

### Login data

- `user`: test
- `password`: test

The target or environment is called `tutorial`

## How to start

Run the docker compose file

```sh
docker-compose up
```

Then login

```sh
fly -t tutorial login -c http://localhost:8080 -u test -p test
```

You will be prompted to open a url in the browser. There login with the credentials (see above)

Now you should be ready to install your first pipeline. We will be using the `pipeline.yml` for that:

```sh
fly -t tutorial sync
fly -t tutorial set-pipeline --pipeline my-pipe --config pipeline.yml
fly -t tutorial unpause-pipeline -p my-pipe
```

1. fly sync is to sync the fly version with the target (might be already fine)
2. set-pipeline will install the pipeline.yml to concourse
3. newly set pipelines are paused automatically, so unpause it. (you can unpause it in the concourse dashboard as well)

> **🤓**
> For convenience you could put the steps above in a script and just run that. See [fly_pipe.sh](./fly_pipe.sh) for example.

## Try it out

Now everything should be set up. In your browser navigate to [Localhost on port 8080](http://localhost:8080)
You should find your pipeline in the concourse dashboard. Try to trigger it manually by clicking on the ⊕ button in the pipeline detail view.
