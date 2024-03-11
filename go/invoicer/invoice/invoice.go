package invoice

import (
	"fmt"
	. "invoicer/company"
	. "invoicer/person"
	"invoicer/utils"
	"io"
)

type Position struct {
	Title string
	Count float64
	Price float64
}

func NewPosition(title string, count, price float64) Position {
	return Position{Count: count, Title: title, Price: price}
}
func GeneratePositions() []Position {
	return []Position{
		NewPosition("Lack ausbessern", 1, 12.99),
		NewPosition("Ausbeulen", 2, 34.95),
	}
}

type Invoice struct {
	Customer      Recipient
	Dealer        Company
	InvoiceNumber string
	Date          string
	Positions     []Position
	Currency      string
	Tax           float64
}

func NewInvoice(customer Recipient, dealer Company, positions []Position, currency string, tax float64) Invoice {
	return Invoice{
		Customer:      customer,
		Dealer:        dealer,
		Positions:     positions,
		Date:          utils.GetCurrentDateAndTimeString(),
		InvoiceNumber: utils.GenerateInvoiceNumber(),
		Currency:      currency,
		Tax:           tax,
	}
}

func (i *Invoice) GetFormattedPrice(price float64) string {

	if i.Currency == "$" {
		return fmt.Sprintf("%s%.2f", i.Currency, price)

	}
	return fmt.Sprintf("%.2f%s", price, i.Currency)

}

func (i *Invoice) ListInvoicePositions(writer io.Writer) {
	fmt.Fprintf(writer, "#) %-20s \tPRICE \tCOUNT \tSUM\n", "DESCRIPTION")
	for num, pos := range i.Positions {
		fmt.Fprintf(writer, "%d) %-20s \t%s\t%.1f\t%s\n", num+1, pos.Title, i.GetFormattedPrice(pos.Price), pos.Count, i.GetFormattedPrice(pos.Count*pos.Price))
	}
}

func (i *Invoice) sumUp() (sum, sumWithTax float64) {

	for _, pos := range i.Positions {
		sum += pos.Price * pos.Count
	}
	sumWithTax = (sum * i.Tax) + sum
	return sum, sumWithTax
}

func (i *Invoice) Print(writer io.Writer) {
	sum, sumWithTax := i.sumUp()

	_, _ = fmt.Fprintf(writer, "%-40s %-11s %10s\n", "### INVOICE ###", "Invoice No:", i.InvoiceNumber)
	_, _ = fmt.Fprintf(writer, "%-40s %-11s %10s\n", "To: ", "Date:", i.Date[0:11])
	_, _ = fmt.Fprintln(writer, i.Customer.GetFullAddress())
	_, _ = fmt.Fprintln(writer, "")

	_, _ = fmt.Fprintln(writer, "Issuer:")
	_, _ = fmt.Fprintln(writer, i.Dealer.GetFullAddress())
	_, _ = fmt.Fprintln(writer, "")
	_, _ = fmt.Fprintln(writer, "Positions:")
	_, _ = fmt.Fprintln(writer, "________________________________________________________________")

	i.ListInvoicePositions(writer)
	_, _ = fmt.Fprintln(writer, "________________________________________________________________")
	_, _ = fmt.Fprintf(writer, "%-25s\t\t\t%s\n", "Total:", i.GetFormattedPrice(sum))
	_, _ = fmt.Fprintf(writer, "%-25s\t\t\t%s\n", "Total with tax:", i.GetFormattedPrice(sumWithTax))

	_, _ = fmt.Fprintln(writer, "")
	_, _ = fmt.Fprintln(writer, "Please transfer the invoice amount to the given account within the payment deadline of 14 days.")
	_, _ = fmt.Fprintln(writer, "\nBank details:")
	i.Dealer.FormattedBillingOuput(writer)
}
