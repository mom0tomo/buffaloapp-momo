package actions

import "github.com/mom0tomo/myapp/models"

func (as *ActionSuite) Test_WidgetsResource_List() {
	as.LoadFixture("list widgets")

	res := as.HTML("/widgets").Get()
	as.Equal(200, res.Code)

	description := res.Body.String()

	var widgets models.Widgets
	as.NoError(as.DB.All(&widgets))
	for _, w := range widgets {
		as.Contains(description, w.Name)
	}
}

func (as *ActionSuite) Test_WidgetsResource_Show() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_New() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Create() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Edit() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Update() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Destroy() {
	as.Fail("Not Implemented!")
}
