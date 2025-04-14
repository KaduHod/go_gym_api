# Musculo Eskeletal API

## 📄 Visão Geral

A Musculo Eskeletal API é uma interface programática que fornece acesso estruturado a dados detalhados sobre o sistema musculoesquelético humano. Esta API é ideal para desenvolvedores que trabalham em aplicações relacionadas à anatomia, biomecânica, fitness, educação médica ou qualquer sistema que necessite de informações precisas sobre músculos, articulações e movimentos do corpo humano.

## 🌟 Recursos Principais

- **Catálogo Anatômico Completo**: Acesso a dados estruturados sobre músculos, articulações e movimentos
- **Mapeamento de Relações**: Conexões entre músculos, movimentos e articulações
- **Informações sobre Exercícios**: Detalhamento dos movimentos envolvidos em exercícios específicos
- **Sistema de Cache Eficiente**: Implementação de ETag para otimização de requisições
- **Autenticação Segura**: Via GitHub OAuth

## 🔗 Endpoints Disponíveis

| Endpoint | Descrição |
|----------|-----------|
| `/joints` | Lista todas as articulações do corpo humano |
| `/movements` | Retorna todos os movimentos possíveis |
| `/muscles` | Fornece hierarquia completa de grupos musculares e suas porções |
| `/muscles/movement-map` | Mapeia relações entre músculos, movimentos e articulações |
| `/muscles/portions` | Lista porções musculares isoladas |
| `/muscles/groups` | Retorna apenas os grupos musculares principais |
| `/exercises` | Catálogo de exercícios disponíveis |
| `/exercises/{id}` | Detalha movimentos específicos para um exercício |

## 🔧 Configuração Técnica

**URL Base**: `https://gymapi.kadu.tec.br/api/v1`
**Formato de Dados**: JSON (`application/json`)
**Versão Atual**: 1.0

### Códigos de Status

- `200 OK`: Requisição bem-sucedida
- `304 Not Modified`: ETag igual, conteúdo não modificado
- `401 Unauthorized`: Autenticação necessária ou inválida
- `404 Not Found`: Recurso não encontrado
- `500 Internal Server Error`: Erro no servidor

## 🔍 Filtragem de Dados

A API permite filtrar resultados usando query strings. Exemplo para o endpoint `/muscles/movement-map`:

- `muscle_group`: Filtrar por grupo muscular
- `muscle_portion`: Filtrar por porção muscular
- `joint`: Filtrar por articulação
- `movement`: Filtrar por movimento

## 📦 Resposta Padrão

Todas as respostas seguem este formato:

```json
{
  "status": "success",
  "data": [...],
  "metadata": {
    "total_itens": 123
  }
}
```

## 🔐 Autenticação

A API utiliza autenticação via token Bearer. Para obter seu token:

1. Faça login com sua conta GitHub no site da API
2. No Dashboard, copie seu token de acesso
3. ⚠️ Guarde seu token com segurança, ele é exibido apenas uma vez
4. Inclua seu token no cabeçalho `Authorization` das requisições

### Exemplo de Requisição Autenticada

```bash
curl --location 'https://gymapi.kadu.tec.br/api/v1/muscles/groups' \
--header 'If-None-Match: "aa450bb55d5a1318432d6b50817fa6fe"' \
--header 'Authorization: Bearer seu_token_aqui'
```

## 📚 Cache e Otimização

- Respostas dos endpoints GET são cacheáveis por até 1 semana
- Utilize o cabeçalho `If-None-Match` com o valor ETag para evitar transferências desnecessárias de dados
- O servidor responderá com status 304 quando o conteúdo não tiver sido modificado

## 👨‍💻 Começando

1. Registre-se e obtenha seu token de autenticação
2. Explore os endpoints disponíveis
3. Implemente os endpoints necessários na sua aplicação
4. Otimize suas requisições usando o sistema de cache

---

Para suporte ou mais informações, entre em contato comigo -> carlosjr.ribas@gmail.com.
