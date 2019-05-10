Warm up
============================

## 1. Mục tiêu cần đạt được:

- Giúp fresher làm quen với môi trường làm việc trên Linux.
- Sử dụng thành thạo một số công cụ có sẵn của hệ thống: vim, vscode, .... Tối thiểu cần cài đặt và build được ứng dụng Hello World! bằng Go trên [Visual Studio Code](https://rominirani.com/setup-go-development-environment-with-visual-studio-code-7ea5d643a51a)
- Nắm vững các nội dung quan trọng của lập trình
- Trao dồi kĩ năng tìm kiếm với Google

## 2. Nguồn tham khảo

- [10 Common Software Architectural Patterns in a nutshell](https://towardsdatascience.com/10-common-software-architectural-patterns-in-a-nutshell-a0b47a1e9013)
- [Golang](https://techtalk.vn/golang-thuc-su-tot-trong-truong-hop-nao.html)
- Google

## 3. Báo cáo công việc
Viết note markdown trình bày các vấn đề tìm hiểu được. Khi gửi mail báo cáo thì:

+ Nên zip lại
+ Đăt tên file "account-module-1.zip", trong đó <account> là account của mỗi người
+ Gửi cho `hieunv2@vng.com.vn`
+ Ghi tiêu đề email: Báo cáo module 1
+ Nội dung email cần theo format:

```
Chào anh,

// .. ghi gì đó vào đây

Cảm ơn anh.
```
Ví dụ:

+ Trong mail cần liệt kê rõ các nội dung chính đã làm được, các nội dung chưa làm được (nếu có), các vấn đề đang gặp khó khăn, khúc mắc không thể tự giải đáp.
+ Cách viết email như trên sẽ áp dụng cho mọi báo cáo khác về sau
+ File báo cáo phải viết dưới bằng [Markdown](https://en.wikipedia.org/wiki/Markdown) (phần mở rộng là `*.md`
+ Trong báo cáo có thể reference tài liệu => nhưng bắt buộc phải ghi lại gạch đầu dòng những ý chính trong báo cáo
Ví dụ:
Chào anh!
Em đã làm xong phần việc của tuần 1.
Nội dung em đã hoàn tất: .................... 
Chi tiết em đã gửi trong file zip đính kèm.
Tuy nhiên, trong quá trình tìm hiểu, em thấy vẫn còn vấn đề chưa hiểu rõ là: ....................
Nhờ anh hổ trợ thêm giúp em vấn đề này nhé!
Em cảm ơn anh!
Em chào anh!

## 4. Nội dung


### 4.1 Caching

- Vai trò của cache. Trả lời được câu hỏi: Tại sao ta cần phải dùng Cache?)

- Vấn đề đồng bộ giữa Cache và Disk. Trả lời được câu hỏi: khi đồng bộ giữa Cache và Disk sẽ phát sinh các nguy cơ gì? Có cách nào hạn chế những nguy cơ đó?

### 4.2 Redis 

- Khái niêm cơ bản về Redis: [tổng quan ](https://viblo.asia/p/tong-quan-ve-redis-NznmMdXzMr69)

- Danh sách [command Redis](https://redis.io/commands)

- Nắm được các 5 loại kiểu dữ liệu string, list, hash, set, zset. Trả lời được câu hỏi: khi nào dùng kiểu dữ liệu nào.

- Nắm được cách đặt tên key cho phù hợp (foo:*....)

- Cài đặt được [Redis server](https://www.linode.com/docs/databases/redis/how-to-install-a-redis-server-on-ubuntu-or-debian8/#update-and-install)

- Bài tập căn bản: Viết chương trình đơn giản sử dụng các command Redis:Get, Set, Del, exists, HSet, HGet, HExists, LPush, LPop, LRem, SAdd, ZAdd, ZRem, ZRange. [OpenSourrce] (https://github.com/go-redis/redis)

- Bài tập nâng cao: Ứng dụng Redis để kiểm tra Access Token: Khi user login thành công sẽ được cấp Access Token và Tên hiển thị (Display name). Mỗi khi thực hiện tác vụ getDisplayName cần kiểm tra access token có đúng hay không, nếu có, trả về display name, nếu không báo lỗi và yêu cầu đăng nhập lại. Mỗi khi login lại, access Token sẽ được thay mới. Mỗi access Token có thời gian tồn tại là 5 phút, sau thời gian trên, bắt buộc phải login lại để cấp access Token mới. Yêu cầu dùng redis tối đa nhất có thể.

### 4.3 Golang in action

- [goroutine](https://tour.golang.org/concurrency/1)

- [channel](https://tour.golang.org/concurrency/2)

- [interfaces](https://gobyexample.com/interfaces)

- Trả lời được câu hỏi: goroutine, channel, interfaces có ưu điểm gì, khi nào dùng?

- Bài tập căn bản: Viết ví dụ dùng goroutine và channel

- Bài tập nâng cao: viết ví dụ (Singleton)[(https://gpcoder.com/4190-huong-dan-java-design-pattern-singleton/)] bằng Golang

### 5. Thời gian hoàn thành 

- Max 2 tuần 

- Hoàn thành càng sớm thì điểm quá trình càng cao 

