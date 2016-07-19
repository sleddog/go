sed 's/NNN/'"$1"'/g' template.md >problem$1.go
open https://projecteuler.net/problem=$1
vi problem$1.go
