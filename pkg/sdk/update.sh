#/bin/bash

cp streamlits_gen_test.go ~/streamlits_gen_test.go

rm -f streamlits*_gen.go
rm -f streamlits*_gen_*test.go

go generate streamlits_def.go

goimports -w streamlits_impl_gen.go

cp ~/streamlits_gen_test.go streamlits_gen_test.go

go generate streamlits_dto_gen.go

rm -f streamlits_gen_integration_test.go
