#/bin/bash

cp sequences_gen_test.go ~/sequences_gen_test.go

rm -f sequences*_gen.go
rm -f sequences*_gen_*test.go

go generate sequences_def.go

goimports -w sequences_impl_gen.go

cp ~/sequences_gen_test.go sequences_gen_test.go

go generate sequences_dto_gen.go

rm -f sequences_gen_integration_test.go
