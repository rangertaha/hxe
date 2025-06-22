/*
Copyright © 2023 Rangertaha  <rangertaha@gmail.com>
*/
package config

import (
	"github.com/nats-io/nats-server/v2/server"
)

type Broker struct {
	Name string `hcl:"name,optional"`
	Host string `hcl:"addr,optional"`
	Port int    `hcl:"port,optional"`
	// DontListen            bool   `hcl:"dont_listen"`
	// ClientAdvertise       string `hcl:",optional"`
	// Trace                 bool   `hcl:",optional"`
	Debug bool `hcl:"debug,optional"`
	// TraceVerbose          bool   `hcl:",optional"`
	// NoLog                 bool   `hcl:",optional"`
	// NoSigs                bool   `hcl:",optional"`
	// NoSublistCache        bool   `hcl:",optional"`
	// NoHeaderSupport       bool   `hcl:",optional"`
	// DisableShortFirstPing bool   `hcl:",optional"`
	// Logtime               bool   `hcl:",optional"`
	// LogtimeUTC            bool   `hcl:",optional"`
	// MaxConn               int    `hcl:"max_connections"`
	// MaxSubs               int    `hcl:"max_subscriptions,optional"`
	// MaxSubTokens          uint8  `hcl:",optional"`
	// // Nkeys                      []*NkeyUser   `hcl:",optional"`
	// // Users                      []*User       `hcl:",optional"`
	// // Accounts                   []*Account    `hcl:",optional"`
	// NoAuthUser      string `hcl:",optional"`
	// SystemAccount   string `hcl:",optional"`
	// NoSystemAccount bool   `hcl:",optional"`
	// Username        string `hcl:",optional"`
	// Password        string `hcl:",optional"`
	// Authorization   string `hcl:",optional"`
	// // AuthCallout                *AuthCallout  `hcl:",optional"`
	// PingInterval        time.Duration `hcl:"ping_interval"`
	// MaxPingsOut         int           `hcl:"ping_max"`
	// HTTPHost            string        `hcl:"http_host"`
	// HTTPPort            int           `hcl:"http_port"`
	// HTTPBasePath        string        `hcl:"http_base_path"`
	// HTTPSPort           int           `hcl:"https_port"`
	// AuthTimeout         float64       `hcl:"auth_timeout"`
	// MaxControlLine      int32         `hcl:"max_control_line"`
	// MaxPayload          int32         `hcl:"max_payload"`
	// MaxPending          int64         `hcl:"max_pending"`
	// NoFastProducerStall bool          `hcl:",optional"`
	// // Cluster             ClusterOpts   `hcl:"cluster,block"`
	// // Gateway                    GatewayOpts   `hcl:"gateway,optional"`
	// // LeafNode                   LeafNodeOpts  `hcl:"leaf,optional"`
	// JetStream          bool   `hcl:"jetstream"`
	// JetStreamStrict    bool   `hcl:",optional"`
	// JetStreamMaxMemory int64  `hcl:",optional"`
	// JetStreamMaxStore  int64  `hcl:",optional"`
	// JetStreamDomain    string `hcl:",optional"`
	// JetStreamExtHint   string `hcl:",optional"`
	// JetStreamKey       string `hcl:",optional"`
	// JetStreamOldKey    string `hcl:",optional"`
	// // JetStreamCipher            StoreCipher   `hcl:",optional"`
	// JetStreamUniqueTag string
	// // JetStreamLimits            JSLimitOpts
	// // JetStreamTpm               JSTpmOpts
	// JetStreamMaxCatchup        int64
	// JetStreamRequestQueueLimit int64
	// StreamMaxBufferedMsgs      int               `hcl:",optional"`
	// StreamMaxBufferedSize      int64             `hcl:",optional"`
	// StoreDir                   string            `hcl:",optional"`
	// SyncInterval               time.Duration     `hcl:",optional"`
	// SyncAlways                 bool              `hcl:",optional"`
	// JsAccDefaultDomain         map[string]string `hcl:",optional"` // account to domain name mapping
	// // Websocket                  WebsocketOpts     `hcl:",optional"`
	// // MQTT                       MQTTOpts          `hcl:",optional"`
	// ProfPort      int    `hcl:",optional"`
	// ProfBlockRate int    `hcl:",optional"`
	// PidFile       string `hcl:",optional"`
	// PortsFileDir  string `hcl:",optional"`
	// LogFile       string `hcl:",optional"`
	// LogSizeLimit  int64  `hcl:",optional"`
	// LogMaxFiles   int64  `hcl:",optional"`
	// Syslog        bool   `hcl:",optional"`
	// RemoteSyslog  string `hcl:",optional"`
	// // Routes                     []*url.URL        `hcl:",optional"`
	// RoutesStr  string  `hcl:",optional"`
	// TLSTimeout float64 `hcl:"tls_timeout"`
	// TLS        bool    `hcl:",optional"`
	// TLSVerify  bool    `hcl:",optional"`
	// TLSMap     bool    `hcl:",optional"`
	// TLSCert    string  `hcl:",optional"`
	// TLSKey     string  `hcl:",optional"`
	// TLSCaCert  string  `hcl:",optional"`
	// // TLSConfig                  *tls.Config       `hcl:",optional"`
	// // TLSPinnedCerts             PinnedCertSet     `hcl:",optional"`
	// TLSRateLimit int64 `hcl:",optional"`
	// // When set to true, the server will perform the TLS handshake before
	// // sending the INFO protocol. For clients that are not configured
	// // with a similar option, their connection will fail with some sort
	// // of timeout or EOF error since they are expecting to receive an
	// // INFO protocol first.
	// TLSHandshakeFirst bool `hcl:",optional"`
	// // If TLSHandshakeFirst is true and this value is strictly positive,
	// // the server will wait for that amount of time for the TLS handshake
	// // to start before falling back to previous behavior of sending the
	// // INFO protocol first. It allows for a mix of newer clients that can
	// // require a TLS handshake first, and older clients that can't.
	// TLSHandshakeFirstFallback time.Duration `hcl:",optional"`
	// AllowNonTLS               bool          `hcl:",optional"`
	// WriteDeadline             time.Duration `hcl:",optional"`
	// MaxClosedClients          int           `hcl:",optional"`
	// LameDuckDuration          time.Duration `hcl:",optional"`
	// LameDuckGracePeriod       time.Duration `hcl:",optional"`

	// // MaxTracedMsgLen is the maximum printable length for traced messages.
	// MaxTracedMsgLen int `hcl:",optional"`

	// // Operating a trusted NATS server
	// TrustedKeys []string `hcl:",optional"`
	// // TrustedOperators         []*jwt.OperatorClaims `hcl:",optional"`
	// // AccountResolver          AccountResolver       `hcl:",optional"`
	// // AccountResolverTLSConfig *tls.Config           `hcl:",optional"`

	// // AlwaysEnableNonce will always present a nonce to new connections
	// // typically used by custom Authentication implementations who embeds
	// // the server and so not presented as a configuration option
	// AlwaysEnableNonce bool

	// // CustomClientAuthentication Authentication `hcl:",optional"`
	// // CustomRouterAuthentication Authentication `hcl:",optional"`

	// // CheckConfig configuration file syntax test was successful and exit.
	// CheckConfig bool `hcl:",optional"`

	// // DisableJetStreamBanner will not print the ascii art on startup for JetStream enabled servers
	// DisableJetStreamBanner bool `hcl:",optional"`

	// // ConnectErrorReports specifies the number of failed attempts
	// // at which point server should report the failure of an initial
	// // connection to a route, gateway or leaf node.
	// // See DEFAULT_CONNECT_ERROR_REPORTS for default value.
	// ConnectErrorReports int

	// // ReconnectErrorReports is similar to ConnectErrorReports except
	// // that this applies to reconnect events.
	// ReconnectErrorReports int

	// // Tags describing the server. They will be included in varz
	// // and used as a filter criteria for some system requests.
	// // Tags jwt.TagList `hcl:",optional"`

	// // OCSPConfig enables OCSP Stapling in the server.
	// // OCSPConfig    *OCSPConfig
	// // tlsConfigOpts *TLSConfigOpts

	// // private fields, used to know if bool options are explicitly
	// // defined in config and/or command line params.
	// inConfig  map[string]bool
	// inCmdLine map[string]bool

	// // private fields for operator mode
	// operatorJWT            []string
	// resolverPreloads       map[string]string
	// resolverPinnedAccounts map[string]struct{}

	// // private fields, used for testing
	// gatewaysSolicitDelay time.Duration
	// overrideProto        int

	// // JetStream
	// maxMemSet   bool
	// maxStoreSet bool
	// syncSet     bool

	// // OCSP Cache config enables next-gen cache for OCSP features
	// // OCSPCacheConfig *OCSPResponseCacheConfigl

	// // Used to mark that we had a top level authorization block.
	// authBlockDefined bool

	// // configDigest represents the state of configuration.
	// configDigest string
}

