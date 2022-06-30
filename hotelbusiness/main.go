package main

import (
	"fmt"
	"sort"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func main() {
	loads := ComputeLoad([]Guest{
		{4, 7},
		{2, 4},
		{2, 3},
	})

	for _, i := range loads {
		fmt.Println(i.StartDate, " ", i.GuestCount)
	}
}

func ComputeLoad(guests []Guest) []Load {
	type CheckDate struct {
		date  int
		delta int
	}
	var dates = make([]CheckDate, len(guests)*2)
	for index, i := range guests {
		dates[2*index] = CheckDate{i.CheckInDate, 1}
		dates[2*index+1] = CheckDate{i.CheckOutDate, -1}
	}
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].date < dates[j].date
	})
	var loads = make([]Load, 0)
	guestCount := 0
	for _, i := range dates {
		guestCount += i.delta
		if len(loads) == 0 {
			loads = append(loads, Load{i.date, guestCount})
			continue
		}
		if loads[len(loads)-1].StartDate == i.date {
			loads[len(loads)-1].GuestCount = guestCount
			continue
		}
		if len(loads) > 1 && loads[len(loads)-1].GuestCount == loads[len(loads)-2].GuestCount {
			loads[len(loads)-1] = Load{i.date, guestCount}
			continue
		}
		loads = append(loads, Load{i.date, guestCount})
	}
	return loads
}
