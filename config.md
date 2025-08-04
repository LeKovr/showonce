
### Main Options

| Name | ENV | Type | Default | Description |
|------|-----|------|---------|-------------|
| listen               | LISTEN               | string | `:8080` | Addr and port which server listens at (srv_listen not used) |
| listen_grpc          | LISTEN_GRPC          | string | `:8081` | Addr and port which GRPC pub server listens at |
| root                 | ROOT                 | string |  | Static files root directory |
| html                 | -                    | string | `html` | Static site subdirectory |
| priv                 | -                    | string | `/my/` | URI prefix for pages which requires auth |
| version              | -                    | bool | `false` | Show version and exit |
| config_gen           | CONFIG_GEN           | ,json,md,mk |  | Generate and print config definition in given format and exit (default: '', means skip) |
| config_dump          | CONFIG_DUMP          | string |  | Dump config dest filename |

### Logging Options {#log}

| Name | ENV | Type | Default | Description |
|------|-----|------|---------|-------------|
| log.debug            | LOG_DEBUG            | bool | `false` | Show debug info |
| log.format           | LOG_FORMAT           | ,text,json |  | Output format (default: '', means use text if DEBUG) |
| log.time_format      | LOG_TIME_FORMAT      | string | `2006-01-02 15:04:05.000` | Time format for text output |
| log.dest             | LOG_DEST             | string |  | Log destination (default: '', means STDERR) |

### Auth Service Options {#as}

| Name | ENV | Type | Default | Description |
|------|-----|------|---------|-------------|
| as.my_url            | -                    | string |  | Own host URL (autodetect if empty) |
| as.cb_url            | -                    | string | `/login` | URL for Auth server's redirect |
| as.do401             | AS_DO401             | bool | `false` | Do not redirect with http.StatusUnauthorized, process it |
| as.host              | AS_HOST              | string | `http://gitea:8080` | Authorization Server host |
| as.team              | AS_TEAM              | string | `dcape` | Authorization Server team which members has access to resource |
| as.client_id         | AS_CLIENT_ID         | string |  | Authorization Server Client ID |
| as.client_key        | AS_CLIENT_KEY        | string |  | Authorization Server Client key |
| as.cache_expire      | -                    | time.Duration | `5m` | Cache expire interval |
| as.cache_cleanup     | -                    | time.Duration | `10m` | Cache cleanup interval |
| as.client_timeout    | -                    | time.Duration | `10s` | HTTP Client timeout |
| as.auth_header       | -                    | string | `X-narra-token` | Use token from this header if given |
| as.cookie_domain     | -                    | string |  | Auth cookie domain |
| as.cookie_name       | -                    | string | `narra_token` | Auth cookie name |
| as.cookie_sign       | AS_COOKIE_SIGN_KEY   | string |  | Cookie sign key (32 or 64 bytes) |
| as.cookie_crypt      | AS_COOKIE_CRYPT_KEY  | string |  | Cookie crypt key (16, 24, or 32 bytes) |
| as.user_header       | AS_USER_HEADER       | string | `X-Username` | HTTP Response Header for username |
| as.basic_realm       | -                    | string | `narra` | Basic Auth realm |
| as.basic_username    | -                    | string | `token` | Basic Auth user name |
| as.basic_useragent   | -                    | string | `docker/` | UserAgent which requires Basic Auth |

### Endpoint Options {#as.ep}

| Name | ENV | Type | Default | Description |
|------|-----|------|---------|-------------|
| as.ep.auth           | -                    | string | `/login/oauth/authorize` | Auth URI |
| as.ep.token          | -                    | string | `/login/oauth/access_token` | Token URI |
| as.ep.user           | -                    | string | `/api/v1/user` | User info URI |
| as.ep.teams          | -                    | string | `/api/v1/user/orgs` | User teams URI |
| as.ep.team_name      | -                    | string | `username` | Teams response field name for team name |

### Storage Options {#db}

| Name | ENV | Type | Default | Description |
|------|-----|------|---------|-------------|
| db.meta_ttl          | -                    | time.Duration | `240h` | Metadata TTL |
| db.data_ttl          | -                    | time.Duration | `24h` | Data TTL |
| db.cleanup           | -                    | time.Duration | `10m` | Cleanup interval |

### Server Options {#srv}

| Name | ENV | Type | Default | Description |
|------|-----|------|---------|-------------|
| srv.listen           | SRV_LISTEN           | string | `:8080` | Addr and port which server listens at |
| srv.maxheader        | -                    | int |  | MaxHeaderBytes |
| srv.rto              | -                    | time.Duration | `10s` | HTTP read timeout |
| srv.wto              | -                    | time.Duration | `60s` | HTTP write timeout |
| srv.grace            | -                    | time.Duration | `10s` | Stop grace period |
| srv.ip_header        | SRV_IP_HEADER        | string | `X-Real-IP` | HTTP Request Header for remote IP |
| srv.user_header      | SRV_USER_HEADER      | string | `X-Username` | HTTP Request Header for username |

### HTTPS Options {#srv.tls}

| Name | ENV | Type | Default | Description |
|------|-----|------|---------|-------------|
| srv.tls.cert         | SRV_TLS_CERT         | string |  | CertFile for serving HTTPS instead HTTP |
| srv.tls.key          | SRV_TLS_KEY          | string |  | KeyFile for serving HTTPS instead HTTP |
| srv.tls.no-check     | -                    | bool | `false` | disable tls certificate validation |

### Version response Options {#srv.vr}

| Name | ENV | Type | Default | Description |
|------|-----|------|---------|-------------|
| srv.vr.prefix        | -                    | string | `/js/version.js` | URL for version response |
| srv.vr.format        | -                    | string | `document.addEventListener('DOMContentLoaded', () => { appVersion.innerText = '%s'; });\n` | Format string for version response |
| srv.vr.ctype         | -                    | string | `text/javascript` | js code Content-Type header |
