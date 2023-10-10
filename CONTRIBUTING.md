# Contributing to appkit

- [Getting Started](#getting-started)
- [Making Changes](#making-changes)
  - [Picking an Issue](#picking-an-issue)
  - [Creating a Branch for Changes](#creating-a-branch-for-changes)
  - [Submitting Changes](#submitting-changes)
- [Submission Standards](#submission-standards)
- [Versioning Policy](#versioning-policy)
  - [Deprecated Features](#deprecated-features)

We welcome community contributions to appkit!

## Getting Started

1. Ensure you have [installed Golang](https://go.dev/dl/) on your development machine. appkit is currently on version `1.20.2`.
2. `git clone git@github.com:decentplatforms/matcha.git`

## Making Changes

### Picking an Issue

We do our best to keep issues up to date and appropriately tagged. Here's a quick rundown on how to pick an issue based on tags.

1. If you're a first time contributor, find one tagged `good first issue` or `patch`. These tend to be short, accessible tasks to help you get to know a specific part of the codebase.
2. Once you're more comfortable with contributing to appkit, pick up `patch`, `minor`, and `bugfix` issues that interest you.
3. If you're feeling a non-code task, there's usually a `documentation` issue or two avaliable for assignment.

Once you have decided on an issue, just leave a comment asking to have it assigned to you. Issues are generally first-come first-serve. If an issue has been inactive for 3 weeks, we will consider reassignment.

### Creating a Branch for Changes

1. Check out the correct development branch. You can find these below in [Versioning Policy](#versioning-policy). `git checkout [version-branch]`
2. Pull in remote changes by doing a fetch/rebase: `git pull --rebase`
3. Create your own development branch: `git branch [my-branch]`
4. Check out your branch: `git checkout [my-branch]`
5. Make your changes!

Please don't make changes on branch `main` or any semver (`*/vX.X.X`) branch. It will only make your life harder. If you accidentally commit to one of these, this guide may help (external link): <https://dev.to/projectpage/how-to-move-a-commit-to-another-branch-in-git-4lj4>

### Submitting Changes

1. Push your changes to your personal fork: `git push origin [my-branch]`
2. Make a pull request to the development branch, and request review from a maintainer.
3. Revise your changes based on maintainer feedback and test results.

Once your changes are approved, they'll be squash-and-merged into the feature branch.

## Submission Standards

appkit contains many small projects in varying lifecycle states, and depending on that state, may be subject to different standards.

### All Projects

- **Style**: Run `gofmt` on your code.
- **Copyright/License**: All contributions must be under the LICENSE, and must only be copyrighted by `appkit Authors`. This doesn't affect your copyright status--if you've contributed to appkit, you are an appkit author, and your contribution/copyright is tracked through version control.
- **No AI**: Please do not use AI tools like ChatGPT and GitHub Copilot in your contributions to appkit.

### Unversioned

Unversioned projects are the wild west. All contributions meeting the above are welcome.

### v1.0.0 and above

1. **Performance**: Changes must not significantly decrease performance unless they are urgent bugfixes. End-to-end benchmarks are provided in the docs folder; additionally, if your feature or change is to a system that is integrated heavily, we suggest you add and check benchmarks for it.
2. **Testing**: Test coverage should stay above 95%. New behavior is expected to have associated unit tests, and PRs that drop coverage by more than 2%, or below 90%, will be automatically rejected.
3. **Documentation**: Decent Platforms strongly encourages in-code documentation. Maintainers may request that you add additional documentation to your code.
4. **Style**: Run `gofmt` on your code.

## Versioning Policy

Projects in appkit follow semantic versioning, formatted as `vX.Y.Z`, where X is major version, Y is minor version, and Z is patch number.

- Major: The change is non-essential and breaks the existing API. Major changes are not currently being accepted.
- Minor: The change doesn't break the existing API, but changes large portions of the internals of the library, or adds significant functionality to the API. Branch off of, and make pull requests to, a version branch like `logf/v1.0.0`.
- Patch: The change is a bugfix, minor performance improvement, minor API change, or auxilary component (like middleware). Branch off of, and make pull requests to, branch `main`.

### Deprecated Features

To maintain the long-term health of the project, you/we may elect to *deprecate* features, meaning that we won't support them going forward. If you have a change that overrides old functionality, it's preferred that you mark the old as deprecated and implement alongside it, rather than delete it from the project entirely; we want to keep behavior the same between versions when it's not a bug problem.
