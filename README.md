# Árvore Genealógica

## Endpoints

### Health
- {GET} localhost:9000/family-tree/v1/health
### Relationship Type
- {POST} localhost:9000/family-tree/v1/relationship_types
    Segue o payload para inserção:
    {
        "type": "cousin"
    }

- {GET} localhost:9000/family-tree/v1/relationships/types
    Pega por queryParam: [?type=cousin]

- {GET} localhost:9000/family-tree/v1/relationships/types/list
    Traz uma lista dos tipos de relações que existem no banco

- {PUT} localhost:9000/family-tree/v1/relationships/types
    Segue o payload para atualização:
    {
        "id": "insira o id do type de relacionamento que deseja atualizar",
        "type": "cousins"
    }
### Person
- {POST} localhost:9000/family-tree/v1/people
    Segue o payload para inserção:
    {
        "name": "Nathan Ferreira"
    }

- {GET} localhost:9000/family-tree/v1/people
    Pega por queryParam: ["?id=id_da_person"] ou ["?name=name_da_person"]

- {GET} localhost:9000/family-tree/v1/people/list
    Traz uma lista de pessoas

- {PUT} localhost:9000/family-tree/v1/people
    Segue o payload para atualização:
    {
        "id": "insira o id da pessoa que deseja atualizar",
        "name": "Nathan Ferreira da Silva"
    }

### Relationship

## Architecture

### Frameworks utilizados
- Netflix/env   - para variáves de ambiente
- Echo v4       - para web service
- Gorm          - para acesso e comunicação com banco
- Gorm Postgres - para comunicação entre gorm e postgres
- Google UUID   - geração de chaves primárias