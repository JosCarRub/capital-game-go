set -e 

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
RED='\033[0;31m'
CYAN='\033[0;36m'
PURPLE='\033[0;35m'
NC='\033[0m'


print_header() {
    local message=$1
    echo -e "${PURPLE}╔══════════════════════════════════════════════════════════════╗${NC}"
    printf "${PURPLE}║%*s%*s║${NC}\n" $(((58 + ${#message}) / 2)) "$message" $(((58 - ${#message}) / 2)) ""
    echo -e "${PURPLE}╚══════════════════════════════════════════════════════════════╝${NC}"
}

print_step() {
    local step_num=$1
    local total_steps=$2
    local message=$3
    echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${YELLOW}📋 PASO ${step_num}/${total_steps}: ${message}${NC}"
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
}

show_progress() {
    local message=$1
    local duration=$2
    local steps=20
    local delay=$(bc <<< "scale=4; $duration / $steps")

    echo -ne "${CYAN}${message}${NC} ["
    for ((i=0; i<=steps; i++)); do
        printf "█"
        sleep "$delay"
    done
    echo -e "] ${GREEN}✓${NC}"
}


command_exists() {
    command -v "$1" >/dev/null 2>&1
}

check_docker() {
    echo -e "${CYAN}🔍 Verificando estado de Docker...${NC}"
    if ! docker info > /dev/null 2>&1; then
        echo -e "\n${RED}┌───────────────────────────────────────────────────────────┐${NC}"
        echo -e "${RED}│  ${YELLOW}ATENCIÓN:${RED} El demonio de Docker no parece estar en ejecución. │${NC}"
        echo -e "${RED}│  Por favor, inicia Docker Desktop y vuelve a intentarlo.    │${NC}"
        echo -e "${RED}└───────────────────────────────────────────────────────────┘${NC}"
        exit 1
    fi
    echo -e "${GREEN}✅ Docker está activo y listo.${NC}"
}

cleanup() {
    echo -e "\n\n${RED}🛑 Script interrumpido. Realizando limpieza...${NC}"
    docker-compose down --volumes > /dev/null 2>&1
    echo -e "${GREEN}✨ Entorno limpio. ¡Hasta pronto!${NC}"
}

usage() {
    echo -e "${YELLOW}Uso:${NC} ./play.sh ${CYAN}<comando>${NC}"
    echo
    echo -e "${YELLOW}Comandos disponibles:${NC}"
    echo -e "  ${GREEN}up${NC}      Levanta el entorno completo (BBDD y App) y ejecuta el juego."
    echo -e "          ${CYAN}Ejemplo:${NC} ./play.sh up"
    echo
    echo -e "  ${GREEN}down${NC}    Detiene y elimina todos los contenedores y volúmenes asociados."
    echo -e "          ${CYAN}Ejemplo:${NC} ./play.sh down"
    echo
    echo -e "  ${GREEN}logs${NC}     Muestra los logs de todos los servicios en tiempo real."
    echo -e "          ${CYAN}Ejemplo:${NC} ./play.sh logs db"
    echo
    echo -e "  ${GREEN}help${NC}    Muestra este mensaje de ayuda."
    echo
    exit 1
}



run_up() {
    trap cleanup INT

    if command_exists clear; then clear; fi
    print_header "🚀 LANZADOR DE CAPITAL GAME GO 🚀"
    echo -e "\n${YELLOW}⚡ Iniciando secuencia de arranque...${NC}"

    print_step 1 3 "Verificando entorno"
    check_docker

    print_step 2 3 "Construcción de la imagen"
    read -p "$(echo -e ${CYAN}"¿Deseas forzar la reconstrucción de la imagen de la app? (s/N): "${NC})" -n 1 -r
    echo
    if [[ $REPLY =~ ^[Ss]$ ]]; then
        show_progress "🛠️  Construyendo imagen de la aplicación..." 5
        docker-compose build app
        echo -e "${GREEN}✅ Imagen construida.${NC}"
    else
        echo -e "${CYAN}⏭️  Omitiendo reconstrucción.${NC}"
    fi

    print_step 3 3 "Iniciando servicios"
    show_progress "🐳 Levantando contenedor de base de datos..." 2
    docker-compose up -d db

    echo -e "${CYAN}🩺 Esperando a que la base de datos esté saludable...${NC}"
    local spinner_chars='-\|/'
    local counter=0
    while [ "$(docker-compose ps -q db | xargs docker inspect -f '{{if .State.Health}}{{.State.Health.Status}}{{end}}')" != "healthy" ]; do
        char_index=$((counter % 4))
        echo -ne "\r${CYAN}   Esperando respuesta... ${YELLOW}${spinner_chars:$char_index:1}${NC}"
        sleep 0.5
        ((counter++))
    done
    echo -e "\r${GREEN}   ¡Base de datos lista y funcionando! ✓${NC}\n"

    echo -e "\n${GREEN}┌─────────────────────────────────────────┐${NC}"
    echo -e "${GREEN}│  🎉 ¡Todo listo! Ejecutando juego...   │${NC}"
    echo -e "${GREEN}└─────────────────────────────────────────┘${NC}\n"
    

    docker-compose run --rm app ./capital-game
    
    echo -e "\n${BLUE}👋 ¡Gracias por jugar! Deteniendo servicios...${NC}"
    docker-compose down --volumes
}

main() {
    if ! command_exists docker || ! command_exists docker-compose; then
        echo -e "${RED}Error: Docker y/o docker-compose no están instalados. Por favor, instálalos para continuar.${NC}"
        exit 1
    fi

    if [ -z "$1" ]; then
        usage
    fi

    case "$1" in
        up) run_up ;;
        down) run_down ;;
        logs) shift; run_logs "$@" ;;
        help|-h|--help) usage ;;
        *) echo -e "${RED}Error: Comando desconocido '$1'${NC}\n"; usage ;;
    esac
}

main "$@"