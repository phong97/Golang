Boot
============================

## 1. Mục tiêu cần đạt được:

- Giúp fresher làm quen với Kafka apache message queue.
- Giúp fresher làm quen với gRPC trong Go.

## 2. Nguồn tham khảo

- [Tổng quan Kafka](https://viblo.asia/p/kafka-la-gi-gDVK2Q7A5Lj)
- [Tổng quan Protobuf](https://developers.google.com/protocol-buffers/?hl=vi)
- [Ví dụ gRPC đơn giản](https://kipalog.com/posts/Su-dung-Protobuf-va-gRPC-de-phat-trien-API-hieu-nang-cao)

## 3. Báo cáo công việc
Viết note markdown trình bày các vấn đề tìm hiểu được. Khi gửi mail báo cáo thì:

+ Nên zip lại
+ Đăt tên file "account-module-1.zip", trong đó <account> là account của mỗi người
+ Gửi cho `hieunv2@vng.com.vn`, `tienht@vng.com.vn`
+ Ghi tiêu đề email: Báo cáo module 2
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
Em đã làm xong phần việc của tuần 2.
Nội dung em đã hoàn tất: .................... 
Chi tiết em đã gửi trong file zip đính kèm.
Tuy nhiên, trong quá trình tìm hiểu, em thấy vẫn còn vấn đề chưa hiểu rõ là: ....................
Nhờ anh hổ trợ thêm giúp em vấn đề này nhé!
Em cảm ơn anh!
Em chào anh!

## 4. Nội dung


### 4.1 Message Queue

- Vai trò của Message Queue. Trả lời được câu hỏi: Tại sao ta cần phải dùng Message Queue?

- Phân biệt các loại message queue hiện có, ví dụ: Rabbit MQ - Kafka - Active MQ 

- Phân biệt hai cơ chế: Publish/‎Subscribe và Produce/Consume. Trả lời được câu hỏi: Lúc nào thì dùng cơ chế nào?

### 4.2 Kafka Apache Message Queue

- Các đối tượng thường gặp trong Kafka. Trả lời được câu hỏi: nêu chi tiết về các định nghĩa có trong Kafka: broker, cluster, message, topic, partition, producer, consumer, consumer group ID

- Phân biệt được cái cơ chế stratergy trong kafka, cách phân bố partitioner, cách dùng cluster

- Trả lời được câu hỏi: các nguy cơ thường gặp khi dùng Kafka Message Queue

- Cài đặt được [Kafka server](https://www.digitalocean.com/community/tutorials/how-to-install-apache-kafka-on-ubuntu-18-04)

- Bài tập căn bản: Viết chương trình đơn giản produce và consume một message bằng Kafka, quản lý offset(dùng markoffset, và commitoffset)

### 4.3 gRPC trong Golang

- [Tổng quan Protobuf](https://viblo.asia/p/protocol-buffers-la-gi-va-nhung-dieu-can-ban-can-biet-ve-no-maGK7D99Zj2)

- [Tổng quan gRPC](https://viblo.asia/p/grpc-va-ung-dung-no-trong-microservices-ORNZqo8N50n)

- Trả lời được câu hỏi: Protobuf là gì? gRPC có ưu điểm gì so với RestAPI của Protocol HTTP?

- Bài tập căn bản: Viết ví dụ dùng gRPC có handle http gateway, có xài thêm swagger để định nghĩa api đã viết.

- Bài tập nâng cao: viết lại bài tập thao tác với Access Token và Redis trong week 1. Tuy nhiên, sử dụng gRPC Stub và gRPC server để call/reply các API gọi tắt serviceA. Đồng thời dùng kafka để cho serviceB lấy về rồi tính token dùng jwt. Sau đó service A sẽ request token từ service B(service B có grpc api để service A gọi lấy token). Lúc này service B sẽ lưu redis có set expiration time cho lần gọi get token.

### 5. Thời gian hoàn thành 

- Max 1.5 tuần 

- Hoàn thành càng sớm thì điểm quá trình càng cao 

