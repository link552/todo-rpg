package game

import (
	"time"
	"sort"
	"todorpg/internal/core"
)

type levelProgression struct {
	Level int
	TotalExp int
	ExpDiff int
}

func getExpTable() []levelProgression  {
	return []levelProgression{
		{Level: 1, TotalExp: 0, ExpDiff: 0},
		{Level: 2, TotalExp: 83, ExpDiff: 83},
		{Level: 3, TotalExp: 174, ExpDiff: 91},
		{Level: 4, TotalExp: 276, ExpDiff: 102},
		{Level: 5, TotalExp: 388, ExpDiff: 112},
		{Level: 6, TotalExp: 512, ExpDiff: 124},
		{Level: 7, TotalExp: 650, ExpDiff: 138},
		{Level: 8, TotalExp: 801, ExpDiff: 151},
		{Level: 9, TotalExp: 969, ExpDiff: 168},
		{Level: 10, TotalExp: 1154, ExpDiff: 185},
		{Level: 11, TotalExp: 1358, ExpDiff: 204},
		{Level: 12, TotalExp: 1584, ExpDiff: 226},
		{Level: 13, TotalExp: 1833, ExpDiff: 249},
		{Level: 14, TotalExp: 2107, ExpDiff: 274},
		{Level: 15, TotalExp: 2411, ExpDiff: 304},
		{Level: 16, TotalExp: 2746, ExpDiff: 335},
		{Level: 17, TotalExp: 3115, ExpDiff: 369},
		{Level: 18, TotalExp: 3523, ExpDiff: 408},
		{Level: 19, TotalExp: 3973, ExpDiff: 450},
		{Level: 20, TotalExp: 4470, ExpDiff: 497},
		{Level: 21, TotalExp: 5018, ExpDiff: 548},
		{Level: 22, TotalExp: 5624, ExpDiff: 606},
		{Level: 23, TotalExp: 6291, ExpDiff: 667},
		{Level: 24, TotalExp: 7028, ExpDiff: 737},
		{Level: 25, TotalExp: 7842, ExpDiff: 814},
		{Level: 26, TotalExp: 8740, ExpDiff: 898},
		{Level: 27, TotalExp: 9730, ExpDiff: 990},
		{Level: 28, TotalExp: 10824, ExpDiff: 1094},
		{Level: 29, TotalExp: 12031, ExpDiff: 1207},
		{Level: 30, TotalExp: 13363, ExpDiff: 1332},
		{Level: 31, TotalExp: 14833, ExpDiff: 1470},
		{Level: 32, TotalExp: 16456, ExpDiff: 1623},
		{Level: 33, TotalExp: 18247, ExpDiff: 1791},
		{Level: 34, TotalExp: 20224, ExpDiff: 1977},
		{Level: 35, TotalExp: 22406, ExpDiff: 2182},
		{Level: 36, TotalExp: 24815, ExpDiff: 2409},
		{Level: 37, TotalExp: 27473, ExpDiff: 2658},
		{Level: 38, TotalExp: 30408, ExpDiff: 2935},
		{Level: 39, TotalExp: 33648, ExpDiff: 3240},
		{Level: 40, TotalExp: 37224, ExpDiff: 3576},
		{Level: 41, TotalExp: 41171, ExpDiff: 3947},
		{Level: 42, TotalExp: 45529, ExpDiff: 4358},
		{Level: 43, TotalExp: 50339, ExpDiff: 4810},
		{Level: 44, TotalExp: 55649, ExpDiff: 5310},
		{Level: 45, TotalExp: 61512, ExpDiff: 5863},
		{Level: 46, TotalExp: 67983, ExpDiff: 6471},
		{Level: 47, TotalExp: 75127, ExpDiff: 7144},
		{Level: 48, TotalExp: 83014, ExpDiff: 7887},
		{Level: 49, TotalExp: 91721, ExpDiff: 8707},
		{Level: 50, TotalExp: 101333, ExpDiff: 9612},
		{Level: 51, TotalExp: 111945, ExpDiff: 10612},
		{Level: 52, TotalExp: 123660, ExpDiff: 11715},
		{Level: 53, TotalExp: 136594, ExpDiff: 12934},
		{Level: 54, TotalExp: 150872, ExpDiff: 14278},
		{Level: 55, TotalExp: 166636, ExpDiff: 15764},
		{Level: 56, TotalExp: 184040, ExpDiff: 17404},
		{Level: 57, TotalExp: 203254, ExpDiff: 19214},
		{Level: 58, TotalExp: 224466, ExpDiff: 21212},
		{Level: 59, TotalExp: 247886, ExpDiff: 23420},
		{Level: 60, TotalExp: 273742, ExpDiff: 25856},
		{Level: 61, TotalExp: 302288, ExpDiff: 28546},
		{Level: 62, TotalExp: 333804, ExpDiff: 31516},
		{Level: 63, TotalExp: 368599, ExpDiff: 34795},
		{Level: 64, TotalExp: 407015, ExpDiff: 38416},
		{Level: 65, TotalExp: 449428, ExpDiff: 42413},
		{Level: 66, TotalExp: 496254, ExpDiff: 46826},
		{Level: 67, TotalExp: 547953, ExpDiff: 51699},
		{Level: 68, TotalExp: 605032, ExpDiff: 57079},
		{Level: 69, TotalExp: 668051, ExpDiff: 63019},
		{Level: 70, TotalExp: 737627, ExpDiff: 69576},
		{Level: 71, TotalExp: 814445, ExpDiff: 76818},
		{Level: 72, TotalExp: 899257, ExpDiff: 84812},
		{Level: 73, TotalExp: 992895, ExpDiff: 93638},
		{Level: 74, TotalExp: 1096278, ExpDiff: 103383},
		{Level: 75, TotalExp: 1210421, ExpDiff: 114143},
		{Level: 76, TotalExp: 1336443, ExpDiff: 126022},
		{Level: 77, TotalExp: 1475581, ExpDiff: 139138},
		{Level: 78, TotalExp: 1629200, ExpDiff: 153619},
		{Level: 79, TotalExp: 1798808, ExpDiff: 169608},
		{Level: 80, TotalExp: 1986068, ExpDiff: 187260},
		{Level: 81, TotalExp: 2192818, ExpDiff: 206750},
		{Level: 82, TotalExp: 2421087, ExpDiff: 228269},
		{Level: 83, TotalExp: 2673114, ExpDiff: 252027},
		{Level: 84, TotalExp: 2951373, ExpDiff: 278259},
		{Level: 85, TotalExp: 3258594, ExpDiff: 307221},
		{Level: 86, TotalExp: 3597792, ExpDiff: 339198},
		{Level: 87, TotalExp: 3972294, ExpDiff: 374502},
		{Level: 88, TotalExp: 4385776, ExpDiff: 413482},
		{Level: 89, TotalExp: 4842295, ExpDiff: 456519},
		{Level: 90, TotalExp: 5346332, ExpDiff: 504037},
		{Level: 91, TotalExp: 5902831, ExpDiff: 556499},
		{Level: 92, TotalExp: 6517253, ExpDiff: 614422},
		{Level: 93, TotalExp: 7195629, ExpDiff: 678376},
		{Level: 94, TotalExp: 7944614, ExpDiff: 748985},
		{Level: 95, TotalExp: 8771558, ExpDiff: 826944},
		{Level: 96, TotalExp: 9684577, ExpDiff: 913019},
		{Level: 97, TotalExp: 10692629, ExpDiff: 1008052},
		{Level: 98, TotalExp: 11805606, ExpDiff: 1112977},
		{Level: 99, TotalExp: 13034431, ExpDiff: 1228825},
	}
}

func calcPriority(task *core.Task) int {
	return task.Short + task.Long
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

func CompleteTask(user *core.User, task *core.Task) {
	task.CompletedOn = time.Now()

	priority := calcPriority(task)

	// TODO: Lots of room for improvement, but this will do for now.
	user.TotalExp += priority * 100

	expTable := getExpTable();

	// Check for level 99.
	if (user.TotalExp >= expTable[98].TotalExp) {
		user.Level = 99
		return
	}

	// Check for level-up!
	for i := range expTable {
		lp1 := expTable[i]

		// NOTE: Guard check above prevents index out of bounds.
		lp2 := expTable[i + 1]

		if (user.TotalExp >= lp1.TotalExp && user.TotalExp < lp2.TotalExp) {
			user.Level = lp1.Level
			break
		}
	}
}
