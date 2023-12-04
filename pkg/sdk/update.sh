#/bin/bash

cp functions_gen_test.go ~/functions_gen_test.go

rm -f functions*_gen.go
rm -f functions*_gen_*test.go

go generate functions_def.go

goimports -w functions_impl_gen.go

cp ~/functions_gen_test.go functions_gen_test.go

rm -f functions_gen_integration_test.go
