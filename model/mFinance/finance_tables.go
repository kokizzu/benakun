package mFinance

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
)

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
	LabelProducts 						= `products`
	LabelProduct 							= `product`
	LabelSuppliers						= `suppliers`
	LabelCustomer							= `customer`
	LabelCustomerReceivables	= `customer:receivables`
	LabelStaff								= `staff`
	LabelBankAccount					= `bankAccounts`
	LabelBankAccountCompany 	= `bankAccounts:company`
	LabelBankAccountStaff 		= `bankAccounts:staff`
	LabelFunders								= `funders`
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

const (
	CoaBank 										= `Bank`
	CoaPersediaanBarangDagangan = `Persediaan Barang Dagangan`
	CoaStok											= `Stock`
	CoaPiutangUsaha							= `Piutang Usaha`
	CoaCOGS											= `Cost of Goods Sold / Harga Pokok Penjualan`
	CoaIncome										= `Income / Pendapatan Usaha`
	CoaHutangDagang							= `Hutang Dagang`
	CoaPerlengkapan							= `Peralatan - Perlengkapan`
	CoaAdm											= `Beban Administrasi dan Umum`
	CoaPrive										= `Prive`
	CoaLiability								= `Liability / Hutang - Kewajiban`
)

const (
	TemplatePenjualan  = `Penjualan`
	TemplatePembelian = `Pembelian`
	TemplatePenanamanModal  = `Penanaman Modal`
	TemplatePengambilanPrive  = `Pengambilan Prive`
	TemplatePembayaranGaji = `Pembayaran Gaji`
	TemplatePembayaranListrik = `Pembayaran Listrik` 
	TemplatePembayaranHutang  = `Pembayaran Hutang`
)

// Generate transaction template
var TransactionTemplatesDefault = []struct{
	Name string
	Color string
}{
	{
		Name: TemplatePenjualan,
		Color: `#2563eb`,
	},
	{
		Name: TemplatePembelian,
		Color: `#2563eb`,
	},
	{
		Name: TemplatePembayaranGaji,
		Color: `#2563eb`,
	},
	{
		Name: TemplatePembayaranListrik,
		Color: `#2563eb`,
	},
}

// Generate transaction template detail to transaction template
var TransactionTemplateDetailsMap = map[string][]struct{
	CoaName string
	IsDebit bool
	Attributes []any
}{
	TemplatePenjualan: {
		{
			CoaName: CoaStok,
			IsDebit: false,
			Attributes: []any{
				AttributesChildOnly, AttributesSales,
			},
		},
		{
			CoaName: CoaBank,
			IsDebit: true,
			Attributes: []any{
				AttributesAutoSum, AttributesChildOnly,
			},
		},
		{
			CoaName: CoaPiutangUsaha,
			IsDebit: true,
			Attributes: []any{
				AttributesAutoSum,
			},
		},
		{
			CoaName: CoaCOGS,
			IsDebit: true,
			Attributes: []any{
				AttributeSelfSum,
			},
		},
		{
			CoaName: CoaIncome,
			IsDebit: false,
			Attributes: []any{
				AttributeSelfSum, AttributesChildOnly,
			},
		},
	},
	TemplatePembelian: {
		{
			CoaName: CoaBank,
			IsDebit: false,
			Attributes: []any{
				AttributesChildOnly,
				AttributesAutoSum,
			},
		},
		{
			CoaName: CoaHutangDagang,
			IsDebit: false,
			Attributes: []any{
				AttributesChildOnly,
				AttributesAutoSum,
			},
		},
		{
			CoaName: CoaStok,
			IsDebit: true,
			Attributes: []any{
				AttributesSales,
				AttributesChildOnly,
			},
		},
		{
			CoaName: CoaPerlengkapan,
			IsDebit: true,
			Attributes: []any{
				AttributesSales,
				AttributesChildOnly,
			},
		},
	},
	TemplatePengambilanPrive: {
		{
			CoaName: CoaBank,
			IsDebit: true,
			Attributes: []any{
				AttributesAutoSum,
			},
		},
		{
			CoaName: CoaPrive,
			IsDebit: false,
		},
	},
	TemplatePembayaranGaji: {
		{
			CoaName: CoaAdm,
			IsDebit: true,
			Attributes: []any{
				AttributesChildOnly,
			},
		},
		{
			CoaName: CoaBank,
			IsDebit: false,
			Attributes: []any{
				AttributesAutoSum,
			},
		},
	},
	TemplatePembayaranListrik: {
		{
			CoaName: CoaAdm,
			IsDebit: true,
			Attributes: []any{
				AttributesChildOnly,
			},
		},
		{
			CoaName: CoaBank,
			IsDebit: true,
			Attributes: []any{
				AttributesAutoSum,
			},
		},
	},
	TemplatePembayaranHutang: {
		{
			CoaName: CoaLiability,
			IsDebit: false,
			Attributes: []any{
				AttributesChildOnly,
			},
		},
		{
			CoaName: CoaBank,
			IsDebit: true,
			Attributes: []any{
				AttributesAutoSum,
			},
		},
	},
}

