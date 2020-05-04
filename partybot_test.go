package partybot

import (
	"testing"
)

func TestPartyBot(t *testing.T) {
	t.Run("creates empty party", func(t *testing.T) {
		// Given

		// When
		party := CreateParty()

		// Then
		if len(party.members) > 0 {
			t.Errorf("A new party should not have members")
		}
	})

	t.Run("add player to new party", func(t *testing.T) {
		//Given
		newPlayer := "Denise"
		party := CreateParty()

		// When
		AddToParty(party, newPlayer) // This could be a method instead of a function

		// Then
		if len(party.members) != 1 {
			t.Errorf("No player was added to party")
		}

		found := false
		for _, v := range party.members {
			if v == newPlayer {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%v was not added to new party %v", newPlayer, party.members)
		}
	})

	t.Run("add player to populated party", func(t *testing.T) {
		//Given
		newPlayer := "Linus"
		party := CreateParty()
		AddToParty(party, "Denise")

		// When
		AddToParty(party, newPlayer) // This could be a method instead of a function

		// Then
		if len(party.members) != 2 {
			t.Errorf("New player was not added to party")
		}

		found := false
		for _, v := range party.members {
			if v == newPlayer {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%v was not added to new party %v", newPlayer, party.members)
		}
	})

	t.Run("add player twice party", func(t *testing.T) {
		//Given
		doublePlayer := "Chell"
		party := CreateParty()
		AddToParty(party, doublePlayer)

		// When
		AddToParty(party, doublePlayer) // This could be a method instead of a function

		// Then
		if len(party.members) != 1 {
			t.Errorf("Player %v was added twice", doublePlayer)
		}

	})
}
