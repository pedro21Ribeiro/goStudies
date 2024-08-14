module example.com/hello

go 1.22.5

replace example.com/greeting => ../greetings

require example.com/greeting v0.0.0-00010101000000-000000000000
