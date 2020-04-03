docker-compose -f part22.yaml down
docker-compose -f part1.yaml down
docker-compose -f addshared.yml down
rm -rf ./node*
rm -rf ./gateway