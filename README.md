# mozart-api-common

This repository contains packages that are used by both `mozart-config-api` and `mozart-template-api`.

They can be imported using `import` statements:

```go
import (
	"github.com/bbc/mozart-api-common/caching"
)
```

Then, as with other dependencies, they are vendored into the local workspace by [godep](https://github.com/tools/godep) and the import statement is rewritten:

```go
import (
  "github.com/bbc/mozart-config-api/src/Godeps/_workspace/src/github.com/bbc/mozart-api-common/caching"
)
```

## Development workflow

This section describes the workflow for making changes to packages in the `mozart-api-common` repo and pulling those changes into `mozart-config-api` and `mozart-template-api`.

1. Make sure you have `mozart-api-common`, `mozart-config-api` and `mozart-template` checked out in the appropriate location in your `$GOPATH`.
2. Create new branches in all three repos to hold your work.
3. Rewrite the relevant `import` statements in the APIs to use the `mozart-api-common` in your `$GOPATH`, instead of in `Godeps/`, e.g. change `import "github.com/bbc/mozart-template-api/src/Godeps/_workspace/src/github.com/bbc/mozart-api-common/caching"` to `import "github.com/bbc/mozart-api-common/caching"`.
4. Make the required changes to the repos.
5. Commit your changes to `mozart-api-common`.
6. Pull these changes into the API repos by running `godep update github.com/bbc/mozart-api-common/...` and `godep save -r ./...` from their `src/` directories.
7. Commit the changes to the API repos.
8. Push your commits up to GitHub and create pull requests for the branches in each of the repos.
9. On :cake: approval merge your changes into the master branch.
