package storage

//interfaces will be used as will be set up as the plugin 
//also used for the testing for the databsase

type Storage interface{
	CreateStudent(name string,email string, age int ) (int64,error)
}