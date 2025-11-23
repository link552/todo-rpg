package game

import (
	"time"
	"sort"
	"todorpg/internal/core"
)

type expRow struct {
	Level int
	TotalExp int
	NextExp int
}

func getExpTable() []expRow{
	return []expRow{
		{Level: 1, TotalExp: 0, NextExp: 83},
		{Level: 2, TotalExp: 83, NextExp: 174},
		{Level: 3, TotalExp: 174, NextExp: 276},
		{Level: 4, TotalExp: 276, NextExp: 388},
		{Level: 5, TotalExp: 388, NextExp: 512},
		{Level: 6, TotalExp: 512, NextExp: 650},
		{Level: 7, TotalExp: 650, NextExp: 801},
		{Level: 8, TotalExp: 801, NextExp: 969},
		{Level: 9, TotalExp: 969, NextExp: 1154},
		{Level: 10, TotalExp: 1154, NextExp: 1358},
		{Level: 11, TotalExp: 1358, NextExp: 1584},
		{Level: 12, TotalExp: 1584, NextExp: 1833},
		{Level: 13, TotalExp: 1833, NextExp: 2107},
		{Level: 14, TotalExp: 2107, NextExp: 2411},
		{Level: 15, TotalExp: 2411, NextExp: 2746},
		{Level: 16, TotalExp: 2746, NextExp: 3115},
		{Level: 17, TotalExp: 3115, NextExp: 3523},
		{Level: 18, TotalExp: 3523, NextExp: 3973},
		{Level: 19, TotalExp: 3973, NextExp: 4470},
		{Level: 20, TotalExp: 4470, NextExp: 5018},
		{Level: 21, TotalExp: 5018, NextExp: 5624},
		{Level: 22, TotalExp: 5624, NextExp: 6291},
		{Level: 23, TotalExp: 6291, NextExp: 7028},
		{Level: 24, TotalExp: 7028, NextExp: 7842},
		{Level: 25, TotalExp: 7842, NextExp: 8740},
		{Level: 26, TotalExp: 8740, NextExp: 9730},
		{Level: 27, TotalExp: 9730, NextExp: 10824},
		{Level: 28, TotalExp: 10824, NextExp: 12031},
		{Level: 29, TotalExp: 12031, NextExp: 13363},
		{Level: 30, TotalExp: 13363, NextExp: 14833},
		{Level: 31, TotalExp: 14833, NextExp: 16456},
		{Level: 32, TotalExp: 16456, NextExp: 18247},
		{Level: 33, TotalExp: 18247, NextExp: 20224},
		{Level: 34, TotalExp: 20224, NextExp: 22406},
		{Level: 35, TotalExp: 22406, NextExp: 24815},
		{Level: 36, TotalExp: 24815, NextExp: 27473},
		{Level: 37, TotalExp: 27473, NextExp: 30408},
		{Level: 38, TotalExp: 30408, NextExp: 33648},
		{Level: 39, TotalExp: 33648, NextExp: 37224},
		{Level: 40, TotalExp: 37224, NextExp: 41171},
		{Level: 41, TotalExp: 41171, NextExp: 45529},
		{Level: 42, TotalExp: 45529, NextExp: 50339},
		{Level: 43, TotalExp: 50339, NextExp: 55649},
		{Level: 44, TotalExp: 55649, NextExp: 61512},
		{Level: 45, TotalExp: 61512, NextExp: 67983},
		{Level: 46, TotalExp: 67983, NextExp: 75127},
		{Level: 47, TotalExp: 75127, NextExp: 83014},
		{Level: 48, TotalExp: 83014, NextExp: 91721},
		{Level: 49, TotalExp: 91721, NextExp: 101333},
		{Level: 50, TotalExp: 101333, NextExp: 111945},
		{Level: 51, TotalExp: 111945, NextExp: 123660},
		{Level: 52, TotalExp: 123660, NextExp: 136594},
		{Level: 53, TotalExp: 136594, NextExp: 150872},
		{Level: 54, TotalExp: 150872, NextExp: 166636},
		{Level: 55, TotalExp: 166636, NextExp: 184040},
		{Level: 56, TotalExp: 184040, NextExp: 203254},
		{Level: 57, TotalExp: 203254, NextExp: 224466},
		{Level: 58, TotalExp: 224466, NextExp: 247886},
		{Level: 59, TotalExp: 247886, NextExp: 273742},
		{Level: 60, TotalExp: 273742, NextExp: 302288},
		{Level: 61, TotalExp: 302288, NextExp: 333804},
		{Level: 62, TotalExp: 333804, NextExp: 368599},
		{Level: 63, TotalExp: 368599, NextExp: 407015},
		{Level: 64, TotalExp: 407015, NextExp: 449428},
		{Level: 65, TotalExp: 449428, NextExp: 496254},
		{Level: 66, TotalExp: 496254, NextExp: 547953},
		{Level: 67, TotalExp: 547953, NextExp: 605032},
		{Level: 68, TotalExp: 605032, NextExp: 668051},
		{Level: 69, TotalExp: 668051, NextExp: 737627},
		{Level: 70, TotalExp: 737627, NextExp: 814445},
		{Level: 71, TotalExp: 814445, NextExp: 899257},
		{Level: 72, TotalExp: 899257, NextExp: 992865},
		{Level: 73, TotalExp: 992895, NextExp: 1096278},
		{Level: 74, TotalExp: 1096278, NextExp: 1210421},
		{Level: 75, TotalExp: 1210421, NextExp: 1336443},
		{Level: 76, TotalExp: 1336443, NextExp: 1475581},
		{Level: 77, TotalExp: 1475581, NextExp: 1629200},
		{Level: 78, TotalExp: 1629200, NextExp: 1798808},
		{Level: 79, TotalExp: 1798808, NextExp: 1986068},
		{Level: 80, TotalExp: 1986068, NextExp: 2192818},
		{Level: 81, TotalExp: 2192818, NextExp: 2421087},
		{Level: 82, TotalExp: 2421087, NextExp: 2673114},
		{Level: 83, TotalExp: 2673114, NextExp: 2951373},
		{Level: 84, TotalExp: 2951373, NextExp: 3258594},
		{Level: 85, TotalExp: 3258594, NextExp: 3597792},
		{Level: 86, TotalExp: 3597792, NextExp: 3972294},
		{Level: 87, TotalExp: 3972294, NextExp: 4385776},
		{Level: 88, TotalExp: 4385776, NextExp: 4842295},
		{Level: 89, TotalExp: 4842295, NextExp: 5346332},
		{Level: 90, TotalExp: 5346332, NextExp: 5902831},
		{Level: 91, TotalExp: 5902831, NextExp: 6517253},
		{Level: 92, TotalExp: 6517253, NextExp: 7195629},
		{Level: 93, TotalExp: 7195629, NextExp: 7944614},
		{Level: 94, TotalExp: 7944614, NextExp: 8771558},
		{Level: 95, TotalExp: 8771558, NextExp: 9684577},
		{Level: 96, TotalExp: 9684577, NextExp: 10692629},
		{Level: 97, TotalExp: 10692629, NextExp: 11805606},
		{Level: 98, TotalExp: 11805606, NextExp: 13034431},
		{Level: 99, TotalExp: 13034431, NextExp: 0},
	}
}

