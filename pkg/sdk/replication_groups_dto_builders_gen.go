// Code generated by dto builder generator; DO NOT EDIT.

package sdk

import ()

func NewCreateReplicationGroupRequest(
	name AccountObjectIdentifier,
) *CreateReplicationGroupRequest {
	s := CreateReplicationGroupRequest{}
	s.name = name
	return &s
}

func (s *CreateReplicationGroupRequest) WithIfNotExists(IfNotExists *bool) *CreateReplicationGroupRequest {
	s.IfNotExists = IfNotExists
	return s
}

func (s *CreateReplicationGroupRequest) WithObjectTypes(ObjectTypes ReplicationGroupObjectTypesRequest) *CreateReplicationGroupRequest {
	s.ObjectTypes = ObjectTypes
	return s
}

func (s *CreateReplicationGroupRequest) WithAllowedDatabases(AllowedDatabases []ReplicationGroupDatabaseRequest) *CreateReplicationGroupRequest {
	s.AllowedDatabases = AllowedDatabases
	return s
}

func (s *CreateReplicationGroupRequest) WithAllowedShares(AllowedShares []ReplicationGroupShareRequest) *CreateReplicationGroupRequest {
	s.AllowedShares = AllowedShares
	return s
}

func (s *CreateReplicationGroupRequest) WithAllowedIntegrationTypes(AllowedIntegrationTypes []ReplicationGroupIntegrationTypeRequest) *CreateReplicationGroupRequest {
	s.AllowedIntegrationTypes = AllowedIntegrationTypes
	return s
}

func (s *CreateReplicationGroupRequest) WithAllowedAccounts(AllowedAccounts []ReplicationGroupAccountRequest) *CreateReplicationGroupRequest {
	s.AllowedAccounts = AllowedAccounts
	return s
}

func (s *CreateReplicationGroupRequest) WithIgnoreEditionCheck(IgnoreEditionCheck *bool) *CreateReplicationGroupRequest {
	s.IgnoreEditionCheck = IgnoreEditionCheck
	return s
}

func (s *CreateReplicationGroupRequest) WithReplicationSchedule(ReplicationSchedule *ReplicationGroupScheduleRequest) *CreateReplicationGroupRequest {
	s.ReplicationSchedule = ReplicationSchedule
	return s
}

func NewReplicationGroupObjectTypesRequest() *ReplicationGroupObjectTypesRequest {
	return &ReplicationGroupObjectTypesRequest{}
}

func (s *ReplicationGroupObjectTypesRequest) WithAccountParameters(AccountParameters *bool) *ReplicationGroupObjectTypesRequest {
	s.AccountParameters = AccountParameters
	return s
}

func (s *ReplicationGroupObjectTypesRequest) WithDatabases(Databases *bool) *ReplicationGroupObjectTypesRequest {
	s.Databases = Databases
	return s
}

func (s *ReplicationGroupObjectTypesRequest) WithIntegrations(Integrations *bool) *ReplicationGroupObjectTypesRequest {
	s.Integrations = Integrations
	return s
}

func (s *ReplicationGroupObjectTypesRequest) WithNetworkPolicies(NetworkPolicies *bool) *ReplicationGroupObjectTypesRequest {
	s.NetworkPolicies = NetworkPolicies
	return s
}

func (s *ReplicationGroupObjectTypesRequest) WithResourceMonitors(ResourceMonitors *bool) *ReplicationGroupObjectTypesRequest {
	s.ResourceMonitors = ResourceMonitors
	return s
}

func (s *ReplicationGroupObjectTypesRequest) WithRoles(Roles *bool) *ReplicationGroupObjectTypesRequest {
	s.Roles = Roles
	return s
}

func (s *ReplicationGroupObjectTypesRequest) WithShares(Shares *bool) *ReplicationGroupObjectTypesRequest {
	s.Shares = Shares
	return s
}

func (s *ReplicationGroupObjectTypesRequest) WithUsers(Users *bool) *ReplicationGroupObjectTypesRequest {
	s.Users = Users
	return s
}

func (s *ReplicationGroupObjectTypesRequest) WithWarehouses(Warehouses *bool) *ReplicationGroupObjectTypesRequest {
	s.Warehouses = Warehouses
	return s
}

func NewReplicationGroupDatabaseRequest() *ReplicationGroupDatabaseRequest {
	return &ReplicationGroupDatabaseRequest{}
}

func (s *ReplicationGroupDatabaseRequest) WithDatabase(Database string) *ReplicationGroupDatabaseRequest {
	s.Database = Database
	return s
}

