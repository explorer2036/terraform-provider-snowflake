#/bin/bash

cp event_tables_gen_test.go ~/event_tables_gen_test.go
# cp event_tables_impl_gen.go ~/event_tables_impl_gen.go
# cp event_tables_validations_gen.go ~/event_tables_validations_gen.go

rm -f event_tables*_gen.go
rm -f event_tables*_gen_*test.go

go generate event_tables_def.go

goimports -w event_tables_impl_gen.go

cp ~/event_tables_gen_test.go event_tables_gen_test.go
# cp ~/event_tables_impl_gen.go event_tables_impl_gen.go
# cp ~/event_tables_validations_gen.go event_tables_validations_gen.go

go generate event_tables_dto_gen.go

rm -f event_tables_gen_integration_test.go
