/**
 * This represents the data model of a member score. 
 * 
 * I chose that naming instead of "members" , since this basically just represents a member's score
 * rather than data about a member. I wold save that nameing for something else
 */
package model

// MemberScore represents one member with a score property
type MemberScore struct {
	Userid int
	Score  int
}
