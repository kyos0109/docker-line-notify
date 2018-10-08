# drone-line-notify

```
docker run -e PLUGIN_TOKEN=${token} -e PLUGIN_MESSAGE="Test" --rm kyos0109/drone-line-notify
```
or in .drone.yml
```
notify:
    image: kyos0109/drone-line-notify:latest
    secrets: [ token_secret ]
    message: |
        Repo: {{.RepoName }}
        Status: {{.BuildStatus}}
        Branch: {{.RepoBranch}}
        Build NUM: {{.BuildNum}}
        Commit ID: {{.CommitID}}
        Author: {{.Author}}
        Commit Msg: {{.CommitMsg}}
    when:
      status: [ success, failure ]
```
