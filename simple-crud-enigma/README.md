# Aplikasi Inventaris Buku

Buatlah sebuah aplikasi sederhana untuk mengelola data buku menggunakan bahasa pemrograman Golang. Aplikasi ini memiliki fitur-fitur berikut:

## Tugas 1 : Membuat Struct Data Book

Buatlah Struct data `Book` dengan properti berikut:

- Id (int): Id dari sebuah `Book`.
- Title (string): Judul dari sebuah `Book`.
- Author (string): Penulis dari sebuah `Book`.
- ReleaseYear (string): Tahun terbit dari sebuah `Book`.
- Pages (int): Jumlah halaman dari sebuah `Book`.

## Tugas 2 : Menampilkan Seluruh Data Buku

Buatlah fitur yang dapat menampilkan seluruh data buku yang ada.

**Spesifikasi**

- Buatlah sebuah fungsi `viewAllBooks()` yang akan menampilkan seluruh buku dengan format yang rapi.
- Gunakan Struct data `Book` untuk menyimpan dan mengelola data buku.Tampilkan daftar buku yang ada dengan format berikut:

```yaml
==================================================
Book - 1
Book Id: [Id]
Book Title: [Title]
Book Author: [Author]
Book ReleaseYear: [ReleaseYear]
Book Pages: [Pages]
==================================================

==================================================
Book - 2
Book Id: [Id]
Book Title: [Title]
Book Author: [Author]
Book ReleaseYear: [ReleaseYear]
Book Pages: [Pages]
==================================================

...
```

## Tugas 3: Menambahkan Buku Baru

Buatlah fitur yang dapat menambahkan buku baru ke dalam data buku.

**Spesifikasi**

- Buatlah sebuah fungsi `addNewBook()` yang akan meminta pengguna memasukkan data buku baru.
- Gunakan Struct data `Book` untuk menyimpan data buku baru.
- Data buku baru yang diminta pengguna meliputi Id, judul, penulis, tahun terbit, dan jumlah halaman.
- Setelah pengguna memasukkan data buku baru, tambahkan buku tersebut ke dalam data buku yang ada.
- Jika proses penambahan berhasil, tampilkan pesan sukses.
- Jika terjadi kesalahan dalam proses penambahan, tampilkan pesan error.

## Tugas 4 : Menulis Data Buku ke File CSV

Buatlah fungsi yang dapat menulis data buku ke dalam file CSV.

**Spesifikasi**

- Buatlah sebuah fungsi `saveDataToCSV()` yang akan menulis data buku ke dalam file CSV.
- Fungsi ini akan menerima parameter berupa slice dari objek `Book` yang akan ditulis ke file CSV.
- Format penulisan ke dalam file CSV harus sesuai dengan format yang digunakan pada Release 5.
- Tambahkan error handling jika terjadi kesalahan dalam menulis file CSV.

## Tugas 5 : Membaca Data Buku dari File CSV

Buatlah fungsi yang dapat membaca data buku dari file CSV dan menginisialisasi objek `Book` berdasarkan data yang dibaca.

**Spesifikasi**

- Buatlah sebuah fungsi `loadDataFromCSV()` yang akan membaca data buku dari file CSV.
- Fungsi ini harus mengembalikan slice dari objek `Book` yang berisi seluruh data buku yang ada.
- Format file CSV akan mengikuti struktur berikut:
  - Kolom pertama berisi ID buku.
  - Kolom kedua berisi judul buku.
  - Kolom ketiga berisi penulis buku.
  - Kolom keempat berisi tahun terbit buku.
  - Kolom kelima berisi jumlah halaman buku.
- Manipulasi data yang dibaca dari file CSV untuk menginisialisasi objek `Book`.
- Tambahkan error handling jika terjadi kesalahan dalam membaca file atau menginisialisasi objek `Book`.

## Tugas 6 : Mengubah Data Buku

Buatlah fitur yang dapat mengubah data buku berdasarkan Id.

**Spesifikasi**

- Buatlah sebuah fungsi `updateBook()` yang akan meminta pengguna memasukkan Id buku yang akan diupdate dan data buku yang baru.
- Cari buku dengan Id yang sesuai dalam data buku yang ada
- Jika buku ditemukan, update data buku tersebut dengan data baru yang dimasukkan oleh pengguna.
- Jika proses update berhasil, tampilkan pesan sukses.
- Jika terjadi kesalahan dalam proses update, tampilkan pesan error.

## Tugas 7 : Menghapus Buku

Buatlah fitur yang dapat menghapus buku berdasarkan Id.

**Spesifikasi**

- Buatlah sebuah fungsi `deleteBook()` yang akan meminta pengguna memasukkan Id buku yang akan dihapus.
- Cari buku dengan Id yang sesuai dalam data buku yang ada.
- Jika buku ditemukan, hapus buku tersebut dari data buku.
- Jika proses penghapusan berhasil, tampilkan pesan sukses.
- Jika terjadi kesalahan dalam proses penghapusan, tampilkan pesan error.


## Release 8 : Integrasi Fungsi dengan Menu Utama

Integrasikan fungsi-fungsi yang telah dibuat ke dalam menu utama aplikasi.

**Spesifikasi**

- Buatlah sebuah menu utama yang akan menampilkan opsi-opsi berikut kepada pengguna:
  - View All Books
  - Add New Book
  - Update Book
  - Delete Book
  - Exit
- Tampilkan pesan kepada pengguna untuk memilih opsi yang diinginkan.
- Terima input dari pengguna berupa nomor opsi yang dipilih.
- Panggil fungsi-fungsi yang sesuai untuk menjalankan opsi yang dipilih.
- Jika pengguna memilih opsi 5 (Exit), maka akan keluar dari program.
