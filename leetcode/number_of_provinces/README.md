# Number of Provinces

[Leetcode link](leetcode.com/problems/number-of-provinces)

Esse desafio poderia ter sido resolvido utilizando três técnicas diferentes:
- Deep First Search (DFS)
- Breadth First Search (BFS)
- Union Find

Decidi seguir pela implementação mais simples: __DFS__, só precisamos percorrer todos os caminhos possíveis.
Usando DFS isso se torna algo trivial de ser feito, a ideia é que para cada nós que nós visitamos marcamos ele como visitado,
e caso as conexões desse nó em especifico ainda não foram visitadas o algoritmo vai visitar.
