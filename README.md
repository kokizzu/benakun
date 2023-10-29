
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
  - [ ] can manage users
  - [ ] can manage standard CoA (Chart of Account that all company must have)
  - [ ] can create entry-rule template
- TenantAdmin
  - [ ] can edit company
  - [ ] can create budgeting
  - [ ] can manage account source
	- [ ] can manage users
	- [ ] can manage accounts
	- [ ] can manage journals
	- [ ] can manage reports
  - [ ] can manage CoA
  - [ ] can manage entry-rule
  - [ ] can manage bank account
  - [ ] can manage external company (customer, supplier, etc)
  - [ ] can manage service/product/stock
  - [ ] can backup/restore database
- ReportViewer
  - [ ] can see all reports
- EntryUser
  - [ ] can input transaction/entry and proof
  - [ ] can request for update/delete entry and proof
- User
  - [ ] as new registered user can create a company
  - [ ] can change password TODO: UI
- Guest
  - [x] can register 
  - [x] can forgot/request reset password
  - [x] can reset password
  - [x] can login with username and password
  - [ ] can join invitation link to a company
  - [ ] can register/login with google account

## Features

- [ ] multi-tenant
- [ ] two language (ID, EN)
- [ ] multi-currency (USD, EUR, IDR)
- [ ] multi-user tenant (one tenant can have many users)
- [ ] Reports / Laporan (TODO: find what are the standards and how to use calculate them)
  - [ ] Profit and Loss / Laporan Laba Rugi --> priority
  - [ ] Balance Sheet / Laporan Posisi/Neracaya Keuangan --> priority
  - [ ] Cash Flow / Laporan Arus Kas
  - [ ] Statement of Changes in Equity / Laporan Perubahan Ekuitas/Modal
  - [ ] Notes to Financial Statements / Catatan atas Laporan Keuangan
  - [ ] Other Comprehensive Income / Laporan Laba Rugi Komprehensif Lain
  - [ ] Statement of Financial Position / Laporan Posisi Keuangan
  - [ ] Statement of Comprehensive Income / Laporan Laba Rugi Komprehensif
  - [ ] Statement of Changes in Equity / Laporan Perubahan Ekuitas
  - [ ] Statement of Cash Flows / Laporan Arus Kas
  - [ ] Notes to Financial Statements / Catatan atas Laporan Keuangan
  - [ ] Other Comprehensive Income / Laporan Laba Rugi Komprehensif Lain
  - [ ] laporan penjualan (TODO: find english term)
  - [ ] pergerakan piutang (TODO: find english term)
  - [ ] pergerakan hutang (TODO: find english term)
- [ ] generate invoice
- [ ] generate receipt
- [ ] upload proof

## Design

- per tenant table
- TODO: database design

## TODO
- add /tos /privacy