# Simple Plan Log
![Coverage](https://img.shields.io/badge/Coverage-87.1%25-brightgreen)

## Struktur Direktori

```shell
.
├── cmd
│   └── api
├── model
├── port
├── repo
└── router
```

- `cmd`: direktori untuk generate executable command
- `model`: model yang akan digunakan untuk menyimpan data
- `port`: berisi kumpulan interface sebagai layer penghubung internal system dan external system
- `repo`: direktori untuk implementasi adapter repository yang sebagai layer penghubung antara model dan database

> Karena goalsnya sederhana, service layer sengaja tidak dibuat
