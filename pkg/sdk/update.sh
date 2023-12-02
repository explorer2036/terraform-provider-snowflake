#/bin/bash

cp procedures_gen_test.go ~/procedures_gen_test.go
# cp procedures_impl_gen.go ~/procedures_impl_gen.go
# cp procedures_validations_gen.go ~/procedures_validations_gen.go

rm -f procedures*_gen.go
rm -f procedures*_gen_*test.go

go generate procedures_def.go

goimports -w procedures_impl_gen.go

cp ~/procedures_gen_test.go procedures_gen_test.go
# cp ~/procedures_impl_gen.go procedures_impl_gen.go
# cp ~/procedures_validations_gen.go procedures_validations_gen.go

go generate procedures_dto_gen.go

rm -f procedures_gen_integration_test.go
