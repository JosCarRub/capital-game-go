
---

# Capital Game GO

## 🇬🇧 English

### About The Project

**Capital Play GO** is an interactive command-line interface (CLI) game developed in Go. The main goal is to guess the capitals of countries from around the world. At the end of each round, players can save their scores to a local leaderboard, adding a competitive and replayable element to the game.

This project is the **Beta Version** of a game originally developed in Python. The primary objective of this migration was to practice and leverage the strengths of the Go programming language and its ability to compile into a single, portable binary.

The application is fully containerized using Docker and orchestrated with Docker Compose, ensuring a consistent and easy-to-set-up development and execution environment.

### Future Work

This Beta version serves as a solid foundation. Future development will focus on:
*   **Code Modularity:** Refactoring the code to be more modular and decoupled, following best practices and design patterns.
*   **Enhanced TUI:** Improving the terminal user interface with more advanced libraries (e.g., Bubble Tea) for a more aesthetic and dynamic experience.
*   **New Features:** Adding new game modes, difficulty levels, and potentially a global online leaderboard.

### Getting Started

Follow these instructions to get a local copy up and running.

#### Prerequisites

*   [Docker](https://www.docker.com/get-started)
*   [Docker Compose](https://docs.docker.com/compose/install/)

#### Installation & Execution

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/your-username/capital-game-go.git
    cd capital-game-go
    ```

2.  **Build the Docker images and start the database service:**
    This command builds the Go application image and starts the MySQL database container in the background.
    ```sh
    docker-compose up --build -d db
    ```

3.  **Run the game!**
    This command starts a new, interactive game session. Your keyboard will be connected directly to the game.
    ```sh
    docker-compose run --rm app
    ```

4.  **(Optional) Clean up the environment:**
    When you're done, you can stop the database container and remove the network.
    ```sh
    docker-compose down
    ```
    To completely reset the project and delete the leaderboard data, also remove the Docker volume:
    ```sh
    docker volume rm capital-game-go_db_data
    ```

---

## 🇪🇸 Español

### Sobre El Proyecto

**Capital Play GO** es un juego interactivo de línea de comandos (CLI) desarrollado en Go. El objetivo principal es adivinar las capitales de países de todo el mundo. Al final de cada ronda, los jugadores pueden guardar sus puntuaciones en una tabla de clasificación local, añadiendo un elemento competitivo y de rejugabilidad.

Este proyecto es la **Versión Beta** de un juego originalmente desarrollado en Python. El objetivo principal de esta migración ha sido practicar y aprovechar las fortalezas del lenguaje de programación Go y su habilidad para compilar en un único binario portable.

La aplicación está completamente contenerizada usando Docker y orquestada con Docker Compose, lo que garantiza un entorno de desarrollo y ejecución consistente y fácil de configurar.

### Trabajo Futuro

Esta Versión Beta sirve como una base sólida. El desarrollo futuro se centrará en:
*   **Modularidad del Código:** Refactorizar el código para hacerlo más modular y desacoplado, siguiendo las mejores prácticas y patrones de diseño.
*   **Mejora de la TUI:** Mejorar la interfaz de usuario de la terminal con librerías más avanzadas (ej: Bubble Tea) para una experiencia más estética y dinámica.
*   **Nuevas Funcionalidades:** Añadir nuevos modos de juego, niveles de dificultad y, potencialmente, una tabla de clasificación global en línea.

### Cómo Empezar

Sigue estas instrucciones para tener una copia local del proyecto funcionando.

#### Prerrequisitos

*   [Docker](https://www.docker.com/get-started)
*   [Docker Compose](https://docs.docker.com/compose/install/)

#### Instalación y Ejecución

1.  **Clona el repositorio:**
    ```sh
    git clone https://github.com/tu-usuario/capital-game-go.git
    cd capital-game-go
    ```

2.  **Construye las imágenes de Docker e inicia el servicio de base de datos:**
    Este comando construye la imagen de la aplicación Go e inicia el contenedor de la base de datos MySQL en segundo plano.
    ```sh
    docker-compose up --build -d db
    ```

3.  **¡Ejecuta el juego!**
    Este comando inicia una nueva sesión de juego interactiva. Tu teclado estará conectado directamente al juego.
    ```sh
    docker-compose run --rm app
    ```

4.  **(Opcional) Limpia el entorno:**
    Cuando termines, puedes detener el contenedor de la base de datos y eliminar la red.
    ```sh
    docker-compose down
    ```
    Para reiniciar completamente el proyecto y borrar los datos de la clasificación, elimina también el volumen de Docker:
    ```sh
    docker volume rm capital-game-go_db_data
    ```

---