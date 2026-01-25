# Contributing to Task CLI

Thanks for wanting to help!  
Even small contributions (typos, docs, bug reports, tests) are very welcome.

website for showcasing the cli tool : https://task-ten-gules.vercel.app

## Quick Ways to Contribute

- Report bugs  
- Suggest small features in the CLI and also in the [website](https://task-ten-gules.vercel.app).
- Fix typos in help text or README  
- Improve Windows/macOS instructions  
- Add/update tests  
- Small refactoring or cleanups  

## Development Setup

```bash
git clone https://github.com/anikchand461/task.git
cd task

go mod tidy
go run ./cmd/task --help
```

Build your own binary:

```bash
go build -o task ./cmd/task
./task list
```

## Workflow

1. Pick or create an issue  
2. Create a branch:  
   `git checkout -b fix/something` or `git checkout -b docs/improve-readme`
3. Make changes  
4. Test locally:  
   `go run ./cmd/task ...` or `./task ...`
5. Commit (small & clear messages)  
6. Push & open a Pull Request

## Pull Request Guidelines

- Title should be short and clear  
- Describe what + why in the PR body  
- Reference related issue if any (`Fixes #42`)  
- One feature / fix per PR is preferred


Feel free to comment "I'd like to work on this" — I'll assign it.

Questions?  
Just open an issue or mention @anikchand461

Thanks for helping make Task CLI better! 🚀
