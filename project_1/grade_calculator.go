package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name string
	subjects map[string]float64
}

func getInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)

}

func getNumberOfSubjects(reader *bufio.Reader) int {
	for {
		input := getInput(reader, "Enter the number of subjects: ")
		num, err := strconv.Atoi(input)
		if err == nil && num > 0 {
			return num
		}
		fmt.Println("Please enter a valid positive number.")
	}
}

func (s *Student) calculateAverage() float64 {
	if len(s.subjects) == 0 {
		return 0
	}

	var sum float64
	for _,grade := range s.subjects {
		sum += grade
	}
	return sum / float64(len(s.subjects))
}

func (s *Student) displayResults() {
    fmt.Printf("\n=== Results for %s ===\n", s.name)
    fmt.Println("\nSubject Grades:")
    for subject, grade := range s.subjects {
        fmt.Printf("%-20s: %.2f\n", subject, grade)
    }
    fmt.Printf("\nAverage Grade: %.2f\n", s.calculateAverage())
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Student Grade Calculator!")

	student := Student{
		subjects: make(map[string]float64),
	}

	student.name = getInput(reader, "Enter student name: ")
	numSubjects := getNumberOfSubjects(reader)
    
    for i := 0; i < numSubjects; i++ {
        fmt.Printf("\nSubject %d:\n", i+1)
        
        subjectName := getInput(reader, "Enter subject name: ")
        
        for {
            gradeStr := getInput(reader, "Enter grade (0-100): ")
            grade, err := strconv.ParseFloat(gradeStr, 64)
            if err == nil && grade >= 0 && grade <= 100 {
                student.subjects[subjectName] = grade
                break
            }
            fmt.Println("Please enter a valid grade between 0 and 100")
        }
    }

	student.displayResults()
}