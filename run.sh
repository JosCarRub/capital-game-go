GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' 


cleanup() {
    echo -e "\n${YELLOW}🛑 Deteniendo todos los servicios...${NC}"
    docker-compose down
    echo -e "${GREEN}¡Limpieza completada! Hasta pronto.${NC}"
}

trap cleanup EXIT


echo -e "${BLUE}=====================================${NC}"
echo -e "${BLUE}🚀 Lanzador de Capital Game GO 🚀${NC}"
echo -e "${BLUE}=====================================${NC}"

# 1: Levantar la base de datos en detached
echo -e "\n${YELLOW}1/3 - Iniciando base de datos...${NC}"

docker-compose up -d db

# 2: base de datos lista?
echo -e "${YELLOW}2/3 - Esperando a que la base de datos esté lista...${NC}"

while [ "$(docker-compose ps -q db | xargs docker inspect -f '{{.State.Health.Status}}')" != "healthy" ]; do
    printf "."
    sleep 2
done
echo -e "\n${GREEN}¡Base de datos lista!${NC}"

# 3: play
echo -e "\n${YELLOW}3/3 - ¡Comenzamos! ${NC}"
echo -e "${BLUE}--------------------------------------------------${NC}\n"
docker-compose run --rm app /capital-game

exit 0