// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package c1

import (
	"encoding/json"
	"time"
)

type GowebappExercise struct {
	ExerciseID   int64  `db:"exercise_id" json:"exerciseID"`
	ExerciseName string `db:"exercise_name" json:"exerciseName"`
}

type GowebappImage struct {
	ImageID     int64  `db:"image_id" json:"imageID"`
	UserID      int64  `db:"user_id" json:"userID"`
	ContentType string `db:"content_type" json:"contentType"`
	ImageData   []byte `db:"image_data" json:"imageData"`
}

type GowebappSet struct {
	SetID      int64 `db:"set_id" json:"setID"`
	ExerciseID int64 `db:"exercise_id" json:"exerciseID"`
	Weight     int32 `db:"weight" json:"weight"`
}

type GowebappUser struct {
	UserID       int64           `db:"user_id" json:"userID"`
	UserName     string          `db:"user_name" json:"userName"`
	PassWordHash string          `db:"pass_word_hash" json:"passWordHash"`
	Name         string          `db:"name" json:"name"`
	Config       json.RawMessage `db:"config" json:"config"`
	CreatedAt    time.Time       `db:"created_at" json:"createdAt"`
	IsEnabled    bool            `db:"is_enabled" json:"isEnabled"`
}

type GowebappWorkout struct {
	WorkoutID  int64     `db:"workout_id" json:"workoutID"`
	SetID      int64     `db:"set_id" json:"setID"`
	UserID     int64     `db:"user_id" json:"userID"`
	ExerciseID int64     `db:"exercise_id" json:"exerciseID"`
	StartDate  time.Time `db:"start_date" json:"startDate"`
}
