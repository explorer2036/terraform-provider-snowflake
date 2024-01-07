#/bin/bash

cp applications_gen_test.go ~/applications_gen_test.go

rm -f applications*_gen.go
rm -f applications*_gen_*test.go

go generate applications_def.go

goimports -w applications_impl_gen.go

cp ~/applications_gen_test.go applications_gen_test.go

go generate applications_dto_gen.go

rm -f applications_gen_integration_test.go
