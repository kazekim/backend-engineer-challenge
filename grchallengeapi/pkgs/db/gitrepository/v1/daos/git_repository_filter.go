package grcgitrepositorydbdaos

type GitRepositoryFilter struct {
	Id *string `db:"id"`
	Name     *string `db:"name"`
	IsActive *bool   `db:"is_active"`
}
