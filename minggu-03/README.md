Docker - Deploying Your First Docker Container
Menjalankan image redis di bacground

![1](01/Pict01.PNG)

Melihat container yang sedang berjalan pada background

![2](01/Pict02.PNG)

Menjalankan redis pada background dengan nama redisHostPort pada port 6379

![3](01/Pict03.PNG)

Menggunakan opsi -p 6379 memungkinkannya untuk mengekspos Redis tetapi pada port yang tersedia secara acak.

![4](01/Pict04.PNG)

Melihat port yang digunakan (acak) dengan perintah:

![5](01/Pict05.PNG)

Menampilkan daftar kontainer mengenai pemetaan port

![6](01/Pict06.PNG)

Setiap data yang perlu disimpan di Host Docker, dan tidak di dalam kontainer, harus disimpan di /opt/docker/data/redis.

![7](01/Pict07.PNG)

Menjalankan container Ubuntu dengan perintah ps untuk melihat semua proses yang berjalan dalam sebuah container.

![8](01/Pict08.PNG)

Mendapatkan akses ke bash shell di dalam sebuah container.

![9](01/Pict09.PNG)