package wikigraphql

import (
	"fmt"
	"time"
	"wiki-api/sql"

	"github.com/graphql-go/graphql"
)

// QueryResolver struct keeping sql.Db connection for the query resolver
type QueryResolver struct {
	db *sql.Db
}

// GetPagesResolver resolver for getting pages from database
func (r *QueryResolver) GetPagesResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("GetPagesResolver initiliazed")
	title, ok := p.Args["Title"].(string)
	if ok {
		pages := r.db.GetPagesByTitle(title)
		return pages, nil
	}
	tags, ok := p.Args["Tags"].(string)
	if ok {
		pages := r.db.GetPagesByTag(tags)
		return pages, nil
	}

	return nil, nil
}

// MutationResolver struct keeping sql.Db connection for the mutation resolver
type MutationResolver struct {
	db *sql.Db
}

// SavePageResolver resolver for saving pages to database
func (r *MutationResolver) SavePageResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("SavePageResolver initiliazed")
	var page sql.WikiPage
	page.Title, _ = p.Args["Title"].(string)
	page.Tags, _ = p.Args["Tags"].(string)
	page.Ingress, _ = p.Args["Ingress"].(string)
	page.MainText, _ = p.Args["MainText"].(string)
	page.SideBarInfo, _ = p.Args["SideBarInfo"].(string)
	page.ProfileImagePath, _ = p.Args["ProfileImagePath"].(string)
	page.BodyImagePath, _ = p.Args["BodyImagePath"].(string)
	page.Visible, _ = p.Args["Visible"].(bool)
	page.Author, _ = p.Args["Author"].(string)
	page.LastChanged = time.Now()
	err := r.db.Save(page)
	if err != nil {
		return nil, err
	}
	return &page, nil
}
