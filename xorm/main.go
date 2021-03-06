package main

import (
   // "fmt"
    "strconv"
    "gopkg.in/macaron.v1"
    "github.com/xingcuntian/binding"
)

func main(){
    m:=macaron.Classic()
    m.Use(macaron.Renderer())
    m.Get("/add/:username",AddHandler)
    m.Get("/list",ListHandler)
    m.Get("/info/:id",InfoHandler)
    m.Get("/update/:id/:name",UpdateHandler)
    m.Get("/delete/:id",DeleteHandler)
    m.Get("/contact/create",CreateHandler)
    // m.Post("/contact/submit",binding.Bind(Contact{}),func(contact Contact)string{
    //     return fmt.Sprintf("Name:%s\n Email:%s\n Message:%s\n Mailing Address:%v",
    //     contact.Name,contact.Email,contact.Message,contact.MailingAddress)
    // })
    m.Post("/contact/submit",binding.Bind(Contact{}),AddContactHandler)
    m.Run()
}


func AddContactHandler(contact Contact) string{
    err:=CreateContact(contact)
    if err!=nil{
        return "create contact fail"
    }else{
        return "create contact success"
    }

}

func CreateHandler(ctx *macaron.Context){
    ctx.HTML(200,"create")
}

func AddHandler(ctx *macaron.Context) string{
    name:= ctx.Params(":username")
    if len(name) > 0{
        err:=CreateAccount(name)
        if err!=nil{
           return "create user fail" 
        }
        return "create user success"
    }else{
        return "create user fail"
    }
}

func ListHandler(ctx *macaron.Context){
 list,_:=ListAccount()
 ctx.Data["List"]=list
 ctx.HTML(200,"list")
}

func InfoHandler(ctx *macaron.Context){
    id:=ctx.Params(":id")
    Numid,_:=strconv.ParseInt(id,10,64)
    info,_:=InfoAccount(Numid)
    ctx.Data["Info"]=info
    ctx.HTML(200,"info")
}

func UpdateHandler(ctx *macaron.Context) string{
    id:=ctx.Params(":id")
    Numid,err:=strconv.ParseInt(id,10,64)
    if err!=nil{
        return "id change fail"
    }
    name:=ctx.Params(":name")
    err=UpdateAccount(name,Numid)
    if err!=nil{
        return "update fail"
    }
    return "update success"+name
}

func DeleteHandler(ctx *macaron.Context)string{
    id:=ctx.Params(":id")
    Numid,err:=strconv.ParseInt(id,10,64)
    if err!=nil{
        return "id change fail"
    }
    err=DeleteAccount(Numid)
    if err!=nil{
        return "delete data fail"
    }
    return "delete data success"
}