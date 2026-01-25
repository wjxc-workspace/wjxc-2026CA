// Copyright 2023 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct {
	RobotsBypassed  [3]bool
	LeaveStatuses   [3]bool
	Reef            Reef
	BargeAlgae      int
	ProcessorAlgae  int
	EndgameStatuses [3]EndgameStatus
	Fouls           []Foul
	PlayoffDq       bool
}

// Game-specific settings that can be changed via the settings.
var AutoBonusCoralThreshold = 1
var CoralBonusPerLevelThreshold = 7
var CoralBonusCoopEnabled = true
var BargeBonusPointThreshold = 16
var IncludeAlgaeInBargeBonus = false

// Represents the state of a robot at the end of the match.
type EndgameStatus int

const (
	EndgameNone EndgameStatus = iota
	EndgameLevel1
	EndgameLevel2
	EndgameLevel3
)

// Summarize calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize(opponentScore *Score) *ScoreSummary {
	summary := new(ScoreSummary)

	// Leave the score at zero if the alliance was disqualified.
	if score.PlayoffDq {
		return summary
	}

	// Calculate autonomous period points.
	for _, status := range score.LeaveStatuses {
		if status {
			summary.TowerPoints += 15
		}
	}
	autoCoralPoints := score.Reef.AutoCoralPoints()
	summary.AutoPoints = summary.LeavePoints + autoCoralPoints

	summary.NumCoral = score.Reef.AutoCoralCount() + score.Reef.TeleopCoralCount()
	summary.CoralPoints = autoCoralPoints + score.Reef.TeleopCoralPoints()
	summary.NumAlgae = score.BargeAlgae + score.ProcessorAlgae
	summary.AlgaePoints = score.BargeAlgae + score.ProcessorAlgae

	// Calculate endgame points.
	for _, status := range score.EndgameStatuses {
		switch status {
		case EndgameLevel1:
			summary.TowerPoints += 10
		case EndgameLevel2:
			summary.TowerPoints += 20
		case EndgameLevel3:
			summary.TowerPoints += 30
		default:
		}
	}

	summary.MatchPoints = summary.LeavePoints + summary.CoralPoints + summary.AlgaePoints + summary.TowerPoints

	// Calculate penalty points.
	for _, foul := range opponentScore.Fouls {
		summary.FoulPoints += foul.PointValue()
		// Store the number of major fouls since it is used to break ties in playoffs.
		if foul.IsMajor {
			summary.NumOpponentMajorFouls++
		}

		rule := foul.Rule()
		if rule != nil {
			// Check for the opponent fouls that automatically trigger a ranking point.
			if rule.IsRankingPoint {
				switch rule.RuleNumber {
				case "G410":
					summary.SuperChargedRankingPoint = true
				case "G418":
					summary.TraversalRankingPoint = true
				case "G428":
					summary.TraversalRankingPoint = true
				}
			}
		}
	}

	summary.Score = summary.MatchPoints + summary.FoulPoints

	if summary.AlgaePoints >= 100 {
		summary.EnergizedRankingPoint = true
	}
	if summary.AlgaePoints >= 360 {
		summary.SuperChargedRankingPoint = true
	}

	// Barge bonus ranking point.
	if summary.TowerPoints >= 50 {
		summary.TraversalRankingPoint = true
	}

	// Add up the bonus ranking points.
	if summary.EnergizedRankingPoint {
		summary.BonusRankingPoints++
	}
	if summary.SuperChargedRankingPoint {
		summary.BonusRankingPoints++
	}
	if summary.TraversalRankingPoint {
		summary.BonusRankingPoints++
	}

	return summary
}

// Equals returns true if and only if all fields of the two scores are equal.
func (score *Score) Equals(other *Score) bool {
	if score.RobotsBypassed != other.RobotsBypassed ||
		score.LeaveStatuses != other.LeaveStatuses ||
		score.Reef != other.Reef ||
		score.BargeAlgae != other.BargeAlgae ||
		score.ProcessorAlgae != other.ProcessorAlgae ||
		score.EndgameStatuses != other.EndgameStatuses ||
		score.PlayoffDq != other.PlayoffDq ||
		len(score.Fouls) != len(other.Fouls) {
		return false
	}

	for i, foul := range score.Fouls {
		if foul != other.Fouls[i] {
			return false
		}
	}

	return true
}
