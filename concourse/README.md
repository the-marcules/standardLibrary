# Concourse Quickstart

## Prequisites

### install fly

easiest way is to use brew

```sh
brew install fly
```

### Login data

- `user`: test
- `password`: test

the target or environment is called `tutorial`

## How to start

run the docker compose file

```sh
docker-compose up
```

to login run

```sh
fly -t tutorial login -c http://localhost:8080 -u test -p test
```

you will be prompted to open a url in the browser. there login with the credentials (see above)

then you are good to install your first pipeline:

```sh
fly -t tutorial sync
fly -t tutorial set-pipeline --pipeline my-pipe --config pipeline.yml
fly -t tutorial unpause-pipeline -p my-pipe
```

1. fly sync is to sync the fly version with the target (might be already fine)
2. set-pipeline will install the pipeline.yml to concourse
3. newly set pipelines are paused automatically, so unpause it. (you can unpause it in the concourse dashboard as well)

Now everything should be set up. In your browser navigate to [Localhost](http://localhost:8080)
You should find your pipeline in the concourse dashboard. Try to trigger it.
