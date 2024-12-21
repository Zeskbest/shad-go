//go:build !solution

package hotelbusiness

import (
	"slices"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func AddGuest(ll []*Load, StartDate, EndDate int) []*Load {
	// foundFirst - whether to insert a new object
	var foundFirst, foundLast bool
	for _, l := range ll {
		if StartDate == l.StartDate {
			foundFirst = true
		}
		if l.StartDate == EndDate {
			foundLast = true
		}
		if StartDate <= l.StartDate {
			// increase amounts on suitable days
			if l.StartDate < EndDate {
				l.GuestCount += 1
			} else {
				break
			}
		}
	}
	if !foundFirst {
		firstIndex := 0
		for i := 0; i < len(ll); i++ {
			if StartDate < ll[i].StartDate {
				firstIndex = i
				break
			}
			if i == len(ll)-1 {
				firstIndex = len(ll)
			}
		}
		count := 0
		if firstIndex != 0 {
			count = ll[firstIndex-1].GuestCount
		}
		ll = slices.Insert(ll, firstIndex, &Load{
			StartDate:  StartDate,
			GuestCount: count + 1,
		})
	}
	if !foundLast {
		lastIndex := 0
		for i := 0; i < len(ll); i++ {
			if EndDate < ll[i].StartDate {
				lastIndex = i
				break
			}
			if i == len(ll)-1 {
				lastIndex = len(ll)
			}
		}
		count := 1
		if lastIndex != 0 {
			count = ll[lastIndex-1].GuestCount
		}
		ll = slices.Insert(ll, lastIndex, &Load{
			StartDate:  EndDate,
			GuestCount: count - 1,
		})
	}
	return ll
}

func Reduce(ll []*Load) []*Load {
	for i := 1; ; i++ {
		if ll[i-1].GuestCount == ll[i].GuestCount {
			ll = append(ll[:i], ll[i+1:]...)
			i-- // len reduced - process this idx one more time
		}
		if i == len(ll)-1 {
			break
		}
	}
	return ll
}

func ComputeLoad(guests []Guest) []Load {
	if len(guests) == 0 {
		return []Load{}
	}

	var lastDay int
	for _, g := range guests {
		if lastDay < g.CheckOutDate {
			lastDay = g.CheckOutDate
		}
	}
	ll := make([]*Load, 0, len(guests)*2)
	for _, g := range guests {
		ll = AddGuest(ll, g.CheckInDate, g.CheckOutDate)
		ll = Reduce(ll)
	}

	res := make([]Load, 0, len(ll))
	for _, l := range ll {
		res = append(res, *l)
	}
	return res
}
