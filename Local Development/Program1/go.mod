module github.com/1Shubham7/hellogo

go 1.20

replace github.com/1Shubham/packageTwo v0.0.0 => ./packageTwo

require (
	github.com/1Shubham/packageTwo v0.0.0
)