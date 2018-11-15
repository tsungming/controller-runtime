# Release Process

The Kubernetes controller-runtime Project is released on an as-needed basis. The process is as follows:

- An issue is proposing a new release with a changelog since the last release
- 2 [OWNERS](OWNERS) must LGTM this release
- An OWNER runs `git tag -s $VERSION` and inserts the changelog and pushes the tag with `git push $VERSION`
- The release issue is closed
- Auto Deleting Merged Branch