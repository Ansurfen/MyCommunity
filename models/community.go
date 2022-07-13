package models

type Community struct {
	Name    string
	Context string
	Status  int8
	Admin   string
}

type CommunityAdmin struct {
	Host   string
	Admins []string
}

type CommunityMember struct {
	Name    string
	Admins  CommunityAdmin
	Members []string
	Limit   int
}

type Application struct {
	To     string
	From   string
	Status int
}

func NewCommunity(name, context string) *Community {
	return &Community{
		Name:    name,
		Context: context,
	}
}

func NewCommunityAdmin(host string) *CommunityAdmin {
	return &CommunityAdmin{
		Host:   host,
		Admins: make([]string, 0),
	}
}

func NewCommunityMember(name string, admin CommunityAdmin, limit int) *CommunityMember {
	return &CommunityMember{
		Name:    name,
		Admins:  admin,
		Members: make([]string, 0),
		Limit:   limit,
	}
}

func NewApplication(from, to string) *Application {
	return &Application{
		From:   from,
		To:     to,
		Status: WAITING,
	}
}
