package company

import (
	"fmt"
	. "invoicer/person"
	"io"
)

type Billing struct {
	iBAN          string
	bIC           string
	instituteName string
}

func NewBilling(iban, bic, institute string) Billing {
	return Billing{
		iBAN:          iban,
		bIC:           bic,
		instituteName: institute,
	}

}

func (b Billing) FormattedBillingOuput(writer io.Writer) {
	fmt.Fprintf(writer, "%-10s\t%s\n%-10s\t%s\n%-10s\t%s\n", "IBAN", b.iBAN, "BIC", b.bIC, "Bank", b.instituteName)
}

type Company struct {
	Name          string
	CorporateForm string
	TradeRegister string
	Address
	Ceo Person
	Billing
}

func NewCompany(name, corporateform, traderegister string, address Address, ceo Person, billing Billing) Company {
	a = Company{
		Address{Street: },
	}
	return Company{
		Name:          name,
		CorporateForm: corporateform,
		TradeRegister: traderegister,
		Address:       address,
		Ceo:           ceo,
		Billing:       billing,
	}
}

func (c Company) GetFullName() string {
	return fmt.Sprintf("%s %s", c.Name, c.CorporateForm)
}

func (c Company) GetFullAddress() string {
	return fmt.Sprintf("%s\n%s", c.GetFullName(), c.GetAddress())
}
