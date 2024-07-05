package mFinance

import "github.com/kokizzu/gotro/D/Tt"

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
				{Name: `Bank`, Label: `bankAccounts`},
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
				{Name: `Aktiva Lain-lain`},
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
)

const (
	TableBusinessTransaction Tt.TableName = `businessTransaction`

	StartDate = `startDate`
	EndDate   = `endDate`
)

const (
	TableTransactionJournal Tt.TableName = `transactionJournal`

	DebitIDR  = `debitIDR`
	CreditIDR = `creditIDR`
	Date      = `date`
	DetailObj = `detailObj`
)

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
	TableTransactionTemplate: {
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
	TableTransactionTplDetail: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
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
}
