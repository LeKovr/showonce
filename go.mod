module it.elfire.ru/itc/showonce

go 1.18

replace SELF => ./

//replace github.com/dopos/narra => ../narra

//replace github.com/LeKovr/go-kit/logger => ../go-kit/logger

require (
	SELF v0.0.0-00010101000000-000000000000
	github.com/LeKovr/go-kit/config v0.2.0
	github.com/LeKovr/go-kit/logger v0.2.0
	github.com/dopos/narra v0.25.0
	github.com/felixge/httpsnoop v1.0.3
	github.com/go-logr/logr v1.2.3
	github.com/json-iterator/go v1.1.12
	github.com/oklog/ulid/v2 v2.0.2
	github.com/patrickmn/go-cache v2.1.0+incompatible
	golang.org/x/sync v0.1.0
)

require (
	github.com/go-logr/zerologr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/rs/zerolog v1.27.0 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/oauth2 v0.0.0-20220630143837-2104d58473e0 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/gorilla/securecookie.v1 v1.1.1 // indirect
)
