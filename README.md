# Árvore Genealógica

## Endpoints

### Health
- {GET} localhost:9000/family-tree/v1/health
### Relationship Type
- {POST} localhost:9000/family-tree/v1/relationship_types<br>
    Segue o payload para inserção:<br>
    {<br>
        "type": "cousin"<br>
    }<br>

- {GET} localhost:9000/family-tree/v1/relationships/types<br>
    Pega por queryParam: [?type=cousin]<br>

- {GET} localhost:9000/family-tree/v1/relationships/types/list<br>
    Traz uma lista dos tipos de relações que existem no banco<br>

- {PUT} localhost:9000/family-tree/v1/relationships/types<br>
    Segue o payload para atualização:<br>
    {<br>
        "id": "insira o id do type de relacionamento que deseja atualizar",<br>
        "type": "cousins"<br>
    }<br>
### Person
- {POST} localhost:9000/family-tree/v1/people<br>
    Segue o payload para inserção:<br>
    {<br>
        "name": "Nathan Ferreira"<br>
    }<br>

- {GET} localhost:9000/family-tree/v1/people<br>
    Pega por queryParam: ["?id=id_da_person"] ou ["?name=name_da_person"]<br>

- {GET} localhost:9000/family-tree/v1/people/list<br>
    Traz uma lista de pessoas<br>

- {PUT} localhost:9000/family-tree/v1/people<br>
    Segue o payload para atualização:<br>
    {<br>
        "id": "insira o id da pessoa que deseja atualizar",<br>
        "name": "Nathan Ferreira da Silva"<br>
    }<br>

### Relationship

## Architecture

### Frameworks utilizados
- Netflix/env   - para variáves de ambiente
- Echo v4       - para web service
- Gorm          - para acesso e comunicação com banco
- Gorm Postgres - para comunicação entre gorm e postgres
- Google UUID   - geração de chaves primárias
