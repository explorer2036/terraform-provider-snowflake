#/bin/bash

cp application_roles_gen_test.go ~/application_roles_gen_test.go
# cp application_roles_validations_gen.go ~/application_roles_validations_gen.go
# cp application_roles_impl_gen.go ~/application_roles_impl_gen.go

rm -f application_roles*_gen.go
rm -f application_roles*_gen_*test.go

go generate application_roles_def.go

goimports -w application_roles_impl_gen.go

cp ~/application_roles_gen_test.go application_roles_gen_test.go
# cp ~/application_roles_validations_gen.go application_roles_validations_gen.go
# cp ~/application_roles_impl_gen.go application_roles_impl_gen.go

go generate application_roles_dto_gen.go

rm -f application_roles_gen_integration_test.go
