package mFinance

import (
	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
)

// TODO: business transaction
// Fields: startDate, endDate

// TODO: transaction journal
// Fields: trxStart, trxEnd, parentTransaction

const (
	Id         = `id`
	CreatedAt  = `createdAt`
	CreatedBy  = `createdBy`
	UpdatedAt  = `updatedAt`
	UpdatedBy  = `updatedBy`
	DeletedAt  = `deletedAt`
	DeletedBy  = `deletedBy`
	RestoredBy = `restoredBy`
	TenantCode = `tenantCode`
)

const (
	TableCoa Tt.TableName = `coa`

	Name     = `name`
	ParentId = `parentId`
	Children = `children`
	Label    = `label` // bankAccounts:company locations
)

const (
	LabelProducts 					= `products`
	LabelProduct 						= `product`
	LabelSuppliers					= `suppliers`
	LabelCustomer						= `customer`
	LabelStaff							= `staff`
	LabelBankAccount				= `bankAccounts`
	LabelBankAccountCompany = `bankAccounts:company`
	LabelBankAccountStaff 	= `bankAccounts:staff`
)

func GetLabelsMap() map[string]string {
	return map[string]string{
		LabelProducts: `Products`,
		LabelProduct: `Product`,
		LabelSuppliers: `Suppliers`,
		LabelCustomer: `Customer`,
		LabelStaff: `Staff`,
		LabelBankAccount: `Bank Account`,
		LabelBankAccountCompany: `Bank Account - Company`,
		LabelBankAccountStaff: `Bank Account - Staff`,
	}
}

type CoaDefault struct {
	Name     string
	Label    string
	Children []CoaDefault
}

func GetCoaDefaults() []CoaDefault {
	return []CoaDefault{
		{
			Name: `Aktiva`,
			Children: []CoaDefault{
				{Name: `Bank`, Label: LabelBankAccount},
				{Name: `Deposito Berjangka`},
				{Name: `Piutang Usaha`},
				{Name: `Persediaan Barang Dagangan`},
				{Name: `Uang Muka`},
				{Name: `Pendapatan yang Masih Harus Diterima`},
				{Name: `Pajak Dibayar Muka`},
				{Name: `Biaya Dibayar Muka`},
				{Name: `Investasi Jangka Panjang`},
				{Name: `Aktiva Tetap`},
				{Name: `Akumulasi Penyusutan Aktiva Tetap`},
				{Name: `Aktiva Tak Berwujud`},
				{Name: `Aktiva Lain-lain`, Label: LabelProducts},
			},
		},
		{
			Name: `Kewajiban`,
			Children: []CoaDefault{
				{Name: `Hutang Dagang`},
				{Name: `Uang Muka Pelanggan`},
				{Name: `Hutang Pajak`},
				{Name: `Biaya yang Masih Harus Dibayar`},
				{Name: `Hutang Jangka Panjang - Lancar`},
				{Name: `Hutang Lain-lain`},
				{Name: `Hutang Jangka Panjang`},
			},
		},
		{
			Name: `Ekuitas`,
			Children: []CoaDefault{
				{Name: `Modal`},
				{Name: `Saldo Laba`},
			},
		},
		{
			Name: `Pendapatan Usaha`,
			Children: []CoaDefault{
				{Name: `Pendapatan Usaha - Penjualan Barang Dagangan`},
				{Name: `Pendapatan Usaha - Jasa Keagenan dan Distributor`},
			},
		},
		{Name: `Harga Pokok Penjualan`},
		{
			Name: `Beban Usaha`,
			Children: []CoaDefault{
				{Name: `Beban Pemasaran`},
				{Name: `Beban Administrasi dan Umum`},
			},
		},
		{
			Name: `Penghasilan Lain-lain`,
			Children: []CoaDefault{
				{Name: `Penghasilan Bunga Deposito`},
				{Name: `Penghasilan Bunga Obligasi`},
				{Name: `Penghasilan Deviden`},
				{Name: `Penghasilan Bunga Jasa Giro`},
				{Name: `Laba Penjualan Aktiva Tetap`},
				{Name: `Penghasilan Sewa`},
				{Name: `Laba Selisih Kurs`},
				{Name: `Penghasilan Lainnya`},
			},
		},
		{
			Name: `Beban Lain-lain`,
			Children: []CoaDefault{
				{Name: `Beban Pajak Jasa Giro`},
				{Name: `Beban Administrasi Jasa Giro`},
				{Name: `Rugi Penjualan Aktiva Tetap`},
				{Name: `Rugi Selisih Kurs`},
				{Name: `Beban Lainnya`},
			},
		},
	}
}

