# lab-sast

Laboratório de estudo de **SAST** (análise estática de segurança). O código daqui tem **falhas de propósito** e não deve ser usado em produção. Não contém código nem dados de empresa.

## O que roda

| Workflow | Ferramenta | O que analisa |
| --- | --- | --- |
| `semgrep.yml` | Semgrep (CLI aberto) | Go e Python |
| `codeql.yml` | CodeQL (GitHub) | Go e Python |
| `checkov.yml` | Checkov | Terraform (IaC) |

Os três enviam o resultado em SARIF para a aba **Security → Code scanning**, onde dá para comparar lado a lado.

## Gabarito: falhas plantadas

Use esta lista para medir o que cada ferramenta achou (verdadeiro positivo), o que deixou passar (falso negativo) e o que apontou sem ser problema (falso positivo).

### Go (`app-go/main.go`)

| # | Falha | Onde |
| --- | --- | --- |
| G1 | Injeção de SQL | `buscarUsuario` |
| G2 | Injeção de comando | `ping` |
| G3 | XSS refletido | `saudacao` |
| G4 | Hash fraco (MD5) para senha | `hashSenha` |
| G5 | Senha fixa no código | `dbPassword` |
| G-ok | Consulta parametrizada (**não** é falha) | `buscarUsuarioSeguro` |

### Python (`app-python/app.py`)

| # | Falha | Onde |
| --- | --- | --- |
| P1 | Injeção de SQL | `usuario` |
| P2 | Injeção de comando (`shell=True`) | `ping` |
| P3 | Execução de código (`eval`) | `calc` |
| P4 | Desserialização insegura (`pickle`) | `carregar` |
| P5 | Hash fraco (MD5) para senha | `hash_senha` |
| P6 | Chave secreta fixa no código | `app.secret_key` |
| P7 | Flask com `debug=True` exposto | `__main__` |
| P-ok | Consulta parametrizada (**não** é falha) | `usuario_seguro` |

### Terraform (`infra/main.tf`)

| # | Falha |
| --- | --- |
| T1 | Bucket S3 sem bloqueio de acesso público |
| T2 | Security group com SSH aberto para 0.0.0.0/0 |
| T3 | Banco RDS sem criptografia |
| T4 | Banco RDS acessível publicamente |
| T5 | Senha do banco fixa no código |

## Como registrar a comparação

Para cada ferramenta, anote: falhas do gabarito encontradas, falsos positivos, tempo do job no Actions e se o alerta mostra o caminho do dado (fonte → ponto vulnerável).
