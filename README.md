# godo 📝

A simple CLI todo manager written in Go. Add, list, and remove tasks directly from your terminal — no fuss, no clutter.

## Installation

Make sure you have Go installed (≥1.18).

```bash
go install github.com/rishabhdubeyy/godo/cmd/godo-cli@latest
```

This will put the binary in your `$GOPATH/bin` (usually `$HOME/go/bin`).
Add it to your `PATH` if it isn’t already:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

## Usage

```bash
# Add a task
godo-cli add "Buy groceries"

# List all tasks
godo-cli list

# Delete a task by index
godo-cli del 2
```

Example output:

```
1. Buy groceries
2. Finish project report
3. Call mom
```

After deleting index `2`:

```
1. Buy groceries
2. Call mom
```

## Features

* ✅ Add tasks quickly from your terminal
* 📋 List tasks in order with auto-incrementing IDs
* 🗑️ Delete tasks by index (IDs re-adjust automatically)
* 💾 Persists tasks locally

## Roadmap

* [ ] Mark tasks as done
* [ ] Support due dates
* [ ] Pretty formatting / colors
* [ ] Sync with cloud storage

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you’d like to add.