const (
	TableTransactions Tt.TableName = `transactions`

	CompletedAt = `completedAt`
	CoaId       = `coaId`
	Price       = `price`
	Description = `descriptions`
	Qty         = `qty`
)

const (
	TableTransactionTemplate Tt.TableName = `transactionTemplate`

	Color    = `color`
	ImageURL = `imageURL`
)

const (
	TableTransactionTplDetail Tt.TableName = `transactionTplDetail`

	IsDebit           = `isDebit`
	IsAlwaysStartDate = `isAlwaysStartDate`
	Attributes				= `attributes`
)

const (
	AttributesAutoSum 	= `autoSum`
	AttributesChildOnly	= `childOnly`
	AttributesSales			= `sales`
)

func IsValidAttributes(attr []any) bool {
	if len(attr) > 0 {
		for _, v := range attr {
			switch v {
			case AttributesAutoSum, AttributesChildOnly, AttributesSales:
				continue
			default:
				return false
			}
		}
	}

	return true
}

const (
	TableBusinessTransaction Tt.TableName = `businessTransaction`

	StartDate 				= `startDate`
	EndDate   				= `endDate`
	CustBankAccountId = `custBankAccountId` 
	TotalIDR 					= `totalIDR`
	TransactionTplId  = `transactionTemplateId`
)

const (
	TableTransactionJournal Tt.TableName = `transactionJournal`

	DebitIDR  = `debitIDR`
	CreditIDR = `creditIDR`
	Date      = `date`
	DetailObj = `detailObj`
)

type TransactionJournalDetailObject struct {
	SalesCount 		uint64 `json:"salesCount"`
	SalesPriceIDR uint64 `json:"salesPriceIDR"`
}

func IsValidDetailObject(in string) bool {
	var trxJournalDetailObj TransactionJournalDetailObject
	err := json.Unmarshal([]byte(in), &trxJournalDetailObj)

	_ = trxJournalDetailObj // throw to hole

	return err == nil
}

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableCoa: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{Name, Tt.String},
			{Label, Tt.String},
			{ParentId, Tt.Unsigned},
			{Children, Tt.Array},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
	TableTransactions: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{CompletedAt, Tt.Integer},
			{Price, Tt.Unsigned},
			{Description, Tt.String},
			{Qty, Tt.Integer},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
	TableTransactionTemplate: { // Ini jenis transaksi yang mungkin terjadi di Perusahaan
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{Name, Tt.String},
			{Color, Tt.String},
			{ImageURL, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
	// Transaction template detail lebih dari 1 utk parent yg sama
	TableTransactionTplDetail: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{ParentId, Tt.Unsigned},
			{TenantCode, Tt.String},
			{CoaId, Tt.Unsigned},
			{IsDebit, Tt.Boolean},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{Attributes, Tt.Array},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
	TableTransactionJournal: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{CoaId, Tt.Unsigned},
			{DebitIDR, Tt.Integer},
			{CreditIDR, Tt.Integer},
			{Description, Tt.String},
			{Date, Tt.String},
			{DetailObj, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
	TableBusinessTransaction: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{StartDate, Tt.String},
			{EndDate, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{TransactionTplId, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
}