func GetCoaDefaults() []CoaDefault {
	return []CoaDefault{
		{
			Name: `Asset / Aktiva`,
			Children: []CoaDefault{
				{
					Name: CoaBank,
					Label: LabelBankAccount,
					Children: []CoaDefault{
						{
							Name: `Cash`,
						},
						{
							Name: `Rekening Utama Perusahaan`,
						},
					},
				},
				{Name: `Deposito Berjangka`},
				{Name: CoaPiutangUsaha, Label: LabelCustomerReceivables},
				{Name: CoaPersediaanBarangDagangan},
				{Name: `Pajak Dibayar Muka`},
				{Name: `Biaya Dibayar Muka`},
				{Name: `Investasi Jangka Panjang`},
				{Name: `Aktiva Tetap`}, 
				{Name: `Akumulasi Penyusutan Aktiva Tetap`}, // Gak yakin
				{Name: `Aktiva Tak Berwujud`},
				{Name: CoaStok, Label: LabelProducts},
				{Name: CoaPerlengkapan},
				{Name: `Aktiva Lain-lain`},
			},
		},
		{
			Name: CoaLiability,
			Children: []CoaDefault{
				{Name: CoaHutangDagang, Label: LabelSuppliers},
				{Name: `Uang Muka Pelanggan`},
				{Name: `Hutang Pajak`},
				{Name: `Biaya yang Masih Harus Dibayar`},
				{Name: `Hutang Jangka Panjang - Lancar`},
				{Name: `Hutang Lain-lain`},
				{Name: `Hutang Jangka Panjang`},
			},
		},
		{
			Name: `Equity / Ekuitas - Modal`,
			Children: []CoaDefault{
				{
					Name: `Penanaman Modal`,
					Children: []CoaDefault{
						{
							Name: `Modal dari Owner 1`,
						},
					},
				},
				{Name: `Saldo Laba`},
			},
		},
		{
			Name: CoaIncome, Label: LabelCustomer,
			// Children: []CoaDefault{
			// 	{Name: `Penjualan Barang Dagangan`},
			// 	{Name: `Pendapatan Jasa`},
			// },
		},
		{
			Name: `Expense / Beban Usaha`,
			Children: []CoaDefault{
				{Name: `Beban Pemasaran`},
				{Name: CoaAdm, Label: LabelStaff},
				{Name: CoaCOGS},
			},
		},
		{
			Name: `Other Incomes / Penghasilan Lain-lain`,
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
			Name: `Other Expenses / Biaya Lain-lain`,
			Children: []CoaDefault{
				{Name: CoaPrive},
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
	AttributeSelfSum		= `selfSum`
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
	SalesCount 		int64 `json:"salesCount"`
	SalesPriceIDR string `json:"salesPriceIDR"` // Currency must be string
}

func IsValidDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	
	return err == nil
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
			{TransactionTplId, Tt.Unsigned},
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
