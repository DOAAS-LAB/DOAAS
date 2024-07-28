# DOAAS

CLI para instalação e atualização de ferramentas utilizadas por DevOps e SREs.


## Cobra

Utilizando cobra-cli https://cobra.dev/ e Go

## Ideia de estrutura

```
Escolha uma opção
  1 - Instalar ferramentas
    1 - AWS CLI
    2 - Terraform
    3 - Docker 
    4 - Etc..
  2 - Verificar versões instaladas:
    - Lista atualizados
    Ferramentas x,y , z estão desatualziadas, gostaria de atualizar?
  3 - Update
    - verifica versões instaladas e compara com as tags do github de cada ferramenta
    - lista as desatualizada
    - ao clicar em em cada ferramenta ela deve atualizar e ao finalizar retornar a lista de ferramentas
    Magic happens
```  

  
