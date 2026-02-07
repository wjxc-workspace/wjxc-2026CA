// Copyright 2019 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Game-specific audience sound timings.

package game

type MatchSound struct {
	Name          string
	FileExtension string
	MatchTimeSec  float64
}

// List of sounds and how many seconds into the match they are played. A negative time indicates that the sound can only
// be triggered explicitly.
var MatchSounds []*MatchSound

func UpdateMatchSounds() {
	MatchSounds = []*MatchSound{
		{
			"start",
			"wav",
			0,
		},
		{
			"end",
			"wav",
			float64(MatchTiming.AutoDurationSec) - 0.5,
		},
		{
			"resume",
			"wav",
			float64(MatchTiming.AutoDurationSec),
		},
		{
			"linear_popping",
			"wav",
			float64(MatchTiming.AutoDurationSec + 10),
		},
		{
			"linear_popping",
			"wav",
			float64(MatchTiming.AutoDurationSec + 35),
		},
		{
			"linear_popping",
			"wav",
			float64(MatchTiming.AutoDurationSec + 60),
		},
		{
			"linear_popping",
			"wav",
			float64(MatchTiming.AutoDurationSec + 85),
		},
		{
			"warning",
			"wav",
			float64(
				MatchTiming.AutoDurationSec + MatchTiming.PauseDurationSec + MatchTiming.TeleopDurationSec - MatchTiming.WarningRemainingDurationSec,
			),
		},
		{
			"end",
			"wav",
			float64(MatchTiming.AutoDurationSec + MatchTiming.PauseDurationSec + MatchTiming.TeleopDurationSec),
		},
		{
			"abort",
			"wav",
			-1,
		},
		{
			"match_result",
			"wav",
			-1,
		},
		{
			"pick_clock",
			"wav",
			-1,
		},
		{
			"pick_clock_expired",
			"wav",
			-1,
		},
	}
}
