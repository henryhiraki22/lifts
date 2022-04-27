package entity

type LiftA struct {
	ID         int    `gorm:"primary_key"`
	Squat      int    `json:"squat" validate:"required"`
	BenchPress int    `json:"bench_press" validate:"required"`
	BarbellRow int    `json:"barbell_row" validate:"required"`
	Date       string `json:"date" validate:"required"`
}

type LiftB struct {
	ID            int    `gorm:"primary_key"`
	Squat         int    `json:"squat" validate:"required"`
	OverHeadPress int    `json:"over_head_press" validate:"required"`
	DeadLift      int    `json:"dead_lift" validate:"required"`
	Date          string `json:"date" validate:"required"`
}
