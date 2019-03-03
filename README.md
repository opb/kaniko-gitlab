Extend Kaniko to support gitlab login through env vars.

Just playing around for the moment.

to get to run in Gitlab CI under a fairly standard docker setup (not kubernetes) you need something like this in your `.gitlab-ci.yml` file:

```
testing kaniko:
  stage: build
  image: opb2k/kaniko-gitlab
  script:
    - echo '{"credHelpers":{"my.gitlab.server:4567":"gitlab-login"}}' > /kaniko/.docker/config.json
    - '/kaniko/executor --destination ${CI_REGISTRY_IMAGE}:latest --context $CI_PROJECT_DIR'
```

Few issues it would be nice to get tidied up:

1. need to use debug image in order to have `sh` available. This is needed because gitlab runners using the docker executor will normally set a shell. Need to try a runner config which won't try to use a shell not sure if this is possible.
2. have to set `--context` because can't configure gitlab to use simply `/workspace`. Could potentially overwrite the way the command is called in our Dockerfile, removing need for this step.
3. need something better for setting the docker `config.json`. Very ugly right now. 

Could probably use a shell script to do some of this (given we have busybox installed anyway) which would leave the interface to use in the `.gitlab-ci.yml` file pretty lean.
 