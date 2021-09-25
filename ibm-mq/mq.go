package ibmmq

import "github.com/ibm-messaging/mq-golang/v5/ibmmq"

type QueueConfig struct {
	ManagerName    string      `mapstructure:"manager_name" json:"managerName,omitempty" gorm:"column:managername" bson:"managerName,omitempty" dynamodbav:"managerName,omitempty" firestore:"managerName,omitempty"`
	ChannelName    string      `mapstructure:"channel_name" json:"channelName,omitempty" gorm:"column:channelname" bson:"channelName,omitempty" dynamodbav:"channelName,omitempty" firestore:"channelName,omitempty"`
	ConnectionName string      `mapstructure:"connection_name" json:"connectionName,omitempty" gorm:"column:connectionname" bson:"connectionName,omitempty" dynamodbav:"connectionName,omitempty" firestore:"connectionName,omitempty"`
	QueueName      string      `mapstructure:"queue_name" json:"queueName,omitempty" gorm:"column:queuename" bson:"queueName,omitempty" dynamodbav:"queueName,omitempty" firestore:"queueName,omitempty"`
	Retry          RetryConfig `mapstructure:"retry" json:"retry,omitempty" gorm:"column:retry" bson:"retry,omitempty" dynamodbav:"retry,omitempty" firestore:"retry,omitempty"`
}

type MQAuth struct {
	UserId   string `mapstructure:"user_id" json:"userId,omitempty" gorm:"column:userid" bson:"userId,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty"`
	Password string `mapstructure:"password" json:"password,omitempty" gorm:"column:password" bson:"password,omitempty" dynamodbav:"password,omitempty" firestore:"password,omitempty"`
}

type RetryConfig struct {
	Retry1 int64 `mapstructure:"1" json:"retry1,omitempty" gorm:"column:retry1" bson:"retry1,omitempty" dynamodbav:"retry1,omitempty" firestore:"retry1,omitempty"`
	Retry2 int64 `mapstructure:"2" json:"retry2,omitempty" gorm:"column:retry2" bson:"retry2,omitempty" dynamodbav:"retry2,omitempty" firestore:"retry2,omitempty"`
	Retry3 int64 `mapstructure:"3" json:"retry3,omitempty" gorm:"column:retry3" bson:"retry3,omitempty" dynamodbav:"retry3,omitempty" firestore:"retry3,omitempty"`
	Retry4 int64 `mapstructure:"4" json:"retry4,omitempty" gorm:"column:retry4" bson:"retry4,omitempty" dynamodbav:"retry4,omitempty" firestore:"retry4,omitempty"`
	Retry5 int64 `mapstructure:"5" json:"retry5,omitempty" gorm:"column:retry5" bson:"retry5,omitempty" dynamodbav:"retry5,omitempty" firestore:"retry5,omitempty"`
	Retry6 int64 `mapstructure:"6" json:"retry6,omitempty" gorm:"column:retry6" bson:"retry6,omitempty" dynamodbav:"retry6,omitempty" firestore:"retry6,omitempty"`
	Retry7 int64 `mapstructure:"7" json:"retry7,omitempty" gorm:"column:retry7" bson:"retry7,omitempty" dynamodbav:"retry7,omitempty" firestore:"retry7,omitempty"`
	Retry8 int64 `mapstructure:"8" json:"retry8,omitempty" gorm:"column:retry8" bson:"retry8,omitempty" dynamodbav:"retry8,omitempty" firestore:"retry8,omitempty"`
	Retry9 int64 `mapstructure:"9" json:"retry9,omitempty" gorm:"column:retry9" bson:"retry9,omitempty" dynamodbav:"retry9,omitempty" firestore:"retry9,omitempty"`
}

func newQueueManagerByConfig(c QueueConfig, auth MQAuth) (*ibmmq.MQQueueManager, error) {
	cd := NewMQCDByChannelAndConnection(c.ChannelName, c.ConnectionName)
	csp := NewMQCSPByConfig(auth)

	cno := ibmmq.NewMQCNO()
	cno.ClientConn = cd
	cno.Options = ibmmq.MQCNO_CLIENT_BINDING + ibmmq.MQCNO_RECONNECT + ibmmq.MQCNO_HANDLE_SHARE_BLOCK
	cno.SecurityParms = csp

	mgr, err := ibmmq.Connx(c.ManagerName, cno)
	return &mgr, err
}
func NewMQCDByChannelAndConnection(channelName string, connectionName string) *ibmmq.MQCD {
	cd := ibmmq.NewMQCD()
	cd.ChannelName = channelName
	cd.ConnectionName = connectionName
	return cd
}
func NewMQCSPByConfig(auth MQAuth) *ibmmq.MQCSP {
	csp := ibmmq.NewMQCSP()
	csp.AuthenticationType = ibmmq.MQCSP_AUTH_USER_ID_AND_PWD
	csp.UserId = auth.UserId
	csp.Password = auth.Password
	return csp
}
