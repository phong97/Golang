Week 2
============================

## 1. Message Queue

### 1.1. Tại sao ta cần phải dùng Message Queue?
- Ứng dụng của chúng ta được xây dựng từ nhiều service độc lập. Một vấn đề quan trọng sinh ra là kết nối và giao tiếp giữa các service này. Nhưng nếu chúng ta sử dụng mô hình Point-To-Point, các service kết nối trực tiếp với nhau thì chỉ có thể áp dụng cho hệ thống có số lượng service nhỏ. Đó là lí do tại sao chúng ta nên sử dụng mô hình Message Queue.
- Các service sẽ không giao tiếp trực tiếp với nhau mà thông qua Message Queue. Nó cung cấp một giao thức giao tiếp bất đồng bộ, tức là trong một quá trình consumer không xử lý các message trong queue thì các message khác vẫn được thêm vào queue từ producer và được xử lý khi consumer sẵn sàng.
- Từ đó, chúng ta nhận thấy các phần khác nhau trong ứng dụng có thể phát triển độc lập hay được viết bằng các ngôn ngữ khác nhau hoặc được duy trì và phát triển bởi các nhóm lập trình riêng biệt.
- Đồng thời, khi khối lượng message tăng lên và chúng ta cần mở rộng ứng dụng của mình, đơn giản tất cả những gì chúng ta cần làm là thêm nhiều consumer, producer để xử lý queue nhanh hơn.

### 1.2. Phân biệt các loại message queue hiện có, ví dụ: Rabbit MQ - Kafka - Active MQ

https://stackshare.io/stackups/activemq-vs-kafka-vs-rabbitmq

https://medium.com/@anhle128/ph%C3%A2n-lo%E1%BA%A1i-gi%E1%BB%AFa-c%C3%A1c-h%E1%BB%87-th%E1%BB%91ng-message-queue-32aea38e2066

https://dzone.com/articles/exploring-message-brokers

### 1.3. Publish/‎Subscribe và Produce/Consume. Lúc nào thì dùng cơ chế nào?

- Publish/‎Subscribe:
    + là người gửi (publisher) gửi message, message sẽ này không được gửi trực tiếp đến người nhận (subscriber) và không biết trước người nhận là ai. Các message được tổ chức thành các lớp (topic).
    + Tương tự, người nhận subscribe vào một hay nhiều topic và chỉ nhận message mình mong muốn mà không biết người gửi là ai.
    + Để thực hiện điều này có một message broker tiếp nhận các message và chuyển chúng đến người nhận mong muốn.
    + Publish/‎Subscribe sử dụng để kết nối các mesage giữa các hệ thống khác nhau mà các hệ thống đó không cần biết thông tin cụ thể về hệ thống khác hay sử dụng cơ chế này hỗ trợ realtime cho ứng dụng của chúng ta.

- Produce/Consume:
    + là người gửi gửi message lên queue và người nhận lấy message theo thứ tự quy định của queue.
    + Nếu queue đầy, những người gửi sẽ đợi đến khi có người nhận lấy một message khỏi queue rồi mới bắt đầu gửi message trởi lại.
    + Tương tự, nếu queue rỗng, những người nhận sẽ đợi đến khi có message được gửi vào queue.
    + Produce/Consume được sử dụng để chia sẻ message trên những hệ thống giống nhau nhằm chia nhỏ công việc hay cân bằng tải...
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