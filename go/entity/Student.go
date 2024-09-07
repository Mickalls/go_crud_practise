package entity

import "fmt"

/*
CREATE TABLE `students` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `age` INT NOT NULL,
  `email` VARCHAR(100),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (Student) TableName() string {
	return "student"
}

func (a Student) Print() {
	fmt.Println("id:", a.ID, " name:", a.Name, " age:", a.Age, " email:", a.Email)
}
