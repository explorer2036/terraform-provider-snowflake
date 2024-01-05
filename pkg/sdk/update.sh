#/bin/bash

cp application_packages_gen_test.go ~/application_packages_gen_test.go
# cp application_packages_impl_gen.go ~/application_packages_impl_gen.go
# cp application_packages_validations_gen.go ~/application_packages_validations_gen.go

rm -f application_packages*_gen.go
rm -f application_packages*_gen_*test.go

go generate application_packages_def.go

goimports -w application_packages_impl_gen.go

cp ~/application_packages_gen_test.go application_packages_gen_test.go
# cp ~/application_packages_impl_gen.go application_packages_impl_gen.go
# cp ~/application_packages_validations_gen.go application_packages_validations_gen.go

go generate application_packages_dto_gen.go

rm -f application_packages_gen_integration_test.go