func NewReplicationGroupShareRequest() *ReplicationGroupShareRequest {
	return &ReplicationGroupShareRequest{}
}

func (s *ReplicationGroupShareRequest) WithShare(Share string) *ReplicationGroupShareRequest {
	s.Share = Share
	return s
}

func NewReplicationGroupIntegrationTypeRequest() *ReplicationGroupIntegrationTypeRequest {
	return &ReplicationGroupIntegrationTypeRequest{}
}

func (s *ReplicationGroupIntegrationTypeRequest) WithIntegrationType(IntegrationType string) *ReplicationGroupIntegrationTypeRequest {
	s.IntegrationType = IntegrationType
	return s
}

func NewReplicationGroupAccountRequest() *ReplicationGroupAccountRequest {
	return &ReplicationGroupAccountRequest{}
}

func (s *ReplicationGroupAccountRequest) WithAccount(Account string) *ReplicationGroupAccountRequest {
	s.Account = Account
	return s
}

func NewReplicationGroupScheduleRequest() *ReplicationGroupScheduleRequest {
	return &ReplicationGroupScheduleRequest{}
}

func (s *ReplicationGroupScheduleRequest) WithIntervalMinutes(IntervalMinutes *ScheduleIntervalMinutesRequest) *ReplicationGroupScheduleRequest {
	s.IntervalMinutes = IntervalMinutes
	return s
}

func (s *ReplicationGroupScheduleRequest) WithCronExpression(CronExpression *ScheduleCronExpressionRequest) *ReplicationGroupScheduleRequest {
	s.CronExpression = CronExpression
	return s
}

func NewScheduleIntervalMinutesRequest(
	Minutes int,
) *ScheduleIntervalMinutesRequest {
	s := ScheduleIntervalMinutesRequest{}
	s.Minutes = Minutes
	return &s
}

func NewScheduleCronExpressionRequest(
	Expression string,
	TimeZone string,
) *ScheduleCronExpressionRequest {
	s := ScheduleCronExpressionRequest{}
	s.Expression = Expression
	s.TimeZone = TimeZone
	return &s
}

func NewCreateSecondaryReplicationGroupRequest(
	name AccountObjectIdentifier,
	Primary *ExternalObjectIdentifier,
) *CreateSecondaryReplicationGroupRequest {
	s := CreateSecondaryReplicationGroupRequest{}
	s.name = name
	s.Primary = Primary
	return &s
}

func (s *CreateSecondaryReplicationGroupRequest) WithIfNotExists(IfNotExists *bool) *CreateSecondaryReplicationGroupRequest {
	s.IfNotExists = IfNotExists
	return s
}

func NewAlterReplicationGroupRequest(
	name AccountObjectIdentifier,
) *AlterReplicationGroupRequest {
	s := AlterReplicationGroupRequest{}
	s.name = name
	return &s
}

func (s *AlterReplicationGroupRequest) WithIfExists(IfExists *bool) *AlterReplicationGroupRequest {
	s.IfExists = IfExists
	return s
}

func (s *AlterReplicationGroupRequest) WithRenameTo(RenameTo *AccountObjectIdentifier) *AlterReplicationGroupRequest {
	s.RenameTo = RenameTo
	return s
}

func (s *AlterReplicationGroupRequest) WithSet(Set *ReplicationGroupSetRequest) *AlterReplicationGroupRequest {
	s.Set = Set
	return s
}

func (s *AlterReplicationGroupRequest) WithSetIntegration(SetIntegration *ReplicationGroupSetIntegrationRequest) *AlterReplicationGroupRequest {
	s.SetIntegration = SetIntegration
	return s
}

func (s *AlterReplicationGroupRequest) WithAddDatabases(AddDatabases *ReplicationGroupAddDatabasesRequest) *AlterReplicationGroupRequest {
	s.AddDatabases = AddDatabases
	return s
}

func (s *AlterReplicationGroupRequest) WithRemoveDatabases(RemoveDatabases *ReplicationGroupRemoveDatabasesRequest) *AlterReplicationGroupRequest {
	s.RemoveDatabases = RemoveDatabases
	return s
}

func (s *AlterReplicationGroupRequest) WithMoveDatabases(MoveDatabases *ReplicationGroupMoveDatabasesRequest) *AlterReplicationGroupRequest {
	s.MoveDatabases = MoveDatabases
	return s
}

