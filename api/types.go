package api

type responseStatus string

const (
	ok    responseStatus = "ok"
	notOk responseStatus = "error"
)

type addPortfoliioSuccess struct {
	PortfolioID string `json:"createdPortfolioId" example:"5edb2a0e550dfc5f16392838"`
}

type addOperationSuccess struct {
	OperationID string `json:"createdOperationId" example:"5edbc0a72c857652a0542fab"`
}

type putPortfoliioSuccess struct {
	HasModified bool `json:"hasModified" example:"true"`
}

type commonResponse struct {
	Status string `json:"status" example:"ok"`
}

type delPortfoliioSuccess struct {
	HasDeleted bool `json:"hasDeleted" example:"true"`
}

type delMutileSuccess struct {
	DeletedItems int64 `json:"DeletedItems" example:"42"`
}

type gmailAuthURLSuccess struct {
	URL string `json:"url" example:"https://google.com"`
}

type getBalanceSuccess struct {
	Balance float64 `json:"balance" example:"42"`
}

type getAverageSuccess struct {
	Average float64 `json:"avg" example:"42"`
}

type portfolioRequest struct {
	Name        string `json:"name" example:"Best portfolio"`
	Description string `json:"description" example:"Best portfolio ever!!!"`
}

type operationRequest struct {
	Currency      string  `json:"currency" example:"USD"`
	Price         float64 `json:"price" example:"293.61"`
	Volume        int64   `json:"vol" example:"100"`
	FIGI          string  `json:"figi" example:"BBG00MVRXDB0"`
	ISIN          string  `json:"isin" example:"US9229083632"`
	Ticker        string  `json:"ticker" example:"VOO"`
	DateTime      string  `json:"date" example:"2020-06-06T15:54:05Z"`
	OperationType string  `json:"type" example:"sell"`
}

type priceRequest struct {
	Time   int64   `json:"time" example:"1467590400"`
	Price  float64 `json:"price" example:"293.61"`
	Volume int     `json:"vol" example:"100"`
}

type errorResponse struct {
	Error string `json:"error" example:"Something went wrong"`
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type tokenResponse struct {
	Status responseStatus `json:"status"`
	Token  string         `json:"token"`
}
