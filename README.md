Atualizar código comentado (parte desatualizada)
1 - codigo não funciona pelo console
2 - encontrar forma de não usar []byte
3 - erro ao usar post em um json que não se encaixa no formato Command, não está sendo tratado.


Possível solução 3: ai ínves de usar NewDecoder().Decode() usar Unmarshal
https://pkg.go.dev/encoding/json#Unmarshal
