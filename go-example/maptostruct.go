package main
//map映射到结构体，这里只支持简单的数据类型，复杂的需要在拓展
import (
    "errors"
    "fmt"
    "reflect"
    "strconv"
    "time"
)

type User struct{
    Name string
    Age int64
    Date time.Time
}

func main(){
    data:=make(map[string]interface{})
    data["Name"] = "你大爷"
    data["Age"] = 100
    data["Date"] ="2016-06-12 00:00:00"
    result:=&User{}
    err:=FillStruct(data,result)
    fmt.Println(err,fmt.Sprintf("%+v",*result))
}

//用map填充结构体
func FillStruct(data map[string]interface{},obj interface{})error{
    for k,v:=range data{
        err:=SetField(obj,k,v)
        if err!=nil{
            return err
        }
    }
    return nil
}

//用map的值替换结构体
func SetField(obj interface{}, name string, value interface{}) error {
    structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
    structFieldValue := structValue.FieldByName(name) //结构体单个属性值
    if !structFieldValue.IsValid() {
        return fmt.Errorf("No such field: %s in obj", name)
    }
    if !structFieldValue.CanSet() {
        return fmt.Errorf("Cannot set %s field value", name)
    }
    structFieldType := structFieldValue.Type() //结构体的类型
    val := reflect.ValueOf(value)              //map值的反射值
    var err error
    if structFieldType != val.Type() {
        val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
        if err != nil {
            return err
        }
    }
    structFieldValue.Set(val)
    return nil
}

//类型转换
func TypeConversion(value string,ntype string)(reflect.Value,error){
    if ntype == "string"{
        return reflect.ValueOf(value),nil
    }else if ntype == "time.Time"{
        
        t,err:=time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
        return reflect.ValueOf(t),err
    }else if ntype == "Time"{
        t,err:=time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
        return reflect.ValueOf(t),err
    }else if ntype == "int"{
        i,err:=strconv.Atoi(value)
        return reflect.ValueOf(i),err
    }else if ntype =="int64"{
        i,err:=strconv.ParseInt(value,10,64)
        return reflect.ValueOf(i),err
    }else if ntype =="float64"{
        i,err:=strconv.ParseFloat(value,64)
        return reflect.ValueOf(i),err
    }
    return reflect.ValueOf(value),errors.New("未知错误:"+ntype)
}