#/bin/bash

cp network_policies_gen_test.go ~/network_policies_gen_test.go

rm -f network_policies*_gen.go
rm -f network_policies*_gen_*test.go

go generate network_policies_def.go

goimports -w network_policies_impl_gen.go

cp ~/network_policies_gen_test.go network_policies_gen_test.go

rm -f network_policies_gen_integration_test.go

go generate network_policies_dto_gen.go
