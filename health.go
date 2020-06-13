package health

type Health struct {
	Status  Status                  `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty"`
	Data    *map[string]interface{} `json:"data,omitempty" gorm:"column:data" bson:"data,omitempty" dynamodbav:"data,omitempty" firestore:"data,omitempty"`
	Details *map[string]Health      `json:"details,omitempty" gorm:"column:details" bson:"details,omitempty" dynamodbav:"details,omitempty" firestore:"details,omitempty"`
}
