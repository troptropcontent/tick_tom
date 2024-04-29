package models

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/troptropcontent/tick_tom/db"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name     string `form:"project[name]"`
	UserID   uint
	User     User
	Sessions []Session `gorm:"polymorphic:Holder;"`
}

// This function returns the total time spent on the project, it takes all the ended session and sums the duration between the started_at and ended_at
func (p Project) TotalTimeSpent() time.Duration {
	row := db.DB.Table("sessions").Where("holder_id = ? AND holder_type = ? AND ended_at > ?", p.ID, "projects", time.Time{}).Select("SUM(ended_at - started_at) as total_time").Row()
	if err := row.Err(); err != nil {
		panic(err)
	}

	rawDuration := ""
	row.Scan(&rawDuration)

	fmt.Println("rawDuration: ", rawDuration)

	if rawDuration == "" {
		return 0
	}

	var myRegexp = regexp.MustCompile(`(?P<hours>\d+):(?P<minutes>\d+):(?P<seconds>\d+)`)
	if !myRegexp.Match([]byte(rawDuration)) {
		panic("rawDuration is not a valid duration")
	}

	match := myRegexp.FindStringSubmatch(rawDuration)
	result := make(map[string]string)
	for i, name := range myRegexp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	hours, _ := strconv.Atoi(result["hours"])

	minutes, _ := strconv.Atoi(result["minutes"])

	seconds, _ := strconv.Atoi(result["seconds"])

	duration, err := time.ParseDuration(fmt.Sprintf("%dh%dm%ds", hours, minutes, seconds))
	if err != nil {
		panic("Error: " + err.Error())
	}

	return duration
}
