package picker

type WithQUoteFieldSuffix interface {
	QuoteFieldSuffix() string
	SetQuoteFieldSuffix(quoteFieldSuffix string)
}
type quoteFieldSuffixParam struct {
	quoteFieldSuffix string
}

func (qf quoteFieldSuffixParam) QuoteFieldSuffix() string {
	return qf.quoteFieldSuffix
}
func (qf *quoteFieldSuffixParam) SetQuoteFieldSuffix(quoteFieldSuffix string) {
	qf.quoteFieldSuffix = quoteFieldSuffix

}
