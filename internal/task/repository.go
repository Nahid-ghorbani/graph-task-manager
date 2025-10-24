package task

import (
	"gorm.io/gorm"
)

// Define methods to do task table's operation
type TaskRepository interface{
	Create(task *Task) error
	Update(task *Task) error
	GetAll(tasks *[]Task) error
	Delete(task *Task) error
	FindTask(task *Task, id any) error
}

type GormRepository struct{
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{DB: db}
}

func (r *GormRepository) Create(task *Task) error {
	if err := r.DB.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) Update(task *Task) error {
	if err := r.DB.Save(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) Delete(task *Task) error {
	if err := r.DB.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) GetAll(tasks *[]Task) error {
	if err := r.DB.Find(&tasks).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) FindTask(task *Task, id any) error {
	if err := r.DB.Find(&task, id).Error; err != nil {
		return err
	}
	return nil
}

