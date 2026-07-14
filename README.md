# pg-gh-3
Playground for Github Actions and Github Agentic Workflows

## Test workflow locally with act
 - Simple test command 
 - ref: https://nektosact.com/usage/index.html
```bash
act -j 'metadata'
```
 
 - In mac pc there is the following warning
```shell
WARN  ⚠ You are using Apple M-series chip and you have not specified container architecture, you might encounter issues while running act. If so, try running it with '--container-architecture linux/amd64'. ⚠
```
 - Use `.actrc` file to specify the architecture and other options
 - ref: https://nektosact.com/usage/index.html#configuration-file
```properties
--container-architecture=linux/amd64
```

## gh - cli commands
```shell
gh auth login
gh auth status
gh repo list
gh extension install github/gh-aw
gh auth status
gh aw add-wizard githubnext/agentics/daily-repo-status
gh aw run daily-repo-status --repo gerassimos/pg-gh-f2 --engine copilot
```
