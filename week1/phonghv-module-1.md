Week 1
============================

## 1. Caching
### 1.1. Tại sao ta cần phải dùng Cache?

- Cache giúp tăng tốc độ xử lý giải quyết vấn đề về hiệu suất và thắt cổ chai. Ví dụ: 100 request cần 100 lần query database để trả về chung một kết quả sẽ tốn thời gian. Chúng ta thay vào đó dùng cache để lưu kết quả của query ở request đầu tiên và 99 request còn lại chỉ cần vào cache lấy về.
- Cache giúp giảm chi phí cho ứng dụng. Chúng ta tưởng tượng ứng dụng sẽ đạt đến một lượng lớn request cần phải đảm bảo xử lý tốt và ổn định. Vì thế, chúng ta nghĩ đến việc nâng cấp CPU, RAM hay load balancing... Và cache cũng là một giải pháp nhưng lại không tốn thêm phí nên hạn chế bớt chi phí cho việc triển khai ứng dụng.
- => truy xuất dữ liệu từ ổ ứng lâu, chúng ta sẽ xảy ra hiện tượng thắt cổ chai nên cần một giải pháp tốt hơn đó là lưu trên RAM nên cache được sinh ra để giải quyết vấn đề này.

### 1.2.  Khi đồng bộ giữa Cache và Disk sẽ phát sinh các nguy cơ gì?

- Dữ liệu trên cache dễ bị lỗi thời. Ví dụ: người dùng thêm comment vào bài viết và load lại trang. Comment hiện thị lên cho bài viết nằm ở cache, người dùng không tìm thấy comment vừa thêm vào do dữ liệu trên cache chưa kịp update lại. Người dùng tưởng ứng dụng bị lỗi. Nếu chúng ta giải quyết trường hợp này bằng cách liên tục update dữ liệu giữa cache và database => Việc Caching không còn giá trị. Chúng ta có thể nghĩ đến việc tính toán dữ liệu nào nên dùng cache hay bao lâu nên là update dữ liệu trên cache một lần.
- Nếu khởi động lại server, dữ liệu trên cache mất hết. Chúng ta caching từ đầu thì nhiều dữ liệu trên cache trước đó có thể đang chờ xử lý hay chờ để lưu xuống database mất hết. Vì thế, chúng ta tính toán sau bao lâu nên backup lại cache một lần và restore lại khi start server. Nếu sau 5 phút backup 1 lần, ở phút 4 giây 59 server vì 1 lí do nào đó bị tạch thì dữ liệu trong 5 phút vừa qua cũng bay màu. Chúng ta cũng có thể nghĩ đến cách khác đó là ghi log lại tất cả các truy vấn mà cache đã làm nhưng file log sẽ rất nặng.

### 1.3. Có cách nào hạn chế những nguy cơ đó?

- Dùng những thư viện cache để hỗ trợ thao tác dữ liệu với cache trên server.
- Lưu trữ cache trong một server riêng, có sử dụng redis hoặc memcache giúp tính toán và lưu trữ cache.

## 2. Redis: khi nào dùng kiểu dữ liệu nào?

### 2.1. String

- String là kiểu dữ liệu cơ bản nhất của Redis và là string, float hoặc interger hay chuỗi binary. Redis có thể làm việc với cả string, từng phần của string hay tăng giảm giá trị của interger/float.
- Có 3 câu lệnh cơ bản với String là GET, SET và DEL.
- Chúng ta có thể sử dụng kiểu String cho nhiều trường hợp như redis key, lưu các đoạn hoặc trang HTML...

### 2.2. List

- Lưu trữ một danh sách có thứ tự của các string, là linked list.
- Nhờ vào cách lưu trữ, redis list có  thời gian add thêm 1 phần tử vào đầu hoặc cuối list rất nhanh. Nhưng việc truy xuất đến phần tử theo index của linked list là lâu hơn rất nhiều so với array.
- Redis thường được dùng trong 2 trường hợp phổ biến là lưu lại các post mới nhất của users trên mạng xã hội và giao tiếp giữa các tiến trình, sử dụng để xây dựng mô hình tương tác giữa consumer và producer, trong đó producer đưa item vào list và consumer dùng các item này theo thứ tự quy định trong list.

### 2.3. Hash

- Lưu trữ tập các map của key và value. Các field và giá trị đều là kiểu Redis string. Key được sắp xếp ngẫu nhiên, không theo thứ tự nào cả.
- Hash thường dùng để biểu diễn một object.
- Tổ chức dữ liệu với hash ít tiêu tốn tài nguyên hơn là string và thứ 2 là thời gian expire time tức là thay vì tổ chức dữ liệu theo string thì expire time cho từng key việc quản lí khó lỏng lẻo còn đối với hash chỉ cần 1 expire time, hết hạn dữ liệu trong key được hủy hết an toàn chặt chẽ.

