Berkomunikasi Antar Kontainer
Environment telah dikonfigurasikan dengan klien Docker dan daemon. Nama mesin yang digunakan daemon Docker dinamakan docker. Menggunakan docker localhost atau 0.0.0.0 untuk mengakses salah satu service.
Menghubungkan ke container aplikasi yang menghubungkan ke data store. Kunci saat membuat link adalah nama container.
Menjalankan sebuah container redis dengan nama redis-server yang akan dijadikan container sumber data store.

![1](img/pict1.png)

Kemudian gunakan opsi --link <nama|id container>:<alias> ketika menjalankan container baru yang ingin dihubungkan kepada container redis-server tersebut.
Mencoba menggunakan image alpine yang akan di-link ke container redis-server, dan didefinisikan aliasnya sebagai redis.
Docker akan men-set beberapa environment variable berdasarkan link ke sebuah container, yaitu informasi seperti port dan alamat IP. 

![2](img/pict2.png)

Selanjutnya, docker akan meng-update file HOSTS alpine container dengan menambahkan sebuah input untuk sumber container berupa  nama asli container, alias, dan hash id seperti berikut:

![3](img/pict3.png)

Selanjutnya mencoba melakukan ping antar container seperti suatu host terhubung dalam satu jaringan

![4](img/pict4.png)

Connect To App
Menghubungkan aplikasi node.js dengan hostname redis

![5](img/pict5.png)

Mengirimkan request HTTP ke aplikasi, selanjutnya akan menyimpan permintaan di Redis dan mengembalikan hitungan

![6](img/pict6.png)

Connect to Redis CLI
Dengan jalan yang sama, kita dapat menghubungkan sumber container ke aplikasi, kita juga dapat terhubung melalui tool CLI dengan perintah seperti berikut:

![7](img/pict7.png)

Untuk keluar dengan menggunakan key QUIT

![8](img/pict8.png)

==========
Docker Networks
Menghubungkan antar container melalui jaringan
Untuk menghubungkan beberapa container menggunakan network,kita harus membuat sebuah “network” terlebih dahulu dengan docker CLI, yang kemudian digunakan untuk meletakkan container-container pada network tersebut sehingga dapat saling berkomunikasi antar satu dengan lainnya.

![9](img/pict9.png)

Connect To Network
Menjalankan container redis yang akan dimasukan kedalam network backend-network:

![10](img/pict10.png)

Network Communication
Docker tidak lagi menambahkan environment variable dan meng-update file /etc/hosts pada container seperti sebelumnya, dibuktikan dengan menjalankan kembali container alpine didalam jaringan backend-network

![11](img/pict11.png)

![12](img/pict12.png)

Karena tidak meng-update environment variable, maka docker menggunakan embedded  DNS Server sebagai alat komunikasi yang digunakan setiap container yang tergabung dalam backend-network tersebut, alamat IP DNS Server ini ditambahkan kepada semua container melalui file/etc/resolv.conf.

![13](img/pict13.png)

Ketika container mencoba mengakses container lain melalui nama yang dikenal, seperti Redis, server DNS akan mengembalikan alamat IP wadah yang benar.

![14](img/pict14.png)

Connect Two Containers
Docker mendukung beberapa network dan container yang terhubung ke lebih dari satu network sekaligus.

![15](img/pict15.png)

Saat menggunakan perintah connect, dimungkinkan untuk melampirkan container yang ada ke network.
$ docker network connect frontend-network redis
Ketika menjalankan web server, karena terdapat pada network yang sama maka akan dapat berkomunikasi dengan instance Redis.

![16](img/pict16.png)

Melakukan test

![17](img/pict17.png)

Connect Container dengan Alias

![18](img/pict18.png)

Ketika container mencoba mengakses layanan melalui nama db, mereka akan diberi alamat IP dari container Redis.

![19](img/pict19.png)

Dengan network yang dibuat, maka dapat menggunakan CLI untuk explore detailnya. Perintah berikut akan mencantumkan semua network yang ada pada host.

![20](img/pict20.png)

Kita kemudian dapat explore network untuk melihat container mana yang dilampirkan dan alamat IP/IP address.

![21](img/pict21.png)

Perintah berikut ini digunakan untuk memutus container redis dari network frontend.

![22](img/pict22.png)
