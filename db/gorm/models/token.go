package models

type Token struct {
	Name string `json:"name" binding:"required""`
	//Id string `json:"id" binding:"required,oneof=USD EUR""`
	Id             string `json:"id" binding:"required"`
	Default_amount int    `json:"default_amount" binding:"required"`
}

type CreateTokenInput struct {
	Id             string `json:"id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Default_amount int    `json:"default_amount" binding:"required"`
}

type UpdateTokenInput struct {
	Name           string `json:"name"`
	Default_amount int    `json:"default_amount"`
}

type GetTokenRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (m *Gorm) DbTokensFind(t *[]Token) {
	m._db.Table("token_info_table").Find(t)
}

func (m *Gorm) DbTokenFind(t *Token, id string) error {
	if err := m._db.Table("token_info_table").Where("id = ?", id).First(t).Error; err != nil {
		return err
	}
	return nil
}

func (m *Gorm) DbTokenCreate(t *Token) {
	m._db.Table("token_info_table").Create(t)
}

func (m *Gorm) DbTokenUpdate(t *Token, u *UpdateTokenInput) {
	m._db.Table("token_info_table").Model(t).Updates(u)
}

func (m *Gorm) DbTokenDelete(t *Token) {
	m._db.Table("token_info_table").Delete(t)
}
