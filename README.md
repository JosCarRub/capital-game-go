# 🌍 Capital Game CLI 🎮

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/Docker-Required-blue)](https://www.docker.com/)

Test your geography knowledge with this addictive capital-guessing game, built with Go and running in Docker, right from your terminal!

<br>

<p align="center">
  <img src="/images/menu.png" alt="Gameplay Screenshot" width="700"/>
</p>

---

## ✨ Features

*   **Modern Terminal UI:** A beautiful and responsive interface built with the Go libraries `bubbletea` and `lipgloss`.
*   **Persistent Leaderboard:** Compete for the high score! Your name and points are saved to a MySQL database.
*   **Zero Dependencies:** The only requirement is Docker. No need to install Go or any other tools on your machine.
*   **Smart Input:** Don't worry about accents or capitalization. The game correctly recognizes `Bogota`, `bogotá`, and `BOGOTÁ` as the same answer.
*   **All-in-One Script:** A powerful `play.sh` script handles everything from building the image to running and cleaning up the environment.

---

## 🛠️ Tech Stack

*   **Language:** [Go](https://golang.org/)
*   **TUI Framework:** [Bubble Tea](https://github.com/charmbracelet/bubbletea) & [Lipgloss](https://github.com/charmbracelet/lipgloss)
*   **Database:** [MySQL 8.0](https://www.mysql.com/)
*   **Containerization:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

---

## 🚀 Getting Started

Follow these simple steps to get the game running on your local machine.

### Prerequisites

You must have the following software installed:
*   [Git](https://git-scm.com/downloads)
*   [Docker](https://docs.docker.com/get-docker/) & [Docker Compose](https://docs.docker.com/compose/install/)

### Installation & Launch

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/JosCarRub/capital-game-go.git
    cd capital-game-go
    ```

2.  **Grant execute permissions to the script:**
    This is a one-time step to make the helper script runnable.
    ```sh
    chmod +x play.sh
    ```

3.  **Launch the game!**
    Use the `up` command to build the Docker images, start the services, and run the game.
    ```sh
    ./play.sh up
    ```
    The first time you run this, it may take a few minutes to download and build the Docker images. Subsequent launches will be much faster.

---

## 📜 Script Commands

The `play.sh` script is your main tool for managing the application.

| Command        | Description                                                              |
|----------------|--------------------------------------------------------------------------|
| `./play.sh up`   | Builds the images (if needed), starts the database and app, and runs the game. |
| `./play.sh down` | Stops and removes all running containers, networks, and volumes.         |
| `./play.sh logs` | Tails the logs from all running services.                                |
| `./play.sh logs db` | Tails the logs specifically from the `db` service.                     |
| `./play.sh help` | Displays the help message with all available commands.                   |

---

## 📂 Project Architecture

The project follows a standard Go project layout to keep the code organized and maintainable.

```
.
├── cmd/                # Application entry points
│   └── capital-game/
│       └── main.go     # The main function that starts the app
├── data/               # Static data files
│   └── countries.json  # List of countries and capitals
├── internal/           # Private application code
│   ├── database/       # Database connection and query logic (Repository)
│   ├── game/           # Core game logic (rules, questions, scoring)
│   ├── style/          # Shared Lipgloss styles for the TUI
│   └── tui/            # Terminal User Interface code (Bubble Tea models and views)
├── Dockerfile          # Instructions to build the Go application's Docker image
├── docker-compose.yml  # Defines the `app` and `db` services
└── play.sh             # The main helper script to manage the project
```

---

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for more details.


---
