Docker - Deploying Your First Docker Container
Menjalankan image redis di bacground

![1](01/01.PNG)

Melihat container yang sedang berjalan pada background

![2](01/02.PNG)

Menjalankan redis pada background dengan nama redisHostPort pada port 6379

![3](01/03.PNG)

Menggunakan opsi -p 6379 memungkinkannya untuk mengekspos Redis tetapi pada port yang tersedia secara acak.

![4](01/04.PNG)

Melihat port yang digunakan (acak) dengan perintah:

![5](01/05.PNG)

Menampilkan daftar kontainer mengenai pemetaan port

![6](01/06.PNG)

Setiap data yang perlu disimpan di Host Docker, dan tidak di dalam kontainer, harus disimpan di /opt/docker/data/redis.

![7](01/07.PNG)

Menjalankan container Ubuntu dengan perintah ps untuk melihat semua proses yang berjalan dalam sebuah container.

![8](01/08.PNG)

Mendapatkan akses ke bash shell di dalam sebuah container.

![9](01/09.PNG)