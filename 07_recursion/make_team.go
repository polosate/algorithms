package _7_recursion

type group struct {
	members  []string
	teamSize int

	result [][]string
}

func NewGroup(members []string, size int) group {
	return group{
		members:  members,
		teamSize: size,
	}
}

func (g *group) left(group, team, ind int) {
	if group == 0 || team == 0 || group < team {
		return
	}
	if group == 1 && team == 1 {
		for i := range g.result {
			g.result[i] = append(g.result[i], g.members[ind])
		}
	} else {
		g.left(group-1, team-1, ind+1)
		g.right(group-1, team, ind+1, g.result)
	}
}

func (g *group) right(group, team, ind int, result [][]string) {
	if group == 0 || team == 0 || group < team {
		return
	}
	if group == 1 && team == 1 {
		result = append(result, []string{g.members[ind]})

	} else {

	}
}

func (g *group) makeTeams() {

}
