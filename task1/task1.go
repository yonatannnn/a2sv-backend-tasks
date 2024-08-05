package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name     string
	Subjects []string
	Grades   []int
}

func (s *Student) addSubject(subject string) {
	s.Subjects = append(s.Subjects, strings.TrimSpace(subject))
}

func (s *Student) addGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

func (s *Student) calculateAverage() float64 {
	var total int
	for _, grade := range s.Grades {
		total += grade
	}
	return float64(total) / float64(len(s.Grades))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	student := Student{Name: name}
	subjects := 0
	for {
		fmt.Println("How many subjects do you have?")
	subjectsStr, _ := reader.ReadString('\n')
	subjectsStr = strings.TrimSpace(subjectsStr)
	sub, err := strconv.Atoi(subjectsStr)
	if err != nil || sub < 1 {
		fmt.Println("Please enter a valid number of subjects")
		continue
	}
	subjects = sub
	break
}
	for i := 0; i < subjects; i++ {

		for {
			fmt.Print("Enter the Subject: ")
			subject, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please enter a valid subject")
			continue
		}
		subject = strings.TrimSpace(subject)
		student.addSubject(subject)
		break
	}
		
		

		for {
			fmt.Print("Enter the Grade: ")
			gradeStr, _ := reader.ReadString('\n')
			gradeStr = strings.TrimSpace(gradeStr)
			intValue, err := strconv.Atoi(gradeStr)
			if err != nil || intValue < 0 || intValue > 100 {
				fmt.Println("Please enter a valid grade (0-100)")
				continue
			}
			student.addGrade(intValue)
			break
		}
	}

	fmt.Printf("Student: %s\n", student.Name)
	for i, subject := range student.Subjects {
		fmt.Printf("Subject: %s, Grade: %d\n", subject, student.Grades[i])
	}
	fmt.Printf("Average: %.2f\n", student.calculateAverage())
}
