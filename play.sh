
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
RED='\033[0;31m'
CYAN='\033[0;36m'
PURPLE='\033[0;35m'
NC='\033[0m'

show_progress() {
    local steps=$1
    local message=$2
    local width=40
    
    echo -ne "${CYAN}${message}${NC} ["
    
    for ((i=0; i<=steps; i++)); do
        local filled=$((i * width / steps))
        local empty=$((width - filled))
        
        printf "\r${CYAN}${message}${NC} ["
        
        for ((j=0; j<filled; j++)); do
            printf "█"
        done
        
        for ((j=0; j<empty; j++)); do
            printf "░"
        done
        
        local percentage=$((i * 100 / steps))
        printf "] ${YELLOW}%d%%${NC}" $percentage
        
        sleep 0.1
    done
    echo
}

show_spinner() {
    local message=$1
    local duration=${2:-3}
    local spinner_chars='-\|/'
    local counter=0
    
    echo -ne "${CYAN}${message}${NC} "
    
    local iterations=$((duration * 10))
    
    for ((i=0; i<iterations; i++)); do
        local char_index=$((counter % 4))
        echo -ne "\b${YELLOW}${spinner_chars:$char_index:1}${NC}"
        sleep 0.1
        ((counter++))
    done
    
    echo -e "\b${GREEN}✓${NC}"
}

show_dots() {
    local message=$1
    local max_dots=${2:-3}
    
    echo -ne "${CYAN}${message}${NC}"
    
    for ((i=1; i<=max_dots; i++)); do
        echo -ne "."
        sleep 0.5
    done
    
    echo -e " ${GREEN}✓${NC}"
}

command_exists() {
    command -v "$1" >/dev/null 2>&1
}

clear_line() {
    echo -ne "\r\033[K"
}

cleanup() {
    echo -e "\n${RED}┌─────────────────────────────────────────┐${NC}"
    echo -e "${RED}│  🛑 Deteniendo todos los servicios...  │${NC}"
    echo -e "${RED}└─────────────────────────────────────────┘${NC}"
    
    show_progress 20 "Limpiando contenedores"
    docker-compose down > /dev/null 2>&1
    
    echo -e "\n${GREEN}┌─────────────────────────────────────────┐${NC}"
    echo -e "${GREEN}│     ✨ ¡Limpieza completada!           │${NC}"
    echo -e "${GREEN}│        Hasta pronto. 👋                │${NC}"
    echo -e "${GREEN}└─────────────────────────────────────────┘${NC}"
}

trap cleanup EXIT

if ! command_exists docker-compose; then
    echo -e "${RED}❌ Error: docker-compose no está instalado${NC}"
    exit 1
fi


if command_exists clear; then
    clear
fi


echo -e "${PURPLE}╔══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${PURPLE}║                                                              ║${NC}"
echo -e "${PURPLE}║  ${CYAN}🚀 LANZADOR DE CAPITAL GAME GO 🚀${PURPLE}                      ║${NC}"
echo -e "${PURPLE}║                                                              ║${NC}"
echo -e "${PURPLE}║  ${BLUE}Sistema de inicio automatizado v2.0${PURPLE}                   ║${NC}"
echo -e "${PURPLE}║                                                              ║${NC}"
echo -e "${PURPLE}╚══════════════════════════════════════════════════════════════╝${NC}"

echo -e "\n${YELLOW}⚡ Iniciando secuencia de arranque...${NC}\n"

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${YELLOW}📋 PASO 1/3: Iniciando base de datos${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

show_progress 30 "🐳 Levantando contenedor de base de datos"

docker-compose up -d db > /dev/null 2>&1

show_spinner "Configurando servicios" 2

echo -e "${GREEN}✅ Base de datos iniciada correctamente${NC}\n"

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${YELLOW}📋 PASO 2/3: Verificando estado de la base de datos${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "${CYAN}🔍 Verificando health check de la base de datos...${NC}"

spinner_chars='-\|/'
counter=0

while [ "$(docker-compose ps -q db | xargs docker inspect -f '{{.State.Health.Status}}')" != "healthy" ]; do
    char_index=$((counter % 4))
    echo -ne "\r${CYAN}🔄 Esperando respuesta de la base de datos ${YELLOW}${spinner_chars:$char_index:1}${NC}"
    sleep 0.5
    ((counter++))
done

clear_line
echo -e "${GREEN}✅ ¡Base de datos lista y funcionando!${NC}\n"

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${YELLOW}📋 PASO 3/3: Iniciando Capital Game${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

show_progress 25 "🎮 Preparando entorno de juego"

echo -e "\n${GREEN}┌─────────────────────────────────────────┐${NC}"
echo -e "${GREEN}│  🎉 ¡Todo listo! Ejecutando juego...   │${NC}"
echo -e "${GREEN}└─────────────────────────────────────────┘${NC}\n"

echo -e "${PURPLE}╔══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${PURPLE}║                    🎮 CAPITAL GAME GO 🎮                    ║${NC}"
echo -e "${PURPLE}╚══════════════════════════════════════════════════════════════╝${NC}\n"

docker-compose run --rm app /capital-game

exit 0