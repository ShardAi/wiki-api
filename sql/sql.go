package sql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	// imported to allow for sql to use "sqlserver" and not used directly
	_ "github.com/denisenkom/go-mssqldb"
)

// Db contains sql db and context structures
type Db struct {
	*sql.DB
	ctx context.Context
}

// New creates a new instance of Db and returns it along or any errors encountered along the way
func New(connectionString string) (*Db, error) {
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return nil, err
	}

	// Use background context
	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &Db{db, ctx}, nil
}

// ConnString creates and returns a new connectionstring to connect to Db
func ConnString(host string, port int) string {
	return fmt.Sprintf("server=%s;port=%d;Integrated Security=SSPI", host, port)
}

// WikiPage structure representing a WikiPage record in the database
type WikiPage struct {
	Title            string    `json:"Title"`
	Tags             string    `json:"Tags"`
	Ingress          string    `json:"Ingress"`
	MainText         string    `json:"MainText"`
	SideBarInfo      string    `json:"SideBarInfo"`
	ProfileImagePath string    `json:"ProfileImagePath"`
	BodyImagePath    string    `json:"BodyImagePath"`
	Visible          bool      `json:"Visible"`
	Author           string    `json:"Author"`
	LastChanged      time.Time `json:"LastChanged"`
}

// Save will update or insert the given page into the database
func (d *Db) Save(page WikiPage) error {
	var visibleBit int
	if page.Visible {
		visibleBit = 1
	} else {
		visibleBit = 0
	}

	statement := fmt.Sprintf("UPDATE ElinorWiki.dbo.WikiPage "+
		"SET Title=%[1]s, "+
		"Tags=%[2]s, "+
		"Ingress=%[3]s, "+
		"MainText=%[4]s, "+
		"SideBarInfo=%[5]s, "+
		"ProfileImagePath=%[6]s, "+
		"BodyImagePath=%[7]s, "+
		"Visible=%[8]d, "+
		"Author=%[9]s, "+
		"LastChanged=getdate() "+
		"WHERE Title=%[1]s "+
		"IF @@ROWCOUNT=0 "+
		"INSERT INTO ElinorWiki.dbo.WikiPage(Title,Tags,Ingress,MainText,SideBarInfo,ProfileImagePath,BodyImagePath,Visible,Author,LastChanged) "+
		"VALUES(%[1]s,%[2]s,%[3]s,%[4]s,%[5]s,%[6]s,%[7]s,%[8]d,%[9]s,getdate())",
		"'"+page.Title+"'",
		"'"+page.Tags+"'",
		"'"+page.Ingress+"'",
		"'"+page.MainText+"'",
		"'"+page.SideBarInfo+"'",
		"'"+page.ProfileImagePath+"'",
		"'"+page.BodyImagePath+"'",
		visibleBit,
		"'"+page.Author+"'")

	fmt.Println("Running statement on database:\n", statement)

	// Run query and scan for result
	_, err := d.QueryContext(d.ctx, statement)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
		return err
	}

	return nil
}

// GetPagesByTitle returns any pages with the given title. This SHOULD only ever be one page
func (d *Db) GetPagesByTitle(title string) []WikiPage {
	titleString := strings.Replace(title, "\"", "", -1)

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := d.PingContext(d.ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	statement := fmt.Sprintf("SELECT "+
		"* "+
		"FROM ElinorWiki.dbo.WikiPage wikiPage "+
		"WHERE Title='%s'", titleString)

	fmt.Println("Running statement on database:\n", statement)

	// Run query and scan for result
	rows, err := d.QueryContext(d.ctx, statement)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	var r WikiPage
	pages := []WikiPage{}
	for rows.Next() {
		err = rows.Scan(
			&r.Title,
			&r.Tags,
			&r.Ingress,
			&r.MainText,
			&r.SideBarInfo,
			&r.ProfileImagePath,
			&r.BodyImagePath,
			&r.Visible,
			&r.Author,
			&r.LastChanged,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		pages = append(pages, r)
	}
	return pages
}

// GetPagesByTag returns any WikiPage that has a tag that matches the one given in the search
func (d *Db) GetPagesByTag(tag string) []WikiPage {

	tagString := "%" + strings.Replace(tag, "\"", "", -1) + "%"

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := d.PingContext(d.ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	statement := fmt.Sprintf("SELECT "+
		"* "+
		"FROM ElinorWiki.dbo.WikiPage wikiPage "+
		"WHERE Tags LIKE '%s'", tagString)

	fmt.Println("Running statement on database:\n", statement)

	// Run query and scan for result
	rows, err := d.QueryContext(d.ctx, statement)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	var r WikiPage
	pages := []WikiPage{}
	for rows.Next() {
		err = rows.Scan(
			&r.Title,
			&r.Tags,
			&r.Ingress,
			&r.MainText,
			&r.SideBarInfo,
			&r.ProfileImagePath,
			&r.BodyImagePath,
			&r.Visible,
			&r.Author,
			&r.LastChanged,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		pages = append(pages, r)
	}
	return pages
}
