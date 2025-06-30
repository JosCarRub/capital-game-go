# Capital Game GO 🌍

<div align="center">

🌍 **Interactive CLI Geography Game** 🌍

*Guess the capitals • Save your scores • Challenge yourself*

<div align="center">

*Built using Go & Docker*

<a href="https://go.dev/">
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
</a>
<a href="https://www.docker.com/">
<img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
</a>

</div>

---

</div>


### About The Project

**Capital Play GO** is an interactive command-line interface (CLI) game developed in Go. The main goal is to guess the capitals of countries from around the world. At the end of each round, players can save their scores to a local leaderboard, adding a competitive and replayable element to the game.

This project is the **Beta Version** of a game originally developed in Python. The primary objective of this migration was to practice and leverage the strengths of the Go programming language and its ability to compile into a single, portable binary.

The application is fully containerized using Docker and orchestrated with Docker Compose, ensuring a consistent and easy-to-set-up development and execution environment.

---

### Future Work

This Beta version serves as a solid foundation. Future development will focus on:

**Code Modularity** → Refactoring the code to be more modular and decoupled, following best practices and design patterns.

**Enhanced TUI** → Improving the terminal user interface with more advanced libraries (e.g., Bubble Tea) for a more aesthetic and dynamic experience.

**New Features** → Adding new game modes, difficulty levels, and potentially a global online leaderboard.

---

### Getting Started

Follow these instructions to get a local copy up and running.

#### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

#### Installation & Execution

**1. Clone the repository**
```sh
git clone https://github.com/your-username/capital-game-go.git
cd capital-game-go
```

**2. Build the Docker images and start the database service**

This command builds the Go application image and starts the MySQL database container in the background.
```sh
docker-compose up --build -d db
```

**3. Run the game!**

This command starts a new, interactive game session. Your keyboard will be connected directly to the game.
```sh
docker-compose run --rm app
```

**4. (Optional) Clean up the environment**

When you're done, you can stop the database container and remove the network.
```sh
docker-compose down
```

To completely reset the project and delete the leaderboard data, also remove the Docker volume:
```sh
docker volume rm capital-game-go_db_data
```

---

<div align="center">
<p align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" alt="Go" width="55" height="55"/>
  &nbsp;&nbsp;&nbsp;&nbsp;

  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" alt="Docker" width="55" height="55"/>
  &nbsp;&nbsp;&nbsp;&nbsp;

</p>

</div>