## Docker - Orchestration using Docker Compose

> Step 1 -  Menentukan Container Pertama

Docker Compose berdasar pada file docker-compose.yml. File docker-compose.yml mendefinisikan semua container dan seting/pengaturan yang dibutuhkan untuk meluncurkan beberapa/kumpulan cluster. Properti memetakan ke bagaimana kita menggunakan perintah menjalankan docker, sekarang disimpan dalam controll sumber dan dibagikan bersama dengan kode kita.

Format file berdasar pada YAML (Yet Another Markup Language).

 ```container_name:property: value- or options```

Menentukan Kontainer Pertama

Pada skenario ini, kita memiliki aplikasi Node.js dan membutuhkan koneksi ke Redis. Untuk memulainya, kita harus mendefinisikan file docker-compose.yml untuk launching aplikasi Node.js.

Dengan format di atas, file harus memberi nama kontainer 'web' dan mengatur properti build ke direktori saat ini. 

 ```web:build: .```

Melakukan pendefinisian container yang disebut web, yang didasarkan pada pembangunan direktori saat ini.

![01](img/img1.png)

> Step 2 - Menentukan Pengaturan

Docker Compose mendukung semua properti yang dapat didefinisikan menggunakan

 ```docker run```.

Untuk menghubungkan dua kontainer bersama untuk menentukan properti tautan dan mendaftar koneksi yang diperlukan. Misalnya, berikut ini akan membuat link ke sumber kontainer redis yang didefinisikan dalam file yang sama dan menetapkan nama yang sama ke alias.

```links:- redis```

Format yang sama digunakan untuk properti lain seperti port

```ports: - "3000"- "8000"```

![02](img/img1b.png)

Tugas:

 Perbarui kontainer web  untuk mengekspos port 3001 dan buat tautan ke kontainer Redis.

> Step 3 -  Mendefinisikan Kontainer Kedua

Pada langkah sebelumnya, Kita menggunakan Dockerfile di direktori saat ini sebagai dasar untuk kontainer Kita. Pada langkah ini, Kita ingin menggunakan gambar yang ada dari Docker Hub sebagai kontainer kedua.

Untuk menemukan kontainer kedua Anda cukup menggunakan format yang sama seperti sebelumnya pada baris baru. Format YAML cukup fleksibel untuk mendefinisikan banyak kontainer dalam file yang sama.

Tugas: Tentukan Kontainer Kedua

Tentukan kontainer kedua dengan nama redis yang menggunakan gambar redis. Mengikuti format YAML, detail kontainer adalah:

```redis:image: redis:alpinevolumes: - /var/redis/data:/data```

![03](img/img1c.png)


> Step 4 -  Docker Up

Dengan file docker-compose.yml dibuat di tempat, Anda dapat meluncurkan semua aplikasi dengan satu perintah ```up```. Jika Anda ingin memunculkan satu kontainer, maka Anda dapat menggunakan ```up <nama>```.

Argumen -d menyatakan untuk menjalankan kontainer di latar belakang, mirip dengan saat digunakan dengan ```docker run```.

Tugas:

Luncurkan aplikasi Anda menggunakan 
```docker-compose up -d```

![04](img/img2.png)

![05](img/img2-2.png)


> Step 5 -  Manajemen Docker

Docker Compose tidak hanya dapat mengelola kontainer awal tetapi juga menyediakan cara mengelola semua kontainer menggunakan satu perintah.

Misalnya, untuk melihat detail kontainer yang diluncurkan, Anda dapat menggunakan 

```docker-compose ps```.

![06](img/img3.png)

Untuk mengakses semua log melalui satu aliran, Anda menggunakan 

```docker-compose logs```.

![07](img/img4.png)


> Step 6 -  Docker Scale

Karena Docker Compose memahami cara meluncurkan kontainer aplikasi, itu juga dapat digunakan untuk mengukur jumlah kontainer yang berjalan.

Scale option memungkinkan Anda untuk menentukan layanan dan jumlah instance yang Anda inginkan. Jika angkanya lebih besar dari instans yang sudah berjalan, maka akan meluncurkan kontainer tambahan. Jika jumlahnya kurang, maka itu akan menghentikan kontainer yang tidak diminta.

Tugas:

Skala jumlah kontainer web yang Anda jalankan menggunakan perintah 

```docker-compose scale web = 3```

![08](img/img5.png)

Anda dapat menurunkannya kembali menggunakan 

```docker-compose scale web = 1```

![09](img/img6.png)

> Step 7 -  Docker Stop

Seperti ketika kami meluncurkan aplikasi, untuk menghentikan satu set kontainer, Anda dapat menggunakan perintah 

```docker-compose stop```.

![10](img/img7.png)

Untuk menghapus semua wadah gunakan perintah 

```docker-compose rm```.

![11](img/img8.png)

* [<<=  Back](README.md)