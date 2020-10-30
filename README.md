### Resultado

    +------------+------------+------------+
    |   COUNT    |    JAVA    |   PYTHON   |
    +------------+------------+------------+
    |   STARS    |   972106   |   935077   |
    |  WATCHERS  |   52853    |   57491    |
    |   FORKS    |   275665   |   369103   |
    |  RELEASES  |    831     |    1211    |
    |    AGE     |    610     |    541     |
    |   LINES    |  36999614  |  13792083  |
    |   CODE     |  26694939  |  12637649  |
    |  COMMENTS  |  7719849   |   441557   |
    |   BLANKS   |  2584826   |   712877   |
    | COMPLEXITY |  1933779   |   250840   |
    +------------+------------+------------+

    +------------+------------+------------+
    |    MEAN    |    JAVA    |   PYTHON   |
    +------------+------------+------------+
    |   STARS    |    9721    |    9350    |
    |  WATCHERS  |    528     |    574     |
    |   FORKS    |    2756    |    3691    |
    |  RELEASES  |     8      |     12     |
    |    AGE     |     6      |     5      |
    |   LINES    |   369996   |   137920   |
    |   CODE     |   266949   |   126376   |
    |  COMMENTS  |   77198    |    4415    |
    |   BLANKS   |   25848    |    7128    |
    | COMPLEXITY |   19337    |    2508    |
    +------------+------------+------------+

## Algorithm

1. Get Repository.URL from CSV
2. Clone the Repository
3. 'cd' into the folder
4. Run command 'scc -f {extension format} -o {output file}'
5. Delete the folder created
6. Repeat with next repository

### Questões

##### RQ 01

- Quais as características dos top-100 repositórios Java mais populares?

---

##### RQ 02

- Quais as características dos top-100 repositórios Python mais populares?

---

##### RQ 03

- Repositórios Java e Python populares possuem características de “boa qualidade” semelhantes?

---

##### RQ 04

- A popularidade influencia nas características dos repositórios Java e Python?

---

### Métricas

Utilizaremos como fatores de qualidade métricas associadas a quatro dimensões:

##### Popularidade

- número de estrelas
- número de watchers
- número de forks dos repositórios coletados

##### Tamanho

X - linhas de código (LOC e SLOC)
X - linhas de comentários.

##### Atividade

- número de releases
- frequência de publicação de releases (número de releases / dias)

##### Maturidade

- idade (em anos) de cada repositório coletado.

---
