package cnfg

import "time"

// https://pkg.go.dev/flag#Value
type DateArg struct {
	Date      *time.Time
	IsDateSet *bool
}

func (v DateArg) String() string {
	if v.Date == nil {
		return ""
	}
	return v.Date.Format("2006-01-02")
}

func (v DateArg) Set(s string) error {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(time.DateOnly, s, loc)
	if err != nil {
		return err
	}
	*v.Date = t
	*v.IsDateSet = true
	return nil
}

type DateArg2 struct {
	Time time.Time
}

func (v *DateArg2) String() string {
	return v.Time.Format("02.01.2006")
}

func (v *DateArg2) Set(s string) error {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(time.DateOnly, s, loc)
	if err != nil {
		return err
	}
	v.Time = t
	return nil
}

type DateArg3 struct {
	Time *time.Time
}

func (v *DateArg3) String() string {
	if v.Time == nil {
		return ""
	}
	return v.Time.Format("02.01.2006")
}

func (v *DateArg3) Set(s string) error {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(time.DateOnly, s, loc)
	if err != nil {
		return err
	}
	v.Time = &t
	return nil
}