### 2.4. Set

- Là tập hợp các string, các phần tử không được sắp xếp.
- Set rất tốt để thể hiện quan hệ giữa các object.
- Ví dụ: chúng ta có 1 set lưu id của các nhân viên lập trình nhóm mobile và 1 set lưu id của các nhân viên lập trình nhóm java. chúng ta cần tìm các nhân viên vừa thuộc nhóm mobile vừa thuộc nhóm java thì chỉ cần union 2 set này lại với nhau.

### 2.5. Sorted Set

- Giống như Set nhưng mỗi phần tử là map của 1 string (member) và 1 floating-point number (score), nó được sắp xếp theo score này.
- Sorted Set dùng cho những bài toán xếp hạng, tìm vị trí top.
- Quan trọng nhất của zset là việc tự đông sắp xếp theo score.

## 3. Golang: goroutine, channel, interfaces có ưu điểm gì, khi nào dùng?

### 3.1. Goroutine

- Gorountine là một hàm thực thi đồng thời các hàm hay method khác trong cùng một không gian địa chỉ.
- Ưu điểm:
    - Chi phí tạo ra goroutine nhỏ hơn so với tạo ra một thread. Kích thước stack khởi tạo cho goroutine nhỏ và nó có thể tự tăng giảm theo nhu cầu của chương trình còn kích thước cho thread thì cần chỉ định và cố định. Do đó, thông thường các ứng dụng Go có hàng ngàn Goroutines chạy đồng thời.
    - Chúng ta chỉ cần khai báo từ khóa go trước cách gọi hàm thông thường hay trước một phương thức, hàm hay phương thức đó sẽ được xử lý đồng thời.
    - Go thiết kế che dấu nhiều sự phức tạp của việc tạo và quản lý luồng dễ dàng cho lập trình viên sử dụng. Có thể một thread trong chương trình có hàng nghìn goroutines. Nếu bất kì goroutine nào trong khối thread này đợi I/0 chẳng hạn, thì một OS thread khác sẽ được tạo và các goroutine còn lại được chuyển sang OS thread mới. Tất cả điều này được bộ xử lý của GO làm hết.
    - Goroutine giao tiếp với nhau bằng channels.

### 3.2. Channel

- Channel dùng trong việc trao đổi dữ liệu giữa các goroutine trong xử lý đồng thời. Mỗi channel chỉ trao đổi một loại dữ liệu.
- Ưu điểm:
    + Channels có thể được coi là các đường ống, dữ liệu có thể được gửi từ một đầu và nhận từ đầu kia bằng channels. Một hành động khác là đóng kênh. An toàn, dễ dàng trong việc trao đổi dữ liệu.
    + Mỗi channel chỉ trao đổi một loại dữ liệu, chúng ta chắc chắn được dữ liệu mà chúng ta trao đổi.
    + Có 2 loại channel: unbuffered channel và buffered channel.
    + Unbuffered channel chỉ chứa tối đa một giá trị dữ liệu trong channle. Unbuffered channel đòi hỏi đồng bộ trong gửi nhận: khi gửi thì goroutine thực hiện gửi bị khóa cho đến khi việc nhận kết thúc và goroutine thực hiện nhận bị khóa cho đến khi nhận xong dữ liệu từ channel.
    + Buffered channel có thể chứa nhiều hơn một giá trị dữ liệu và chứa tối đa số giá trị được khai báo. Goroutine gửi chỉ bị khóa khi khả năng chứa của channel bị đầy và goroutine nhận chỉ bị khóa khi channel trống dữ liệu.
    
### 3.3. Interface

- Interface dùng khi muốn tạo phương thức chung cho các đối tượng hoặc muốn giao tiếp qua thông qua một đối tượng chung.
- Ưu điểm:
    + Các Interface làm cho code linh hoạt hơn, có thể mở rộng và đó là cách để đạt được tính đa hình trong Golang.
    + Interface là kiểu dữ liệu trừu tượng, chỉ khai báo các phương thức mô tả hành động, không khai báo dữ liệu lưu trữ. Khi các kiểu dữ liệu khác có khai báo đầy đủ các phương thức mà interface đã khai báo thì các kiểu dữ liệu này đáp ứng yêu cầu và các biến của chúng đều có thể được lưu trữ và sử dụng bởi các biến kiểu interface.
    + Interface có trường hợp đặc biệt là interface rỗng với khai báo interface{}. Lúc này mọi kiểu dữ liệu đều đáp ứng nó và có thể dùng nó đại diện cho các kiểu dữ liệu ở mọi nơi.