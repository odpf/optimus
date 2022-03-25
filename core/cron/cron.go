package cron

import (
	"time"

	roboCron "github.com/robfig/cron/v3"
)

type ScheduleSpec struct {
	schd roboCron.Schedule
}

// Next accepts the time and returns the next run time that should
// be used for execution
func (s *ScheduleSpec) Next(t time.Time) time.Time {
	return s.schd.Next(t)
}

// ParseCronSchedule can parse standard cron notation
// it returns a new crontab schedule representing the given
// standardSpec (https://en.wikipedia.org/wiki/Cron). It requires 5 entries
// representing: minute, hour, day of month, month and day of week, in that
// order. It returns a descriptive error if the spec is not valid.
//
// It accepts
//   - Standard crontab specs, e.g. "* * * * ?"
//   - Descriptors, e.g. "@midnight", "@every 1h30m"
func ParseCronSchedule(interval string) (*ScheduleSpec, error) {
	roboCronSchedule, err := roboCron.ParseStandard(interval)
	if err != nil {
		return nil, err
	}

	return &ScheduleSpec{
		schd: roboCronSchedule,
	}, nil
}

// Interval accepts the time and returns
func (s *ScheduleSpec) Interval(t time.Time) time.Duration {
	start := s.Next(t)
	next := s.Next(start)
	return next.Sub(start)
}
