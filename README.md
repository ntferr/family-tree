# Árvore Genealógica

## Endpoints

### Health
- {GET} localhost:9000/family-tree/v1/health
### Relationship Type
- {POST} localhost:9000/family-tree/v1/relationship_types
    segue o payload para inserção 
    {
        "type": "insira o type de relação que quer salvar"
    }

- {GET} localhost:9000/family-tree/v1/relationship_types
    Pega por parâmetro inserada na ulr exemplo [?type=cousin]

- {GET} localhost:9000/family-tree/v1/relationship_types/list
    Traz uma lista dos tipos de relações que existem no banco
### Person
- 
### Relationship

## Architecture

### Frameworks utilizados
- Netflix/env   - para variáves de ambiente
- Echo v4       - para web service
- Gorm          - para acesso e comunicação com banco
- Gorm Postgres - para comunicação entre gorm e postgres
- Google UUID   - geração de chaves primárias