# Todo RPG

Todo RPG is a lightweight web application written in Go that combines task management ("To-Do") with role-playing game (RPG) elements. It provides a fun and engaging way to track daily tasks while leveling up your character based on completed objectives.

## Features

- Simple, intuitive web interface for managing tasks with priorities
- SQLite database backend for persistence
- RPG-style progression system incentivizing productivity (coming soon)
- Easily extendable and configurable

## Installation

1. Ensure you have [Go](https://go.dev/dl/) installed (version 1.25.3 recommended).

2. Clone this repository:

```bash
git clone https://github.com/link552/todo-rpg.git
cd todo-rpg
```

3. Set the port for HTTP (option, defaults to `8080`):

```bash
export HTTP_PORT=8080
```

4. Set the path to the SQLite database (optional, defaults to `./todorpg.db`):

```bash
export SQLITE_DB_PATH="./todorpg.db"
```

5. Run the application:

```bash
go run cmd/todorpg/todorpg.go
```

6. Open a browser and navigate to [http://localhost:8080](http://localhost:8080) (or your specified HTTP port) to access the app.

## Project Structure

- `cmd` - Application entry points
- `internal` - Internal source code
- `web` - Static web files and templates