func (s *AlterReplicationGroupRequest) WithAddShares(AddShares *ReplicationGroupAddSharesRequest) *AlterReplicationGroupRequest {
	s.AddShares = AddShares
	return s
}

func (s *AlterReplicationGroupRequest) WithRemoveShares(RemoveShares *ReplicationGroupRemoveSharesRequest) *AlterReplicationGroupRequest {
	s.RemoveShares = RemoveShares
	return s
}

func (s *AlterReplicationGroupRequest) WithMoveShares(MoveShares *ReplicationGroupMoveSharesRequest) *AlterReplicationGroupRequest {
	s.MoveShares = MoveShares
	return s
}

func (s *AlterReplicationGroupRequest) WithAddAccounts(AddAccounts *ReplicationGroupAddAccountsRequest) *AlterReplicationGroupRequest {
	s.AddAccounts = AddAccounts
	return s
}

func (s *AlterReplicationGroupRequest) WithRemoveAccounts(RemoveAccounts *ReplicationGroupRemoveAccountsRequest) *AlterReplicationGroupRequest {
	s.RemoveAccounts = RemoveAccounts
	return s
}

func (s *AlterReplicationGroupRequest) WithRefresh(Refresh *bool) *AlterReplicationGroupRequest {
	s.Refresh = Refresh
	return s
}

func (s *AlterReplicationGroupRequest) WithSuspend(Suspend *bool) *AlterReplicationGroupRequest {
	s.Suspend = Suspend
	return s
}

func (s *AlterReplicationGroupRequest) WithResume(Resume *bool) *AlterReplicationGroupRequest {
	s.Resume = Resume
	return s
}

func NewReplicationGroupSetRequest() *ReplicationGroupSetRequest {
	return &ReplicationGroupSetRequest{}
}

func (s *ReplicationGroupSetRequest) WithObjectTypes(ObjectTypes *ReplicationGroupObjectTypesRequest) *ReplicationGroupSetRequest {
	s.ObjectTypes = ObjectTypes
	return s
}

func (s *ReplicationGroupSetRequest) WithAllowedDatabases(AllowedDatabases []ReplicationGroupDatabaseRequest) *ReplicationGroupSetRequest {
	s.AllowedDatabases = AllowedDatabases
	return s
}

func (s *ReplicationGroupSetRequest) WithAllowedShares(AllowedShares []ReplicationGroupShareRequest) *ReplicationGroupSetRequest {
	s.AllowedShares = AllowedShares
	return s
}

func (s *ReplicationGroupSetRequest) WithReplicationSchedule(ReplicationSchedule *ReplicationGroupScheduleRequest) *ReplicationGroupSetRequest {
	s.ReplicationSchedule = ReplicationSchedule
	return s
}

func (s *ReplicationGroupSetRequest) WithEnableEtlReplication(EnableEtlReplication *bool) *ReplicationGroupSetRequest {
	s.EnableEtlReplication = EnableEtlReplication
	return s
}

func NewReplicationGroupSetIntegrationRequest() *ReplicationGroupSetIntegrationRequest {
	return &ReplicationGroupSetIntegrationRequest{}
}

func (s *ReplicationGroupSetIntegrationRequest) WithObjectTypes(ObjectTypes *ReplicationGroupObjectTypesRequest) *ReplicationGroupSetIntegrationRequest {
	s.ObjectTypes = ObjectTypes
	return s
}

func (s *ReplicationGroupSetIntegrationRequest) WithAllowedIntegrationTypes(AllowedIntegrationTypes []ReplicationGroupIntegrationTypeRequest) *ReplicationGroupSetIntegrationRequest {
	s.AllowedIntegrationTypes = AllowedIntegrationTypes
	return s
}

func (s *ReplicationGroupSetIntegrationRequest) WithReplicationSchedule(ReplicationSchedule *ReplicationGroupScheduleRequest) *ReplicationGroupSetIntegrationRequest {
	s.ReplicationSchedule = ReplicationSchedule
	return s
}

func NewReplicationGroupAddDatabasesRequest() *ReplicationGroupAddDatabasesRequest {
	return &ReplicationGroupAddDatabasesRequest{}
}

func (s *ReplicationGroupAddDatabasesRequest) WithDatabases(Databases []ReplicationGroupDatabaseRequest) *ReplicationGroupAddDatabasesRequest {
	s.Databases = Databases
	return s
}

func NewReplicationGroupRemoveDatabasesRequest() *ReplicationGroupRemoveDatabasesRequest {
	return &ReplicationGroupRemoveDatabasesRequest{}
}

