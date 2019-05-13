Week 1
============================

## 1. Caching
### 1.1. Tại sao ta cần phải dùng Cache?

- Cache giúp tăng tốc độ xử lý giải quyết vấn đề về hiệu suất và thắt cổ chai. Ví dụ: 100 request cần 100 lần query database để trả về chung một kết quả sẽ tốn thời gian. Chúng ta thay vào đó dùng cache để lưu kết quả của query ở request đầu tiên và 99 request còn lại chỉ cần vào cache lấy về.
- Cache giúp giảm chi phí cho ứng dụng. Chúng ta tưởng tượng ứng dụng sẽ đạt đến một lượng lớn request cần phải đảm bảo xử lý tốt và ổn định. Vì thế, chúng ta nghĩ đến việc nâng cấp CPU, RAM hay load balancing... Và cache cũng là một giải pháp nhưng lại không tốn thêm phí nên hạn chế bớt chi phí cho việc triển khai ứng dụng.

### 1.2.  Khi đồng bộ giữa Cache và Disk sẽ phát sinh các nguy cơ gì?

- Dữ liệu trên cache dễ bị lỗi thời. Ví dụ: người dùng thêm comment vào bài viết và load lại trang. Comment hiện thị lên cho bài viết nằm ở cache, người dùng không tìm thấy comment vừa thêm vào do dữ liệu trên cache chưa kịp update lại. Người dùng tưởng ứng dụng bị lỗi. Nếu chúng ta giải quyết trường hợp này bằng cách liên tục update dữ liệu giữa cache và database => Việc Caching không còn giá trị. Chúng ta có thể nghĩ đến việc tính toán dữ liệu nào nên dùng cache hay bao lâu nên là update dữ liệu trên cache một lần.
- Nếu khởi động lại server, dữ liệu trên cache mất hết. Chúng ta caching từ đầu thì nhiều dữ liệu trên cache trước đó có thể đang chờ xử lý hay chờ để lưu xuống database mất hết. Vì thế, chúng ta tính toán sau bao lâu nên backup lại cache một lần và restore lại khi start server. Nếu sau 5 phút backup 1 lần, ở phút 4 giây 59 server vì 1 lí do nào đó bị tạch thì dữ liệu trong 5 phút vừa qua cũng bay màu. Chúng ta cũng có thể nghĩ đến cách khác đó là ghi log lại tất cả các truy vấn mà cache đã làm nhưng file log sẽ rất nặng.

### 1.3. Có cách nào hạn chế những nguy cơ đó?

- Dùng những thư viện cache để hỗ trợ thao tác dữ liệu với cache trên server.
- Lưu trữ cache trong một server riêng, có sử dụng redis hoặc memcache giúp tính toán và lưu trữ cache.

## 2. Redis: khi nào dùng kiểu dữ liệu nào?

### 2.1. String
- 

### 2.2. List

- Lưu trữ danh sách 