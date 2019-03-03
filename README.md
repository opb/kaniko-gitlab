Extend Kaniko to support gitlab login through env vars.

Just playing around for the moment.

to get to run in Gitlab CI under a fairly standard docker setup (not kubernetes) you need something like this in your `.gitlab-ci.yml` file:

```
testing kaniko:
  stage: build
  image:
    name: opb2k/kaniko-gitlab
    entrypoint: [""]
  script:
    - echo '{"credHelpers":{"my.gitlab.server:4567":"gitlab-login"}}' > /kaniko/.docker/config.json
    - '/kaniko/executor --destination ${CI_REGISTRY_IMAGE}:latest --context $CI_PROJECT_DIR'
```

Few issues it would be nice to get tidied up:

1. need to set entrypoint
 