package repository

import (
	"fmt"

	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/pkg/sheet"
	"gorm.io/gorm"
)

type CostDB struct {
	db *gorm.DB
}

func CostRepo(db *gorm.DB) *CostDB {

	return &CostDB{
		db: db,
	}
}

func (c *CostDB) Create(user_id uint, category_id uint, currency_id uint, summary int) int {
	res := c.db.Create(&domain.Costs{
		CurrencyID: currency_id,
		UsersID:    user_id,
		CategoryID: category_id,
		Summary:    summary,
	})

	if res.Error != nil {
		fmt.Print(res.Error)
	}
	return int(res.RowsAffected)
}

func (c *CostDB) GetByID(id string) *domain.Costs {
	costs := &domain.Costs{}
	err := c.db.First(costs, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return costs
}
func (c *CostDB) GetAll() *[]domain.Costs {
	costs := &[]domain.Costs{}
	err := c.db.Find(costs).Error
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(costs)
	return costs
}
func (c *CostDB) GetSelectsFields(s []string) *domain.Costs {
	costs := &domain.Costs{}
	err := c.db.Select(s).Find(costs).Error
	if err != nil {
		fmt.Print(err)
	}
	return costs
}
func (c *CostDB) UpdateFields(id string, field *domain.Costs) *domain.Costs {
	costs := &domain.Costs{}
	err := c.db.Model(costs).Where("id = ?", id).Updates(field).Error
	if err != nil {
		fmt.Print(err)
	}
	return costs
}

func (c *CostDB) Delete(id string) *domain.Costs {
	costs := &domain.Costs{}
	err := c.db.Delete(costs, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return costs
}

func (c *CostDB) Report() error {
	costs := domain.CostsReport{}
	c.db.Raw("select cs.name as category, sum(c.summary) as costs from costs c join categories cs on c.category_id = cs.id where c.users_id = 1 GROUP by cs.name").Scan(&costs)

	f := sheet.Excelize()

	f.SetSheetName("Sheet1", "Отчет о тратах")
	for col, header := range []string{"Category", "Costs"} {
		cellName := fmt.Sprintf("%c1", 'A'+col)
		f.SetCellValue("Отчет о тратах", cellName, header)
	}

	for rowIdx, value := range costs {
		f.SetCellValue("Отчет о тратах", fmt.Sprintf("A%d", rowIdx+2), value.Category)
		f.SetCellValue("Отчет о тратах", fmt.Sprintf("B%d", rowIdx+2), value.Costs)
	}

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
