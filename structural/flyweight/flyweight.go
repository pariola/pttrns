package flyweight

type TeamType int

const (
	A TeamType = iota
	B
)

type Team struct {
	Id   int32
	Name string
}

type TeamFactory struct {
	teams map[TeamType]*Team
}

func (f TeamFactory) GetTeam(t TeamType) *Team {

	team, ok := f.teams[t]

	if ok {
		return team
	}

	switch t {
	case A:
		team = &Team{
			Id:   int32(A),
			Name: "Team A",
		}
	case B:
		team = &Team{
			Id:   int32(B),
			Name: "Team B",
		}
	}

	f.teams[t] = team

	return team
}
