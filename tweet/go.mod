module github.com/pooriaAcademy/event-driven-design-golang/tweet

go 1.15

replace github.com/pooriaAcademy/event-driven-design-golang/user => ../user

require (
	github.com/gorilla/mux v1.8.0
	github.com/pooriaAcademy/event-driven-design-golang/user v0.0.0-00010101000000-000000000000
)