func calcPriority(task *core.Task) int {
	return task.Short + task.Long
}

func calcToPercent(userExp int, levelTotalExp int, nextExp int) float32 {
	return (float32(userExp - levelTotalExp) / float32(nextExp - levelTotalExp)) * 100
}

func PrioritizeTasks(tasks *[]core.Task) {
	// Sort tasks by priority.
	// Priority = short-term priority value + long-term priority value.
	sort.Slice(*tasks, func(i, j int) bool {
		return calcPriority(&(*tasks)[i]) > calcPriority(&(*tasks)[j])
	})

	// Set ordered priority numbers.
	for i := range *tasks {
		(*tasks)[i].Priority = i + 1
	}
}

func SortTasksByCompletedOn(tasks *[]core.Task) {
	sort.Slice(*tasks, func(i, j int) bool {
		timeA := (*tasks)[i].CompletedOn
		timeB := (*tasks)[j].CompletedOn
		return timeA.After(timeB)
	})
}

func GetExpToNextLevel(level int) int {
	expTable := getExpTable();

	if (level > 0 && len(expTable) >= level) {
		return expTable[level - 1].NextExp
	}

	return 0
}

func CompleteTask(user *core.User, task *core.Task) []core.LevelProgress {
	task.CompletedOn = time.Now()
	task.Priority = 0

	expTable := getExpTable();

	// If we are max level then there is no more progress to be had!
	if user.Level == expTable[len(expTable) - 1].Level {
		return []core.LevelProgress{}
	}

	levelTrack := user.Level
	expTrack := user.TotalExp
	seqTrack := 1

	priority := calcPriority(task)

	// TODO: Lots of room for improvement, but this will do for now.
	user.TotalExp += (priority / 2) * task.Energy * 100

	var lps []core.LevelProgress

	for i := range expTable {
		row := expTable[i]

		if expTrack >= row.TotalExp && expTrack < row.NextExp {
			// First record to set initial exp progress state.
			if (len(lps) == 0) {
				var lp core.LevelProgress
				lp.FromLevel = levelTrack
				lp.ToLevel = levelTrack + 1
				lp.ToPercent = calcToPercent(expTrack, row.TotalExp, row.NextExp)
				lp.Sequence = seqTrack
				lps = append(lps, lp)
				seqTrack++
			}

			if user.TotalExp >= row.TotalExp && user.TotalExp < row.NextExp {
				// Exp gain.
				var lp core.LevelProgress
				lp.FromLevel = levelTrack
				lp.ToLevel = levelTrack + 1
				lp.ToPercent = calcToPercent(user.TotalExp, row.TotalExp, row.NextExp)
				lp.Sequence = seqTrack
				lps = append(lps, lp)
				break
			} else {
				// Level up.
				{
					// Move progress bar to 100%.
					var lp core.LevelProgress
					lp.FromLevel = levelTrack
					lp.ToLevel = levelTrack + 1
					lp.ToPercent = 100
					lp.Sequence = seqTrack
					lps = append(lps, lp)
				}

				seqTrack++

				{
					// Move progress bar back to 0%.
					var lp core.LevelProgress
					lp.FromLevel = levelTrack + 1
					lp.ToLevel = levelTrack + 2
					lp.ToPercent = 0
					lp.Sequence = seqTrack
					lps = append(lps, lp)
				}

				levelTrack++
				expTrack = row.NextExp
				seqTrack++
			}
		}
	}

	user.Level = levelTrack

	return lps
}
