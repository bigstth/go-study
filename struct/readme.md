เวลาส่งค่าเข้าไปใน function Go จะ copy variable ไปไว้ในอีก address memory ก่อนแล้วส่งอันนั้นไปในฟังก์ชั่น ทำให้ ค่าที่ได้ในฟังก์ชั่นจะมาจากคนละ address กับ variable ต้นทาง เวลาแก้ไขค่าไป ตัวต้นทางจะไม่ถูกแก้ไขด้วย ยกเว้นตัว referenceType ที่มันมี pointer ไปที่ originalValue ทำให้มันสามารถแก้ไขค่าต้นทางได้ ถึงแม้จะไม่ได้รับ \*type ในฟังก์ชั่น

turn address into value with \*address

turn value into address with &value

![typetobecare](https://raw.githubusercontent.com/bigstth/go-study/main/struct/typeToBeCare.png)
![referencetypeworking](https://raw.githubusercontent.com/bigstth/go-study/main/struct/referenceTypeWorking.png)