func (s *ReplicationGroupRemoveDatabasesRequest) WithDatabases(Databases []ReplicationGroupDatabaseRequest) *ReplicationGroupRemoveDatabasesRequest {
	s.Databases = Databases
	return s
}

func NewReplicationGroupMoveDatabasesRequest() *ReplicationGroupMoveDatabasesRequest {
	return &ReplicationGroupMoveDatabasesRequest{}
}

func (s *ReplicationGroupMoveDatabasesRequest) WithDatabases(Databases []ReplicationGroupDatabaseRequest) *ReplicationGroupMoveDatabasesRequest {
	s.Databases = Databases
	return s
}

func (s *ReplicationGroupMoveDatabasesRequest) WithMoveTo(MoveTo *AccountObjectIdentifier) *ReplicationGroupMoveDatabasesRequest {
	s.MoveTo = MoveTo
	return s
}

func NewReplicationGroupAddSharesRequest() *ReplicationGroupAddSharesRequest {
	return &ReplicationGroupAddSharesRequest{}
}

func (s *ReplicationGroupAddSharesRequest) WithShares(Shares []ReplicationGroupShareRequest) *ReplicationGroupAddSharesRequest {
	s.Shares = Shares
	return s
}

func NewReplicationGroupRemoveSharesRequest() *ReplicationGroupRemoveSharesRequest {
	return &ReplicationGroupRemoveSharesRequest{}
}

func (s *ReplicationGroupRemoveSharesRequest) WithShares(Shares []ReplicationGroupShareRequest) *ReplicationGroupRemoveSharesRequest {
	s.Shares = Shares
	return s
}

func NewReplicationGroupMoveSharesRequest() *ReplicationGroupMoveSharesRequest {
	return &ReplicationGroupMoveSharesRequest{}
}

func (s *ReplicationGroupMoveSharesRequest) WithShares(Shares []ReplicationGroupShareRequest) *ReplicationGroupMoveSharesRequest {
	s.Shares = Shares
	return s
}

func (s *ReplicationGroupMoveSharesRequest) WithMoveTo(MoveTo *AccountObjectIdentifier) *ReplicationGroupMoveSharesRequest {
	s.MoveTo = MoveTo
	return s
}

func NewReplicationGroupAddAccountsRequest() *ReplicationGroupAddAccountsRequest {
	return &ReplicationGroupAddAccountsRequest{}
}

func (s *ReplicationGroupAddAccountsRequest) WithAccounts(Accounts []ReplicationGroupAccountRequest) *ReplicationGroupAddAccountsRequest {
	s.Accounts = Accounts
	return s
}

func (s *ReplicationGroupAddAccountsRequest) WithIgnoreEditionCheck(IgnoreEditionCheck *bool) *ReplicationGroupAddAccountsRequest {
	s.IgnoreEditionCheck = IgnoreEditionCheck
	return s
}

func NewReplicationGroupRemoveAccountsRequest() *ReplicationGroupRemoveAccountsRequest {
	return &ReplicationGroupRemoveAccountsRequest{}
}

func (s *ReplicationGroupRemoveAccountsRequest) WithAccounts(Accounts []ReplicationGroupAccountRequest) *ReplicationGroupRemoveAccountsRequest {
	s.Accounts = Accounts
	return s
}

func NewShowReplicationGroupRequest() *ShowReplicationGroupRequest {
	return &ShowReplicationGroupRequest{}
}

func (s *ShowReplicationGroupRequest) WithInAccount(InAccount *AccountObjectIdentifier) *ShowReplicationGroupRequest {
	s.InAccount = InAccount
	return s
}

func NewShowDatabasesInReplicationGroupRequest(
	name AccountObjectIdentifier,
) *ShowDatabasesInReplicationGroupRequest {
	s := ShowDatabasesInReplicationGroupRequest{}
	s.name = name
	return &s
}

func NewShowSharesInReplicationGroupRequest(
	name AccountObjectIdentifier,
) *ShowSharesInReplicationGroupRequest {
	s := ShowSharesInReplicationGroupRequest{}
	s.name = name
	return &s
}

func NewDropReplicationGroupRequest(
	name AccountObjectIdentifier,
) *DropReplicationGroupRequest {
	s := DropReplicationGroupRequest{}
	s.name = name
	return &s
}

func (s *DropReplicationGroupRequest) WithIfExists(IfExists *bool) *DropReplicationGroupRequest {
	s.IfExists = IfExists
	return s
}
