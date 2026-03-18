# Catatan Modul 06: Maps Deep Dive

-   **CRUD**: Create (`m[k]=v`), Read (`v := m[k]`), Update (`m[k]=newV`), Delete (`delete(m, k)`).
-   **Comma OK Idiom**: Sintaks `val, ok := m[key]` adalah cara standar Go untuk mengecek keberadaan data.
-   **Zero Value**: Jika memanggil key yang tidak ada tanpa _Comma OK_, Go mengembalikan _zero value_ dari tipe datanya (misal `0` untuk int, `""` untuk string).
-   **Unordered**: Map tidak menjamin urutan data saat di-loop.
