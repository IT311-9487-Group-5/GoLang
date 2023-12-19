package main

import (
	"fmt"
)

// StudentCredentials represents student login credentials.
type StudentCredentials struct {
	Username string
	Password string
}

// Student represents a college student.
type Student struct {
	Name     string
	Address  string
	Contacts int
	Enrolled bool // Add a field to track enrollment status
	Block    int  // Add a field to track the enrolled block index
}

// Course represents a college course.
type Course struct {
	Name     string
	Schedule string
}

// Block represents a time block with a set of courses.
type Block struct {
	Name    string
	Courses []Course
}

// CollegeEnrollmentSystem represents the main enrollment system.
type CollegeEnrollmentSystem struct {
	Blocks   []Block
	Students []Student
}

// NewCourse creates a new course.
func NewCourse(name, schedule string) Course {
	return Course{
		Name:     name,
		Schedule: schedule,
	}
}

// NewBlock creates a new block with a set of courses.
func NewBlock(name string, courses []Course) Block {
	return Block{
		Name:    name,
		Courses: courses,
	}
}

func NewCollegeEnrollmentSystem() CollegeEnrollmentSystem {
	blockA := NewBlock("Block A", []Course{
		NewCourse("CFE 105A", "Monday,Thursday,Friday 9:00 AM - 10:30 AM"),
		NewCourse("IT 111", "Tuesday,Saturday, Wednesday 11:00 AM - 12:30 PM"),
	})

	blockB := NewBlock("Block B", []Course{
		NewCourse("CFE 105A", "Monday,Thursday,Friday 9:00 AM - 10:30 AM"),
		NewCourse("IT 111", "Tuesday,Saturday, Wednesday 11:00 AM - 12:30 PM"),
	})

	// Create two student profiles with different details.
	student1 := Student{
		Name:     "Rodrigo, Juan",
		Address:  "Segundo St., Legarda-Burnham, Baguio City ",
		Contacts: 6393976111874, // Change to an integer
	}

	student2 := Student{
		Name:     "Crisanto, Alejandro",
		Address:  "Upper Bonifacio St., Baguio City ",
		Contacts: 6394865432103, // Change to an integer
	}

	return CollegeEnrollmentSystem{
		Blocks:   []Block{blockA, blockB},
		Students: []Student{student1, student2},
	}
}

func (ces *CollegeEnrollmentSystem) ShowAvailableBlocks() {
	fmt.Println("Available Blocks:")
	for i, block := range ces.Blocks {
		fmt.Printf("%d. %s\n", i+1, block.Name)
	}
}

func main() {
	// Create an array to store student credentials.
	credentials := []StudentCredentials{
		{
			Username: "juan",
			Password: "rodrigo321",
		},
		{
			Username: "alejandro",
			Password: "crisanto321",
		},
	}

	var username, password string

	// Prompt for login credentials until a valid login is provided.
	for {
		fmt.Print("Enter your username: ")
		fmt.Scan(&username)
		fmt.Print("Enter your password: ")
		fmt.Scan(&password)

		// Check if the entered credentials are valid.
		validLogin := false
		var studentIndex int
		for i, cred := range credentials {
			if cred.Username == username && cred.Password == password {
				validLogin = true
				studentIndex = i
				break
			}
		}

		if validLogin {
			// Once logged in, create the enrollment system with the corresponding student information.
			enrollmentSystem := NewCollegeEnrollmentSystem()
			student := &enrollmentSystem.Students[studentIndex]

			// Display the main menu for the logged-in student.
			fmt.Printf("\nWelcome: %s\n", student.Name)

			for {
				fmt.Println("\nMain Menu:")
				fmt.Println("1. View Profile")
				fmt.Println("2. Enroll in a Block")
				fmt.Println("3. View Enrolled Block and Subjects")
				fmt.Println("4. Exit")

				var choice int
				fmt.Print("Enter your choice: ")
				fmt.Scan(&choice)

				switch choice {
				case 1:
					// View Profile
					fmt.Printf("Name: %s\n", student.Name)
					fmt.Printf("Address: %s\n", student.Address)
					fmt.Printf("Contact: %d\n", student.Contacts)

				case 2:
					// Enroll in a Block
					if student.Enrolled {
						fmt.Println("You are already enrolled in a block.")
					} else {
						enrollmentSystem.ShowAvailableBlocks()
						var selectedBlockIndex int
						fmt.Print("Enter the number of the block you want to enroll in: ")
						_, err := fmt.Scan(&selectedBlockIndex)
						if err != nil || selectedBlockIndex < 1 || selectedBlockIndex > len(enrollmentSystem.Blocks) {
							fmt.Println("Invalid input. Please enter a valid block number.")
							continue
						}

						selectedBlock := enrollmentSystem.Blocks[selectedBlockIndex-1]
						student.Block = selectedBlockIndex - 1
						student.Enrolled = true
						fmt.Printf("You have successfully enrolled in %s.\n", selectedBlock.Name)
						fmt.Println("Course Schedule:")
						for _, course := range selectedBlock.Courses {
							fmt.Printf("%s: %s\n", course.Name, course.Schedule)
						}
					}

				case 3:
					// View Enrolled Block and Subjects
					if student.Enrolled {
						enrolledBlock := enrollmentSystem.Blocks[student.Block]
						fmt.Printf("You are enrolled in Block %s.\n", enrolledBlock.Name)
						fmt.Println("Enrolled Courses:")
						for _, course := range enrolledBlock.Courses {
							fmt.Printf("%s: %s\n", course.Name, course.Schedule)
						}
					} else {
						fmt.Println("You haven't enrolled in any block yet.")
					}

				case 4:
					fmt.Println("Goodbye!")
					return // Exit the program.

				default:
					fmt.Println("Invalid choice. Please select a valid option.")
				}
			}
		} else {
			fmt.Println("Invalid login credentials. Please try again.")
		}
	}
}
