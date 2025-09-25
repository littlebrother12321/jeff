package controllers

import (
	"fmt"
	"firstbee/models"
	"reflect"
	"strconv"
)

type AboutListController struct {
	BaseAdminController
}

type AboutDetailController struct {
	BaseAdminController
}

// READ HELPERS
type Item struct {
	Key   string
	Value any // Using 'any' to allow any type for Value
}

// Function to convert struct to ordered data based on struct field order
func structToOrderedData(v any) []Item {
	var orderedData []Item
	// Use reflection to get the value and type of the struct
	val := reflect.ValueOf(v)
	typ := val.Type()
	// Iterate over the struct fields in their defined order
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		orderedData = append(orderedData, Item{Key: fieldType.Name, Value: field.Interface()})
	}
	return orderedData
}

func read(c *BaseAdminController, createModel func(id uint64) any) any {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Invalid ID"}
		c.ServeJSON()
		return nil
	}

	ptr := createModel(id)
	err = models.O.Read(ptr)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "Not Found"}
		c.ServeJSON()
		return nil
	}

	return ptr
}

// List
func (c *AboutListController) Get() {
	c.RequireAuth()

	c.Data["Title"] = "About"
	c.TplName = "admin/list.tpl"
	q := models.O.QueryTable("about_model")
	var contacts []*models.AboutModel
	q.All(&contacts)
	var l [][]Item
	for _, contact := range contacts {
		l = append(l, structToOrderedData(*contact))
	}
	c.Data["List"] = l
	c.Data["BaseHref"] = "/admin/about"
}

// Create (NOT IMPLEMENTED)
// func (c *ContactListController) Post() {
// 	c.RequireAuth()

// 	c.Data["Title"] = "Contacts"
// 	c.TplName = "admin/create.tpl"
// }

// Read
func (c *AboutDetailController) Get() {
	c.RequireAuth()

	c.Data["Title"] = "About"
	c.TplName = "admin/read.tpl"

	contact := read(&c.BaseAdminController, func(id uint64) any {
		contact := &models.AboutModel{Id: id}
		return (any)(contact)
	})

	if m, ok := contact.(*models.AboutModel); ok {
		val := structToOrderedData(*m)
		c.Data["Title"] = fmt.Sprintf("Contact %d", m.Id)
		c.Data["Item"] = val
		c.Data["Id"] = m.Id
		c.Data["BaseHref"] = "/admin/about"
	}
}

// Update (NOT IMPLEMENTED)
// func (c *ContactDetailController) Post() {
// 	c.RequireAuth()

// 	c.Data["Title"] = "Contacts"
// 	c.TplName = "admin/update.tpl"
// }

// Delete
func (c *AboutDetailController) Delete() {
	c.RequireAuth()

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	contact := &models.AboutModel{Id: id}
	_, err := models.O.Delete(contact)
	fmt.Println(err)
	c.Data["Title"] = "About"
}
