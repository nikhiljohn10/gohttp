module example/http

go 1.19

replace example/hi => ../lib/hi

replace example/routes => ../lib/routes

require example/routes v0.0.0-00010101000000-000000000000

require example/hi v0.0.0-00010101000000-000000000000 // indirect
