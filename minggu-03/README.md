Docker - Deploying Your First Docker Container
Menjalankan image redis di bacground

![1](01/docker run -d redis.PNG)

Melihat container yang sedang berjalan pada background

![2](01/docker ps.PNG)

Menjalankan redis pada background dengan nama redisHostPort pada port 6379

![3](01/docker run port.PNG)

Menggunakan opsi -p 6379 memungkinkannya untuk mengekspos Redis tetapi pada port yang tersedia secara acak.

![4](01/redisDynamic.PNG)

Melihat port yang digunakan (acak) dengan perintah:

![5](01/redisDynamic 6379.PNG)

Menampilkan daftar kontainer mengenai pemetaan port

![6](01/docker ps.PNG)

Setiap data yang perlu disimpan di Host Docker, dan tidak di dalam kontainer, harus disimpan di /opt/docker/data/redis.

![7](01/docker Run -d.PNG)

Menjalankan container Ubuntu dengan perintah ps untuk melihat semua proses yang berjalan dalam sebuah container.

![8](01/docker run ubuntu.PNG)

Mendapatkan akses ke bash shell di dalam sebuah container.

![9](01/d run ubuntu bash.PNG)