เวลาส่งค่าเข้าไปใน function Go จะ copy variable ไปไว้ในอีก address memory ก่อนแล้วส่งอันนั้นไปในฟังก์ชั่น ทำให้ ค่าที่ได้ในฟังก์ชั่นจะมาจากคนละ address กับ variable ต้นทาง เวลาแก้ไขค่าไป ตัวต้นทางจะไม่ถูกแก้ไขด้วย ยกเว้นตัว referenceType ที่มันมี pointer ไปที่ originalValue ทำให้มันสามารถแก้ไขค่าต้นทางได้ ถึงแม้จะไม่ได้รับ \*type ในฟังก์ชั่น
![typetobecare](https://raw.githubusercontent.com/bigstth/go-learning/main/go-struct/typeToBeCare.png)
![referencetypeworking](https://raw.githubusercontent.com/bigstth/go-learning/main/go-struct/referenceTypeWorking.png)