func (b *Broker) Options() *server.Options {
	return &server.Options{}
}

type ClusterOpts struct{}

// 	// creator, err := broker.Get(config.Tdb.Type)
// 	// if err != nil {
// 	// 	log.Error().Err(err).Msgf("Unable to load config file")
// 	// }

// 	// Tdb := creator()
// 	// gohcl.DecodeBody(config.Tdb.Config, nil, Tdb)
// 	// // gohcl.EncodeIntoBody()

// 	return config
// }

// func LoadAgent(paths ...string) (agent *Agent, err error) {
// 	for _, path := range paths {
// 		if path != "" {
// 			log.Info().Msgf("Loading agent config: %s", path)

// 			if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "$HOME/") {
// 				if home, err := os.UserHomeDir(); err == nil {
// 					path = strings.ReplaceAll(path, "~", home)
// 					path = strings.ReplaceAll(path, "$HOME", home)
// 				}
// 			}
// 			if _, err = os.Stat(path); err != nil {
// 				log.Error().Err(err).Msgf("File not found: %s", path)
// 				return nil, err
// 			}

// 			agent = NewAgent()
// 			if err = agent.Load(path); err != nil {
// 				log.Error().Err(err).Msgf("Unable to load config file: %s", path)
// 				return nil, err
// 			}
// 			return

// 		}

// 	}

// 	// if err != nil {
// 	// 	log.Error().Err(err).Msg("Unable to laod config file")
// 	// }

// 	returnj
// }

// func (a *Agent) Load(filepath string) (err error) {
// 	return hclsimple.DecodeFile(filepath, utils.CtxFunctions, &AgentConfig{Config: a})
// }

// func (a *Agent) Save(filepath string) (err error) {
// 	f := hclwrite.NewEmptyFile()
// 	block := gohcl.EncodeAsBlock(a, `agent`)
// 	f.Body().AppendBlock(block)

// 	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
// 	file.Write(f.Bytes())

// 	return
// }
