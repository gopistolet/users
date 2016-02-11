package users

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSaveLoad(t *testing.T) {

	Convey("UserStore Load/Save/Construct", t, func() {

		s, err := NewUserStore("users.json")
		sActual := UserStore{Users: map[string]User{"DenBeke": User{Name: "DenBeke", Mail: "denbeke@test.com"}}}
		So(err, ShouldEqual, nil)
		So(s.Users, ShouldResemble, sActual.Users)

		s.Users = map[string]User{"DenBeke": User{Name: "Hello world!", Mail: "denbeke@test.com"}}
		sActual = UserStore{Users: map[string]User{"DenBeke": User{Name: "Hello world!", Mail: "denbeke@test.com"}}}
		So(s.Users, ShouldResemble, sActual.Users)

		So(s.Users, ShouldResemble, sActual.Users)
		err = s.Save()
		So(err, ShouldEqual, nil)
		err = s.Load()
		So(err, ShouldEqual, nil)

		s.Users = map[string]User{"DenBeke": User{Name: "DenBeke", Mail: "denbeke@test.com"}}
		err = s.Save()
		So(err, ShouldEqual, nil)

	})

}

func TestAddGet(t *testing.T) {

	Convey("UserStore Get/Add/Exists", t, func() {

		s := UserStore{filename: "users.json", Users: map[string]User{"DenBeke": User{Name: "DenBeke", Mail: "denbeke@test.com"}}}
		So(s.Exists("DenBeke"), ShouldEqual, true)
		So(s.Exists("Some random name"), ShouldEqual, false)

		So(*s.Get("DenBeke"), ShouldResemble, User{Name: "DenBeke", Mail: "denbeke@test.com"})

		err := s.Add(User{Name: "GoPistolet", Mail: "gopistolet@test.com"})
		So(err, ShouldEqual, nil)
		So(s.Exists("GoPistolet"), ShouldEqual, true)
		So(*s.Get("GoPistolet"), ShouldResemble, User{Name: "GoPistolet", Mail: "gopistolet@test.com"})

		err = s.Delete("GoPistolet")
		So(err, ShouldEqual, nil)
		So(s.Exists("GoPistolet"), ShouldNotEqual, true)
		So(len(s.Users), ShouldEqual, 1)

	})

}
