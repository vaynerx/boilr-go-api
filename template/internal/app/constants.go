package app

// ProgramName defines this application name
const ProgramName = "{{ Project }}"

// ProgramVersion set this application version
// This is supposed to be automatically populated by the Makefile using the value from the VERSION file
// (-ldflags '-X main.ProgramVersion=$(shell cat VERSION)')
var ProgramVersion = "0.0.0"

// ProgramRelease contains this program release number (or build number)
// This is automatically populated by the Makefile using the value from the RELEASE file
// (-ldflags '-X main.ProgramRelease=$(shell cat RELEASE)')
var ProgramRelease = "0"

// ConfigPath list the paths where to look for configuration files (in order)
var ConfigPath = [...]string{
	"./",
	"config/",
	"$HOME/." + ProgramName + "/",
}

// RemoteConfigProvider is the remote configuration source ("consul", "etcd")
const RemoteConfigProvider = ""

// RemoteConfigEndpoint is the remote configuration URL (ip:port)
const RemoteConfigEndpoint = ""

// RemoteConfigPath is the remote configuration path where to search fo the configuration file ("/config/{{ Project }}")
const RemoteConfigPath = ""

// RemoteConfigSecretKeyring is the path to the openpgp secret keyring used to decript the remote configuration data ("/etc/{{ Project }}/configkey.gpg")
const RemoteConfigSecretKeyring = "" // #nosec

// EnvironmentVariablesPrefix prefix to add to the configuration environment variables
const EnvironmentVariablesPrefix = "{{ toUpper Project }}"

// ServerShutdownTimeout timeout in seconds before forcing the server to close
const ServerShutdownTimeout = 10

// ----------

// SkaffoldAPIMainHost is the network address for the api micro-service.
const SkaffoldAPIMainHost = "api-main"

// SkaffoldAPIMainPort is the network port for the api micro-service.
const SkaffoldAPIMainPort = "50000"

// SkaffoldAPIMainAddress is the network address for the api micro-service (host:port).
const SkaffoldAPIMainAddress = "api-main:50000"

// ----------

// SkaffoldSystemAccountHost is the network host for the account micro-service.
const SkaffoldSystemAccountHost = "service-account"

// SkaffoldSystemAccountPort is the network port for the account micro-service.
const SkaffoldSystemAccountPort = "60000"

// SkaffoldSystemAccountAddress is the network address for the account micro-service (host:port).
const SkaffoldSystemAccountAddress = "service-account:60000"

// ----------

// SkaffoldSystemAuthHost is the network host for the auth micro-service.
const SkaffoldSystemAuthHost = "service-auth"

// SkaffoldSystemAuthPort is the network port for the auth micro-service.
const SkaffoldSystemAuthPort = "60010"

// SkaffoldSystemAuthAddress is the network address for the auth micro-service (host:port).
const SkaffoldSystemAuthAddress = "service-auth:60010"

// ----------

// SkaffoldSystemMediaHost is the network host for the media micro-service.
const SkaffoldSystemMediaHost = "service-media"

// SkaffoldSystemMediaPort is the network port for the media micro-service.
const SkaffoldSystemMediaPort = "60020"

// SkaffoldSystemMediaAddress is the network address for the media micro-service (host:port).
const SkaffoldSystemMediaAddress = "service-media:60020"

// ----------

// SkaffoldSystemYoutubeHost is the network host for the youtube micro-service.
const SkaffoldSystemYoutubeHost = "service-youtube"

// SkaffoldSystemYoutubePort is the network port for the youtube micro-service.
const SkaffoldSystemYoutubePort = "60030"

// SkaffoldSystemYoutubeAddress is the network address for the youtube micro-service (host:port).
const SkaffoldSystemYoutubeAddress = "service-youtube:60030"

// ----------

// SkaffoldClientAuthHost is the network host for the auth front-end client.
const SkaffoldClientAuthHost = "client-auth"

// SkaffoldClientAuthPort is the network port for the auth front-end client.
const SkaffoldClientAuthPort = "60100"

// SkaffoldClientAuthAddress is the network address for the auth front-end client (host:port).
const SkaffoldClientAuthAddress = "client-auth:60100"

// ----------

// SkaffoldClientMainHost is the network host for the main front-end client.
const SkaffoldClientMainHost = "client-main"

// SkaffoldClientMainPort is the network port for the main front-end client.
const SkaffoldClientMainPort = "60110"

// SkaffoldClientMainAddress is the network address for the main front-end client (host:port).
const SkaffoldClientMainAddress = "client-main:60110"

// ----------

// SkaffoldAppHydraHost is the network host for the public Hydra service.
const SkaffoldAppHydraHost = "app-hydra"

// SkaffoldAppHydraPort is the network port for the public Hydra service.
const SkaffoldAppHydraPort = "60200"

// SkaffoldAppHydraAddress is the network address for the public Hydra service (host:port).
const SkaffoldAppHydraAddress = "app-hydra:60200"

// SkaffoldAppHydraAdminHost is the network host for the admin Hydra service.
const SkaffoldAppHydraAdminHost = "app-hydra"

// SkaffoldAppHydraAdminPort is the network port for the admin Hydra service.
const SkaffoldAppHydraAdminPort = "60201"

// SkaffoldAppHydraAdminAddress is the network address for the admin Hydra service (host:port).
const SkaffoldAppHydraAdminAddress = "app-hydra:60221"

// ----------

// SkaffoldAppPostgresqlHost is the network host for the public Postgresql service.
const SkaffoldAppPostgresqlHost = "app-postgresql"

// SkaffoldAppPostgresqlPort is the network port for the public Postgresql service.
const SkaffoldAppPostgresqlPort = "60210"

// SkaffoldAppPostgresqlAddress is the network address for the public Postgresql service (host:port).
const SkaffoldAppPostgresqlAddress = "app-postgresql:60230"

// SkaffoldAppPostgresqlUsername is the username used to authenticate requests for connection to the database.
const SkaffoldAppPostgresqlUsername = "postgresql"

// SkaffoldAppPostgresqlPassword is the password used to authenticate requests for connection to the database.
const SkaffoldAppPostgresqlPassword = "postgresql"

// ----------

// SkaffoldAppPrismalHost is the network host for the public Prismal service.
const SkaffoldAppPrismaHost = "app-prisma"

// SkaffoldAppPrismalPort is the network port for the public Prismal service.
const SkaffoldAppPrismaPort = "60220"

// SkaffoldAppPrismalAddress is the network address for the public Prismal service (host:port).
const SkaffoldAppPrismaAddress = "app-prisma:60220"

// SkaffoldAppPrismalSecret is the secret used to authenticate requests for connection to the ORM.
const SkaffoldAppPrismaSecret = "pri"
