module gohttp

go 1.19

replace gohttp/server => ./server

require gohttp/server v0.0.0-00010101000000-000000000000

require (
	gohttp/lib/hi v0.0.0-00010101000000-000000000000 // indirect
	gohttp/lib/routes v0.0.0-00010101000000-000000000000 // indirect
)

replace gohttp/lib/routes => ./lib/routes

replace gohttp/lib/hi => ./lib/hi
