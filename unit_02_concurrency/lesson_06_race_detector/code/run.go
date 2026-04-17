package lesson_06_race_detector

import "fmt"

func Run() {
	fmt.Println("\n=== Lesson 6: Race Detector 数据竞争检测实战 ===")

	DemoDataRace()
	DemoRaceDetector()
	DemoFixingRaces()
}
