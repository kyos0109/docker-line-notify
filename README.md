# drone-line-notify

```
docker run -e PLUGIN_TOKEN=${token} -e PLUGIN_MESSAGE="Test" --rm kyos0109/drone-line-notify
```
# or in .drone.yml
```
  notify:
    image: kyos0109/drone-line-notify:latest
    secrets: [ token_secret ]
    message: |
        Status: {{.BuildStatus}}
        Repo: {{.RepoName }}
        Branch: {{.RepoBranch}}
        Build Num: {{.BuildNum}}
        Commit ID: {{.CommitID}}
        Author: {{.Author}}
        Commit Msg: {{.CommitMsg}}
        Link: {{.ResultLink}}
    when:
      status: [ success, failure ]
```
