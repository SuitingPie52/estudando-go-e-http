Atualizar código comentado (parte desatualizada)
1 - codigo não funciona pelo console
2 - encontrar forma de não usar []byte
3 - erro ao usar post em um json que não se encaixa no formato Command, não está sendo tratado.

Possível solução 3: ai ínves de usar NewDecoder().Decode() usar Unmarshal
https://pkg.go.dev/encoding/json#Unmarshal

nota sobre git:

lembrar de commitar antes de git push --set-upstream (link).git (branch)
git config --global user.name "(name)"
git config --global user.email (email)

como declarar struct anônima (?)

dog := struct {
    name   string
    isGood bool
}{
    "Rex",
    true,
}
