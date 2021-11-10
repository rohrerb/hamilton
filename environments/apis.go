package environments

// API represent an API configuration for Microsoft Graph or Azure Active Directory Graph.
type Api struct {
	// The Application ID for the API.
	AppId ApiAppId

	// The endpoint for the API, including scheme.
	Endpoint ApiEndpoint
}

var (
	MsGraphGlobal = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		Endpoint: MsGraphGlobalEndpoint,
	}

	MsGraphGermany = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		Endpoint: MsGraphGermanyEndpoint,
	}

	MsGraphChina = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		Endpoint: MsGraphChinaEndpoint,
	}

	MsGraphUSGovL4 = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		Endpoint: MsGraphUSGovL4Endpoint,
	}

	MsGraphUSGovL5 = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		Endpoint: MsGraphUSGovL5Endpoint,
	}

	MsGraphCanary = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		Endpoint: MsGraphCanaryEndpoint,
	}

	AadGraphGlobal = Api{
		AppId:    PublishedApis["AzureActiveDirectoryGraph"],
		Endpoint: AadGraphGlobalEndpoint,
	}

	AadGraphGermany = Api{
		AppId:    PublishedApis["AzureActiveDirectoryGraph"],
		Endpoint: AadGraphGermanyEndpoint,
	}

	AadGraphChina = Api{
		AppId:    PublishedApis["AzureActiveDirectoryGraph"],
		Endpoint: AadGraphChinaEndpoint,
	}

	AadGraphUSGov = Api{
		AppId:    PublishedApis["AzureActiveDirectoryGraph"],
		Endpoint: AadGraphUSGovEndpoint,
	}

	ResourceManagerPublic = Api{
		AppId:    PublishedApis["AzureServiceManagement"],
		Endpoint: ResourceManagerPublicEndpoint,
	}

	ResourceManagerGermany = Api{
		AppId:    PublishedApis["AzureServiceManagement"],
		Endpoint: ResourceManagerGermanyEndpoint,
	}

	ResourceManagerChina = Api{
		AppId:    PublishedApis["AzureServiceManagement"],
		Endpoint: ResourceManagerChinaEndpoint,
	}

	ResourceManagerUSGov = Api{
		AppId:    PublishedApis["AzureServiceManagement"],
		Endpoint: ResourceManagerUSGovEndpoint,
	}

	BatchManagementPublic = Api{
		AppId:    PublishedApis["AzureBatch"],
		Endpoint: BatchManagementPublicEndpoint,
	}

	BatchManagementGermany = Api{
		AppId:    PublishedApis["AzureBatch"],
		Endpoint: BatchManagementGermanyEndpoint,
	}

	BatchManagementChina = Api{
		AppId:    PublishedApis["AzureBatch"],
		Endpoint: BatchManagementChinaEndpoint,
	}

	BatchManagementUSGov = Api{
		AppId:    PublishedApis["AzureBatch"],
		Endpoint: BatchManagementUSGovEndpoint,
	}

	DataLakePublic = Api{
		AppId:    PublishedApis["AzureDataLake"],
		Endpoint: DataLakePublicEndpoint,
	}

	KeyVaultPublic = Api{
		AppId:    PublishedApis["AzureBatch"],
		Endpoint: KeyVaultPublicEndpoint,
	}

	KeyVaultGermany = Api{
		AppId:    PublishedApis["AzureBatch"],
		Endpoint: KeyVaultGermanyEndpoint,
	}

	KeyVaultChina = Api{
		AppId:    PublishedApis["AzureBatch"],
		Endpoint: KeyVaultChinaEndpoint,
	}

	KeyVaultUSGov = Api{
		AppId:    PublishedApis["AzureBatch"],
		Endpoint: KeyVaultUSGovEndpoint,
	}

	OperationalInsightsPublic = Api{
		AppId:    PublishedApis["LogAnalytics"],
		Endpoint: OperationalInsightsPublicEndpoint,
	}

	OperationalInsightsUSGov = Api{
		AppId:    PublishedApis["LogAnalytics"],
		Endpoint: OperationalInsightsUSGovEndpoint,
	}

	OSSRDBMSPublic = Api{
		AppId:    PublishedApis["OssRdbms"],
		Endpoint: OSSRDBMSPublicEndpoint,
	}

	OSSRDBMSGermany = Api{
		AppId:    PublishedApis["OssRdbms"],
		Endpoint: OSSRDBMSGermanyEndpoint,
	}

	OSSRDBMSChina = Api{
		AppId:    PublishedApis["OssRdbms"],
		Endpoint: OSSRDBMSChinaEndpoint,
	}

	OSSRDBMSUSGov = Api{
		AppId:    PublishedApis["OssRdbms"],
		Endpoint: OSSRDBMSUSGovEndpoint,
	}

	ServiceBusPublic = Api{
		AppId:    PublishedApis["AzureServiceBus"],
		Endpoint: ServiceBusPublicEndpoint,
	}

	ServiceBusGermany = Api{
		AppId:    PublishedApis["AzureServiceBus"],
		Endpoint: ServiceBusGermanyEndpoint,
	}

	ServiceBusChina = Api{
		AppId:    PublishedApis["AzureServiceBus"],
		Endpoint: ServiceBusChinaEndpoint,
	}

	ServiceBusUSGov = Api{
		AppId:    PublishedApis["AzureServiceBus"],
		Endpoint: ServiceBusUSGovEndpoint,
	}

	ServiceManagementPublic = Api{
		AppId:    PublishedApis["AzureServiceManagement"],
		Endpoint: ServiceManagementPublicEndpoint,
	}

	ServiceManagementGermany = Api{
		AppId:    PublishedApis["AzureServiceManagement"],
		Endpoint: ServiceManagementGermanyEndpoint,
	}

	ServiceManagementChina = Api{
		AppId:    PublishedApis["AzureServiceManagement"],
		Endpoint: ServiceManagementChinaEndpoint,
	}

	ServiceManagementUSGov = Api{
		AppId:    PublishedApis["AzureServiceManagement"],
		Endpoint: ServiceManagementUSGovEndpoint,
	}

	SQLDatabasePublic = Api{
		AppId:    PublishedApis["AzureSqlDatabase"],
		Endpoint: SQLDatabasePublicEndpoint,
	}

	SQLDatabaseGermany = Api{
		AppId:    PublishedApis["AzureSqlDatabase"],
		Endpoint: SQLDatabaseGermanyEndpoint,
	}

	SQLDatabaseChina = Api{
		AppId:    PublishedApis["AzureSqlDatabase"],
		Endpoint: SQLDatabaseChinaEndpoint,
	}

	SQLDatabaseUSGov = Api{
		AppId:    PublishedApis["AzureSqlDatabase"],
		Endpoint: SQLDatabaseUSGovEndpoint,
	}

	StoragePublic = Api{
		AppId:    PublishedApis["AzureStorage"],
		Endpoint: StoragePublicEndpoint,
	}

	SynapsePublic = Api{
		AppId:    PublishedApis["AzureSynapseGateway"],
		Endpoint: SynapsePublicEndpoint,
	}
)

type ApiCliName string

const (
	MsGraphCliName         ApiCliName = "ms-graph"
	AadGraphCliName        ApiCliName = "aad-graph"
	ResourceManagerCliName ApiCliName = "arm"
	BatchCliName           ApiCliName = "batch"
	DataLakeCliName        ApiCliName = "data-lake"
	OSSRDBMSCliName        ApiCliName = "oss-rdbms"
)
