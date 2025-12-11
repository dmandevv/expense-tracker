# Expense Tracker 📝

A simple expense tracker application to manage your finances. This project is an exercise from the roadmap.sh project: https://roadmap.sh/projects/expense-tracker 📚

**What it does**
- **Launches** a REPL for the user to add, delete, and view their expenses
- **Saves** data to a .txt file store locally

**Prerequisites**

- Go 1.25.1 or newer installed

## Installation 📦

There are a few easy ways to get and run this project.

- **Install using Go**

```shell
go install github.com/dmandevv/expense-tracker@latest
```
This installs the binary into your `$GOBIN` (or `$GOPATH/bin`) so you can run it globally.

- **Clone and build from source**

```shell
git clone https://github.com/dmandevv/expense-tracker.git
cd expense-tracker
go build -o expense-tracker .
```

After building you can run the generated `expense-tracker` binary.

- **Run directly (quick, no install)**

```shell
go run .
```

## Setup ⚙️

Expense tracker saves data to a YAML file. By default the app uses `./expenses.txt` in the current working directory. You can override the path using the `SAVE_DATA_PATH` environment variable.

Create a `.env` file in the project root (or set the variable in your shell) with the following content:

```env
SAVE_DATA_PATH="DIRECTORY_PATH/FILE_NAME.txt"
```

- The program will create the file automatically when you save tasks.

## Usage ✅

Enter "help" for a complete list of commands.
Enter "exit" to close the program.

- Add a new expense

```shell
add --description "Groceries" --amount 46.70
```
- List all expenses

```shell
list
```

- List total amount of expenses for the current year

```shell
summary
```

- List total amount of expenses for April (4th month) of the current year

```shell
summary --month 4
```

- Change the info for expense with ID 1 ("list" command can help you find IDs)

```shell
update --id 1 --description "Groceries" --amount 23.79
```
- Delete the expense with ID of 7

```shell
delete --id 7
```
