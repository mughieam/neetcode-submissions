/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })
	
	end := math.MinInt
	for _, val := range intervals {
		if end > val.start {
			return false
		}
		end = max(end, val.end)
	}
	return true
}
