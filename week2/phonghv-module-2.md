Week 2
============================

## 1. Message Queue

### 1.1. Tại sao ta cần phải dùng Message Queue?

https://aws.amazon.com/message-queue/


### 1.2. Phân biệt các loại message queue hiện có, ví dụ: Rabbit MQ - Kafka - Active MQ

https://stackshare.io/stackups/activemq-vs-kafka-vs-rabbitmq

https://medium.com/@anhle128/ph%C3%A2n-lo%E1%BA%A1i-gi%E1%BB%AFa-c%C3%A1c-h%E1%BB%87-th%E1%BB%91ng-message-queue-32aea38e2066

https://dzone.com/articles/exploring-message-brokers

### 1.3. Publish/‎Subscribe và Produce/Consume. Lúc nào thì dùng cơ chế nào?


## 3. gRPC trong Golang

### 3.1. Protobuf là gì?

- Serializable là kỹ thuật chuyển đổi một cấu trúc dữ liệu hay đối tượng thành dạng byte stream dùng để lưu trữ hoặc truyền qua mạng.
- Deserialization là hoạt động ngược lại của Serializable đọc byte stream và chuyển đổi ngược lai.
- Protobuf là  a language-neutral, platform-neutral, sử dụng để serialize dữ liệu thành dạng byte stream.
- Dữ liệu được serialize bởi protobuf có thể được deserialize bằng các ngôn ngữ khác nhau.

### 3.2. gRPC có ưu điểm gì so với RestAPI của Protocol HTTP?
 
 - gRPC sử dụng protobuf nên làm cho payload nhanh hơn, nhỏ hơn và đơn giản hơn. Dữ liệu gửi bằng gPRC an toàn, bảo mật hơn so với xml hay json của RestAPI do khó mà deserialize nếu không biết schema.
 - Một điều nữa, nhờ vào việc ta định nghĩa dữ liệu nhận được trong message của protobuf đảm bảo tính đúng đắn mà dữ liệu nhận được. Ví dụ: server đợi tham số là một số nguyên, gRPC sẽ không cho phép ta gửi một chuỗi lên còn với RestAPI là xml hay json nên chúng ta không chắc chắn được điều này.
 - Chúng ta có thể sử dụng gRPC trên nhiều ngôn ngữ khác nhau. Ví dụ: ta có một web service viết bằng Go, một ứng dụng viết bằng Java vẫn có thể sử dụng web service này thành ra chúng ta thấy web service gRPC có khả năng mở rộng cao.
 - gRPC sử dụng HTTP/2 nên có hiệu suất rất cao và các API sử dụng dữ liệu nhị phân thay vì xml hay json như REST API làm cho việc truyền tải gọn nhẹ và hiểu quả hơn. Và đồng thời, gPRC sử dụng HTTP/2 tốt hơn RestAPI.