interface 101

type ที่มี element ข้างใน เหมือน interface จะถูกนับว่ามี type เหมือน interface นั้นอัตโนมัติ
ทำให้เราสามารถส่ง eb หรือ sb เข้าไปใน printGreeting() ที่รับ type เป็น bot(common name) และเรียกใช้งาน b.getGreeting ได้เลย

เพื่อลดความซ้ำซ้อน ต้องมาทำแยกทีละฟังก์ชั่นเช่น printEnglishGreeting , print...

interface ไม่สามารถ assign ค่าได้ ต้องเป็น concreteType

![concretevsinterface](https://raw.githubusercontent.com/bigstth/go-study/main/interface/concreteTypeVSinterfaceType.png)
