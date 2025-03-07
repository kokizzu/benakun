
# BenAkun Free Accounting Software

free double-entry bookeeping accounting software for Indonesian personal or soho business (UMKM).

## How to Contribute

read README.md of https://github.com/kokizzu/street (the structure of this project is the same)

## Purpose

people without prior accounting knowledge should able to use this software, and still conform to government standard.

based on:
- http://iaiglobal.or.id/v03/files/draft_ed_sak_emkm_kompilasi.pdf
- http://iaiglobal.or.id/v03/files/file_sak/emkm/
- https://www.youtube.com/watch?v=DojvxP_rbH0
- https://www.youtube.com/watch?v=kP_dHnRC3yk

## Draft Specification

https://docs.google.com/document/d/11u00Ch0oE0EEIGylYg4FlDKLoTWySnHonpFznSPb5Ik/edit?usp=sharing

## Roles

- SuperAdmin
  - [ ] can manage tenants
  - [x] can manage users
  - [ ] can manage standard CoA (Chart of Account that all company must have)
  - [ ] can create entry-rule template
- TenantAdmin
  - [x] can edit company
  - [x] change organizational structure
  - [x] can create budgeting
  - [x] can manage account source
  - [x] can manage users
  - [x] can manage accounts
  - [x] can manage journals
  - [ ] can manage reports
  - [x] can manage CoA
  - [x] can manage entry-rule
  - [x] can manage bank account
  - [ ] can manage external company (customer, supplier, etc)
  - [ ] can manage service/product/stock
  - [ ] can backup/restore database
  - [ ] can send/revoke invitation link to his own company (send email), `tenant:X1:invited:YYYYMMDD` or `tenant:X1:revoked:YYYYMMDD`
  - [ ] can remove someone from his own company (send email), `tenant:X1:terminated:YYYYMMDD`
- ReportViewer
  - [ ] can see all reports (see below)
- EntryUser
  - [ ] can input transaction/entry and proof
  - [ ] can request for update/delete entry and proof
- User
  - [x] as new registered user can create a company
    - [ ] can update company name and head title
  - [ ] can see and switch company based on host/url
  - [ ] can see list of invitation/rejection to join a company
  - [ ] can change password
  - [ ] can join/reject invitation link to a company (send email to inviter), `tenant:X1:accepted:YYYYMMDD` or `tenant:X1:rejected:YYYYMMDD`
  - [ ] can leave from a company (send email), `tenant:X1:left:YYYYMMDD`, when accepting invite to new company, the state of old company automatically set to this.
- Guest
  - [x] can register 
  - [x] can forgot/request reset password
  - [x] can reset password
  - [x] can login with username and password
  - [ ] can register/login with google account
 
- Daemon / Cron
  - [ ] can remind invitation link to company
  - [ ] can auto-revoke invitation link to company after certain period (eg. 3 days)

## Features

- [ ] multi-tenant
- [ ] multiple database connection mapping (one tenant can be using specific database connection)
- [ ] two language (ID, EN)
- [ ] multi-currency (USD, EUR, IDR)
- [ ] multi-user tenant (one tenant can have many users)
- [ ] Reports / Laporan (TODO: find what are the standards and how to use calculate them)
  - [x] Profit and Loss / Laporan Laba Rugi --> priority
  - [x] Balance Sheet / Laporan Posisi / Neraca Keuangan --> priority
  - [ ] Cash Flow / Laporan Arus Kas
  - [ ] Statement of Changes in Equity / Laporan Perubahan Ekuitas/Modal
  - [ ] Notes to Financial Statements / Catatan atas Laporan Keuangan
  - [ ] Other Comprehensive Income / Laporan Laba Rugi Komprehensif Lain
  - [x] Statement of Financial Position / Laporan Posisi Keuangan
  - [ ] Statement of Comprehensive Income / Laporan Laba Rugi Komprehensif
  - [ ] Statement of Changes in Equity / Laporan Perubahan Ekuitas
  - [ ] Statement of Cash Flows / Laporan Arus Kas
  - [ ] Notes to Financial Statements / Catatan atas Laporan Keuangan
  - [ ] Other Comprehensive Income / Laporan Laba Rugi Komprehensif Lain
  - [ ] Sales Report / Laporan Penjualan
  - [ ] Accounts Receivable Report / Laporan Piutang
    - [ ] Customer Receivable Report (TODO: find Indonesian term)
    - [ ] Receivable Aging Report / Laporan Umur Piutang
    - [ ] Receivable Payment History Report (TODO: find Indonesian term)
  - [ ] Accounts Payable Report / Laporan Hutang
    - [ ] Voucher Activity Report (TODO: find Indonesian term)
    - [ ] Payable Aging Report / Laporan Umur Hutang
    - [ ] Payable Payment History Report (TODO: find Indonesian term)
- [ ] generate invoice
- [ ] generate receipt
- [ ] upload proof

## Design

- per tenant table
- TODO: database design

## TODO
- add /tos /privacy

## FAQ
- **Q**: got error when starting docker container: `LuajitError: builtin/config/applier/mkdir.lua:26: mkdir.apply[wal.dir]: failed to create directory /var/lib/tarantool/sys_env/app/instance-001: Error creating directory /var/lib/tarantool/sys_env: fio: Permission denied`
  - **a**: run ``sudo chown `whoami` -Rv _tmpdb``

## Testing

- **Coverage testing**: run ``USE_COMPOSE=true go test -timeout=9999s -cover -run TestFunctionName ./dir -v``