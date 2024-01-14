#/bin/bash

cp network_rules_gen_test.go ~/network_rules_gen_test.go

rm -f network_rules*_gen.go
rm -f network_rules*_gen_*test.go

go generate network_rules_def.go

goimports -w network_rules_impl_gen.go

cp ~/network_rules_gen_test.go network_rules_gen_test.go

go generate network_rules_dto_gen.go

rm -f network_rules_gen_integration_test.go
