package database__test

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/CinematicCow/Lumora/internal/database"
	"github.com/CinematicCow/Lumora/internal/models"
	. "github.com/smartystreets/goconvey/convey"
)

func createDBFile(t *testing.T) *os.File {
	f, err := os.CreateTemp("../../../tmp/", "testdb")
	if err != nil {
		t.Fatal("Error while creating test db: ", err)
	}
	return f
}

func openDBFile(t *testing.T) *os.File {
	f := createDBFile(t)

	d, err := os.Open(f.Name())
	if err != nil {
		t.Fatal("Error while opening test db: ", err)
	}
	return d

}

func TestDBOperations(t *testing.T) {
	Convey("Given a db file", t, func() {
		Convey("when reading from an empty db", func() {
			db := openDBFile(t)
			defer db.Close()

			Convey("it should return an error and an empty slice", func() {
				d, err := database.ReadFromDB(db)
				log.Default().Println(err)
				So(err, ShouldBeNil)
				So(d, ShouldBeEmpty)
			})
		})

		Convey("when writing to the db", func() {
			db := createDBFile(t)
			defer db.Close()

			Convey("it should write data without errors", func() {
				md := &models.Data{
					Key:   []byte("test-key"),
					Value: []byte("test-value"),
				}

				err := database.WriteToDB(db, md)
				So(err, ShouldBeNil)
			})
		})

		Convey("when reading from the db", func() {
			db := openDBFile(t)
			defer db.Close()

			Convey("it should return the data without errors", func() {
				d, err := database.ReadFromDB(db)
				So(err, ShouldBeNil)

				Convey("each element should match the written data", func() {
					for _, i := range d {
						So(i, ShouldHaveSameTypeAs, models.Data{})
					}

				})
			})
		})

		Convey("when comparing the written and read data", func() {
			db := createDBFile(t)
			defer db.Close()

			Convey("it should match the written data", func() {
				md := &models.Data{
					Key:   []byte("Jon"),
					Value: []byte("Doe"),
				}

				err := database.WriteToDB(db, md)
				So(err, ShouldBeNil)

				_, err = db.Seek(0, 0)
				So(err, ShouldBeNil)

				d, err := database.ReadFromDB(db)
				So(err, ShouldBeNil)

				So(len(d), ShouldEqual, 1)
				So(d[0].Key, ShouldResemble, md.Key)
				So(d[0].Value, ShouldResemble, md.Value)
			})
		})

		Convey("when writing and reading multiple data to/from the db", func() {
			db := createDBFile(t)
			defer db.Close()

			Convey("it should write and read multiple data without errors", func() {

				s := 10

				for i := 0; i < s; i++ {
					md := &models.Data{
						Key:   []byte("test-key" + strconv.Itoa(i)),
						Value: []byte("test-value + " + strconv.Itoa(i)),
					}
					err := database.WriteToDB(db, md)
					So(err, ShouldBeNil)
				}

				_, err := db.Seek(0, 0)
				So(err, ShouldBeNil)

				d, err := database.ReadFromDB(db)
				So(err, ShouldBeNil)

				So(len(d), ShouldEqual, s)
			})
		})
	})
}
