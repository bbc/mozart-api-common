# mozart-api-common

This repository contains packages that are used by both `mozart-config-api` and `mozart-template-api`.

## Development workflow

This section aims to set out the workflow for making changes to packages in the `mozart-api-common` repo and pulling those changes into `mozart-config-api` and `mozart-template-api`.

These instructions assume you have those three repositories checket out to the appropriate location in your `$GOPATH`.

1. Create new branches in all three repos to hold your work.
2. Rewrite the relevant `import` statements in the APIs to use the `mozart-api-common` in your $GOPATH, instead of in `Godeps/`, e.g. change `import "github.com/bbc/mozart-template-api/src/Godeps/_workspace/src/github.com/bbc/mozart-api-common/caching"` to `import "github.com/bbc/mozart-api-common/caching"`.
3. Make the required changes to the repos.
4. Commit your changes to `mozart-api-common`.
5. Pull these changes into the API repos by running `godep update github.com/bbc/mozart-api-common/...` and `godep save -r ./...` from their `src/` directories.
6. Commit the changes to the API repos.
7. Push your commits up to GitHub and create pull requests for the branches in each of the repos.
8. On :cake: approval merge your changes into the master branch.
