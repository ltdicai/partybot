package partybot

type Party struct {
	members []string
}

func CreateParty() *Party {
	return &Party{}

}

func AddToParty(party *Party, player string) *Party {
	for _, v := range party.members {
		if v == player {
			return party
		}
	}
	party.members = append(party.members, player)
	return party
}
