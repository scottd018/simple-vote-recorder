// Construct a class to represent a voting system.
//
// Interface:
//  // record the vote for a specific voter and initiative
//  void recordVote(string voterId, int initiativeId, boolean vote);
//
//  // return the total vote count (either for or against) as specified by "vote"
//  int getResultCount(int initiativeId, boolean vote);
//
// The class will be initialized with
//  - a list of valid voterIds
//  - a list of valid initiativeIds
//
// Notes:
//  - the system should record only one vote for each [ voterId + initiativeId ]
//  - the voterId must be valid (based on data provided during class initialization)
//  - the initiativeId must be valid (based on data provides during class initialization)
//
// Sample voting:
//  recordVote("alice", 1, True)
//  recordVote("alice", 2, False)
//  recordVote("bob", 1, True)
//  recordVote("bob", 2, True)
//  recordVote("charles", 1, True)
//  recordVote("charles", 2, False)
//
// Sample results:
//     getResultCount(1, True) → 3
//     getResultCount(1, False) → 0
//     getResultCount(2, True) → 1
//     getResultCount(2, False) → 2
//

package main

import "fmt"

func validInitiativeIds() []int {
	return []int{
		1,
		2,
	}
}

func validVoterIds() []string {
	return []string{
		"alice",
		"bob",
		"charles",
	}
}

func validateVoterId(voterId string) {
	for _, id := range validVoterIds() {
		if id == voterId {
			return
		}
	}

	panic(fmt.Errorf("%s is not a valid voter id", voterId))
}

func validateIniativeid(initiativeId int) {
	for _, id := range validInitiativeIds() {
		if id == initiativeId {
			return
		}
	}

	panic(fmt.Errorf("%v is not a valid initiative id", initiativeId))
}

func (counter *VoteCounter) recordVote(voterId string, initiativeId int, vote bool) {
	// validate the voter id
	validateVoterId(voterId)

	// validate the initiative id
	validateIniativeid(initiativeId)

	// validate that the initiative id and voter id are not currently stored
	if len(counter.votes) > 0 {
		for _, recordedVote := range counter.votes {
			if recordedVote.voterId == voterId && recordedVote.initiativeId == initiativeId {
				return
			}
		}
	}

	counter.votes = append(counter.votes, &VoteRecorder{
		voterId:      voterId,
		initiativeId: initiativeId,
		vote:         vote,
	})
}

func (counter *VoteCounter) getResultCount(intiativeId int, vote bool) int {
	var count int

	for _, voteCount := range counter.votes {
		if voteCount.initiativeId == intiativeId && voteCount.vote == vote {
			count++
		}
	}

	return count
}

type VoteCounter struct {
	votes []*VoteRecorder
}

type VoteRecorder struct {
	voterId      string
	initiativeId int
	vote         bool
}

func main() {
	// create the vote counter
	counter := &VoteCounter{}

	// record the votes
	counter.recordVote("alice", 1, true)
	counter.recordVote("alice", 2, false)
	counter.recordVote("bob", 1, true)
	counter.recordVote("bob", 2, true)
	counter.recordVote("charles", 1, true)
	counter.recordVote("charles", 2, false)

	// print the votes
	fmt.Printf("%v\n", counter.getResultCount(1, true))
	fmt.Printf("%v\n", counter.getResultCount(1, false))
	fmt.Printf("%v\n", counter.getResultCount(2, true))
	fmt.Printf("%v\n", counter.getResultCount(2, false))
}
