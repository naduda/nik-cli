package scheduler

import (
	"errors"
	"fmt"
	"github.com/go-co-op/gocron"
	"strconv"
	"strings"
	"time"
)

func Run(every, at string) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(every)
	atTime := startAtTime(at)
	fmt.Println(atTime)
	s.StartAt(atTime)
	_, err := s.Do(func() {
		if err := testJob(); err != nil {
			fmt.Println(err.Error())
		}
	})
	if err != nil {
		panic(err.Error())
	}
	s.StartAsync()
}

func testJob() error {
	fmt.Println("scheduler works")
	if time.Now().Minute() == 31 {
		return errors.New("some error")
	}
	return nil
}

func startAtTime(at string) time.Time {
	atArr := strings.Split(at, ":")
	n := time.Now().Add(1 * time.Second)
	if len(atArr) != 2 {
		return n
	}

	h, err := strconv.Atoi(atArr[0])
	if err != nil || h > 23 {
		return n
	}

	m, err := strconv.Atoi(atArr[1])
	if err != nil || m > 59 {
		return n
	}

	return time.Date(n.Year(), n.Month(), n.Day(), h, m, 0, 0, n.Location())
}
