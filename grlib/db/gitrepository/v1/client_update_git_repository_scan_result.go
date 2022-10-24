package grgitrepositorydb

// UpdateGitRepositoryScanResultById create git repository scan result update helper
func (c *defaultClient) UpdateGitRepositoryScanResultById(id string) UpdateGitRepositoryScanResultDB {

	return NewUpdateGitRepositoryScanResultDBById(c.db, id)
}
