package service

import (
	"CRUD/go/dao"
	"CRUD/go/entity"
	"fmt"
)

// CreatStudent 新建一个学生数据
func CreateStudent(stu *entity.Student) (err error) {
	if err = dao.SqlConn.Create(stu).Error; err != nil {
		return err
	}
	return
}

// GetAllStudent 获取所有学生的信息
func GetAllStudent() (stu []entity.Student, err error) {
	if err = dao.SqlConn.Find(&stu).Error; err != nil {
		return stu, err
	}
	return stu, nil
}

// GetStudentByID 根据id查学生信息
func GetStudentByID(id string) (*entity.Student, error) {
	stu := &entity.Student{} //必须在这里初始化
	if err := dao.SqlConn.Model(&entity.Student{}).Where("id=?", id).First(stu).Error; err != nil {
		return stu, err
	}
	return stu, nil
}

// DeleteStudentByID 根据id删除学生信息
func DeleteStudentByID(id string) (err error) {
	result := dao.SqlConn.Delete(&entity.Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("删除失败，没有查询到 id=%s 的学生信息", id)
	}
	return nil
}

// UpdateStudentByID 根据学生ID更新ID对应学生的信息
func UpdateStudentByID(id string, stu *entity.Student) (err error) {
	result := dao.SqlConn.Model(&entity.Student{}).Where("id=?", id).Updates(stu)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("更新失败，没有查询到 id=%s 的学生信息", id)
	}
	return nil
}
