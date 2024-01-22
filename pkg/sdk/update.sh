#/bin/bash

cp replication_groups_gen_test.go ~/replication_groups_gen_test.go

rm -f replication_groups*_gen.go
rm -f replication_groups*_gen_*test.go

go generate replication_groups_def.go

goimports -w replication_groups_impl_gen.go

cp ~/replication_groups_gen_test.go replication_groups_gen_test.go

go generate replication_groups_dto_gen.go

rm -f replication_groups_gen_integration_test.go
