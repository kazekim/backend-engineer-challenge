package grgitrepositorydb

// UpdateGitRepositoryById create git repository update helper
func (c *defaultClient) UpdateGitRepositoryById(id string) UpdateGitRepositoryDB {

	return NewUpdateGitRepositoryDBById(c.db, id)
}
