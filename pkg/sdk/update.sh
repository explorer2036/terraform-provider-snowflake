#/bin/bash

cp external_functions_gen_test.go ~/external_functions_gen_test.go

rm -f external_functions*_gen.go
rm -f external_functions*_gen_*test.go

go generate external_functions_def.go

goimports -w external_functions_impl_gen.go

cp ~/external_functions_gen_test.go external_functions_gen_test.go

go generate external_functions_dto_gen.go

rm -f external_functions_gen_integration_test.go
